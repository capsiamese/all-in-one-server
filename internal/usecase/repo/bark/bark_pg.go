package bark

import (
	"aio/internal/entity"
	"aio/internal/pb"
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
		Select("device_token", "device_key", "name").
		From(DeviceTable)

	obj := &entity.BarkDevice{}
	if device.DeviceToken != "" {
		if err := r.p.GetCache(ctx, device.DeviceToken, obj); err != nil {
			if !errors.Is(err, postgres.ErrNotSetCache) {
				r.l.Errorln(err)
			}
		} else {
			return obj, nil
		}
		b = b.Where(squirrel.Eq{"device_token": device.DeviceToken})
	} else if device.DeviceKey != "" {
		if err := r.p.GetCache(ctx, device.DeviceKey, obj); err != nil {
			if !errors.Is(err, postgres.ErrNotSetCache) {
				r.l.Errorln(err)
			}
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
	err := gob.NewEncoder(buf).Encode(message.Params)
	if err != nil {
		return err
	}
	b := r.p.Builder.Insert(HistoryTable).
		Columns("device_key", "device_token", "ts", "send_from", "title", "content", "params").
		Values(device.DeviceKey, device.DeviceToken, time.Now().Unix(), device.Name, message.Title, message.Content, buf.Bytes())
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

func (r *Repo) FetchHistory(ctx context.Context, device *entity.BarkDevice, offset, limit int) ([]*pb.BarkHistory, error) {
	q := r.p.Builder.
		Select("id", "device_key", "ts", "send_from", "title", "content", "params").
		From(HistoryTable).
		Where(squirrel.Eq{"device_token": device.DeviceToken}).
		Offset(uint64(offset)).
		Limit(uint64(limit))
	query, params, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.p.X.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.BarkHistory, 0)
	for rows.Next() {
		item := new(pb.BarkHistory)
		p := make([]byte, 0)
		err = rows.Scan(&item.Id, &item.Key, &item.Ts, &item.From, &item.Title, &item.Content, &p)
		if err != nil {
			return nil, err
		}
		item.Params = make(map[string]string)
		err = gob.NewDecoder(bytes.NewReader(p)).Decode(&item.Params)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (r *Repo) DropHistory(ctx context.Context, key string, id int64) error {
	q, p, err := r.p.Builder.
		Delete(HistoryTable).
		Where(squirrel.Eq{"device_key": key, "id": id}).
		ToSql()
	if err != nil {
		return err
	}
	res, err := r.p.X.ExecContext(ctx, q, p...)
	if err != nil {
		return err
	}
	if i, err := res.RowsAffected(); err != nil {
		r.l.Warningf("drop history key:%s id:%d %w", key, id, err)
	} else if i != 1 {
		r.l.Infoln("drop history id:%d not exist", id)
	}
	return nil
}
