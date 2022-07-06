package usecase

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"notification/ent"
	"time"
)

var _ Extension = (*ExtensionUseCase)(nil)

type ExtensionUseCase struct {
	cli *ent.Client
}

func NewExtension(cli *ent.Client) *ExtensionUseCase {
	return &ExtensionUseCase{
		cli: cli,
	}
}

func (e *ExtensionUseCase) Register(ctx context.Context, name, extensionID string) (*ent.ExtensionClient, error) {
	ext, err := e.cli.ExtensionClient.
		Create().
		SetExtensionID(extensionID).
		SetName(name).
		SetLastAccessTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_ = ext
	return nil, nil
}

func (e *ExtensionUseCase) Connect(ctx context.Context, uid uuid.UUID, wsConn *websocket.Conn) error {
	return errors.New("")
}
