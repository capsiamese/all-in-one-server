package usecase

import (
	"context"
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

func (e *ExtensionUseCase) Register(ctx context.Context) error {
	ext, err := e.cli.ExtensionClient.
		Create().
		SetExtensionID("").
		SetName("").
		SetLastAccessTime(time.Now()).
		Save(ctx)
	if err != nil {
		return err
	}
	_ = ext
	return nil
}
