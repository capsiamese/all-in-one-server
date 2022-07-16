package webapi

import (
	"aio/internal/entity"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"context"
)

var _ usecase.BarkWebAPI = (*DumbBarkAPNsAPI)(nil)

type DumbBarkAPNsAPI struct {
	l logger.Interface
}

func NewDumbBark(l logger.Interface) usecase.BarkWebAPI {
	return &DumbBarkAPNsAPI{
		l: l,
	}
}

func (d *DumbBarkAPNsAPI) Push(ctx context.Context, msg *entity.APNsMessage) error {
	d.l.Infoln("dumb bark")
	return nil
}
