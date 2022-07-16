package bark

import (
	"aio/internal/entity"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"aio/pkg/postgres"
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"time"
)

var _ usecase.BarkRepo = (*Repo)(nil)

const DeviceTable = "t_bark_device"
const HistoryTable = "t_bark_history"

type Repo struct {
	p *postgres.Postgres
	l logger.Interface
}

func New(p *postgres.Postgres, l logger.Interface) *Repo {
	return &Repo{p, l}
}

func (r *Repo) Store(ctx context.Context, device *entity.BarkDevice) error {
	if device.DeviceToken == "" ||
		device.DeviceKey == "" {
		return fmt.Errorf("invalid device")
	}
	sql, args, err := r.p.Builder.
		Insert(DeviceTable).
		Columns("device_key, device_token", "name").
		Values(device.DeviceKey, device.DeviceToken, device.Name).
		ToSql()
	if err != nil {
		return err
	}

	r.l.Infoln(sql, args)

	_, err = r.p.X.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	if err = r.p.SetCache(ctx, device.DeviceToken, device); err != nil && !errors.Is(err, postgres.ErrNotSetCache) {
		r.l.Errorln(err)
	}

	if err = r.p.SetCache(ctx, device.DeviceKey, device); err != nil && !errors.Is(err, postgres.ErrNotSetCache) {
		r.l.Errorln(err)
	}

	return nil
}

func (r *Repo) Get(ctx context.Context, device *entity.BarkDevice) (*entity.BarkDevice, error) {
	b := r.p.Builder.
		Select("device_token", "device_key").
		From(DeviceTable)

	obj := &entity.BarkDevice{}
	if device.DeviceToken != "" {
		if err := r.p.GetCache(ctx, device.DeviceToken, obj); err != nil && !errors.Is(err, postgres.ErrNotSetCache) {
			r.l.Errorln(err)
		} else {
			return obj, nil
		}
		b = b.Where(squirrel.Eq{"device_token": device.DeviceToken})
	} else if device.DeviceKey != "" {
		if err := r.p.GetCache(ctx, device.DeviceKey, obj); err != nil && !errors.Is(err, postgres.ErrNotSetCache) {
			r.l.Errorln(err)
		} else {
			return obj, nil
		}
		b = b.Where(squirrel.Eq{"device_key": device.DeviceKey})
	} else {
		return nil, fmt.Errorf("invalid device")
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	r.l.Infoln(sql, args)

	err = r.p.X.GetContext(ctx, obj, sql, args...)
	return obj, err
}

func (r *Repo) SaveMessage(ctx context.Context, device *entity.BarkDevice, message *entity.APNsMessage) error {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := gob.NewEncoder(buf).Encode(message.Data)
	if err != nil {
		return err
	}
	b := r.p.Builder.Insert(HistoryTable)
	b.Columns("device_key", "device_token", "data", "ts", "send_from")
	b.Values(device.DeviceKey, device.DeviceToken, buf.Bytes(), time.Now().Unix(), device.Name)
	q, p, err := b.ToSql()
	if err != nil {
		return err
	}
	_, err = r.p.X.ExecContext(ctx, q, p...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FetchHistory(ctx context.Context, device *entity.BarkDevice, offset, limit int) ([]*entity.BarkHistory, error) {
	q := r.p.Builder.Select(HistoryTable)
	q.Columns("id", "device_key", "device_token", "data", "ts", "send_from")
	q.Where(squirrel.Eq{"device_token": device.DeviceToken})
	q.Offset(uint64(offset))
	q.Limit(uint64(limit))
	query, params, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.p.X.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	list := make([]*entity.BarkHistory, 0)
	for rows.Next() {
		item := new(entity.BarkHistory)
		data := make([]byte, 0)
		err = rows.Scan(&item.Id, &item.DeviceKey, &item.DeviceToken, &data, &item.Ts, &item.From)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		m := make(map[string]any)
		err = gob.NewDecoder(buf).Decode(&m)
		if err != nil {
			return nil, err
		}
		item.Data = m
		list = append(list, item)
	}
	return list, nil
}
