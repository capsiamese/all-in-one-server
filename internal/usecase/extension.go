package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"notification/ent"
	"notification/ent/extensionclient"
	"notification/internal/entity"
	"notification/pkg/logger"
	"time"
)

var _ Extension = (*ExtensionUseCase)(nil)

type ExtensionUseCase struct {
	cli *ent.Client
	log logger.Interface
}

func NewExtension(cli *ent.Client, l logger.Interface) *ExtensionUseCase {
	return &ExtensionUseCase{
		cli: cli,
		log: l,
	}
}

func (e *ExtensionUseCase) tx(ctx context.Context, callback func(ctx context.Context, tx *ent.Tx) error) error {
	tx, err := e.cli.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return err
	}
	err = callback(ctx, tx)
	if err != nil {
		return fmt.Errorf("transection:%v Rollback:%v", err, tx.Rollback())
	}
	return tx.Commit()
}

func (e *ExtensionUseCase) Register(ctx context.Context, name, extensionID string) (*ent.ExtensionClient, error) {
	ext, err := e.cli.ExtensionClient.Query().
		Where(extensionclient.And(
			extensionclient.NameEQ(name),
			extensionclient.ExtensionIDEQ(extensionID),
		)).Only(ctx)

	if err == nil {
		return ext, nil
	} else {
		e.log.Warningln("query client:%v", err)
	}

	ext, err = e.cli.ExtensionClient.
		Create().
		SetName(name).
		SetExtensionID(extensionID).
		SetClientUID(uuid.NewV4()).
		SetLastAccessTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return ext, nil
}

func (e *ExtensionUseCase) Connect(ctx context.Context, uid uuid.UUID, wsConn *websocket.Conn) error {
	return errors.New("")
}

func (e *ExtensionUseCase) Add(c context.Context, uid uuid.UUID, group ...*entity.GroupInfo) error {
	return e.tx(c, func(ctx context.Context, tx *ent.Tx) error {
		var err error
		var client *ent.ExtensionClient
		client, err = tx.ExtensionClient.Query().Where(extensionclient.ClientUIDEQ(uid)).Only(ctx)
		if err != nil {
			return err
		}
	LOOP:
		for _, g := range group {
			create := tx.Group.Create().
				SetUID(uuid.NewV4()).
				SetCreatedAt(time.Now()).
				SetClient(client)
			if g.Name != "" {
				create.SetName(g.Name)
			}
			if g.Option != nil {
				create.SetOption(*g.Option)
			}
			var eg *ent.Group
			eg, err = create.Save(ctx)
			if err != nil {
				break
			}
			for i, t := range g.Tabs {
				_, err = tx.Tab.Create().
					SetName(t.Name).
					SetURL(t.Url).
					SetSeq(int32(i + 1)).
					SetFavicon(t.Favicon).
					SetGroup(eg).
					Save(ctx)
				if err != nil {
					break LOOP
				}
			}
		}
		return nil
	})
}

func (e *ExtensionUseCase) Pull(ctx context.Context, uid uuid.UUID) (*ent.ExtensionClient, error) {
	return e.cli.ExtensionClient.
		Query().
		Where(
			extensionclient.ClientUIDEQ(uid),
		).
		WithGroups(
			func(query *ent.GroupQuery) {
				query.WithTabs()
			},
		).
		Only(ctx)
}
