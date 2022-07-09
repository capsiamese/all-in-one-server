package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"notification/ent"
	"notification/ent/extensionclient"
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

func (e *ExtensionUseCase) Register(ctx context.Context, name, extensionID string) (*ent.ExtensionClient, error) {
	ext, err := e.cli.ExtensionClient.Query().
		Where(extensionclient.And(
			extensionclient.NameEQ(name),
			extensionclient.ExtensionIDEQ(extensionID),
		)).First(ctx)

	if err == nil {
		return ext, nil
	} else {
		if !errors.Is(err, sql.ErrNoRows) {
			e.log.Errorln("query extension client:%v", err)
		}
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

func (e *ExtensionUseCase) Add(ctx context.Context, name string, group ...*ent.Group) error {
	return errors.New("implement me")
}
