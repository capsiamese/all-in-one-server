package usecase

import (
	"context"
	"notification/ent"
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
	return nil
}
