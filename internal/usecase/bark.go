package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"notification/internal/entity"
	"notification/pkg/logger"
)

var _ Bark = (*BarkUseCase)(nil)

type BarkUseCase struct {
	repo BarkRepo
	api  BarkWebAPI
	log  logger.Interface
}

func NewBark(r BarkRepo, a BarkWebAPI, l logger.Interface) *BarkUseCase {
	return &BarkUseCase{
		repo: r,
		api:  a,
		log:  l,
	}
}

func (uc *BarkUseCase) Push(ctx context.Context, key string, msg *entity.APNsMessage) error {
	d, err := uc.repo.Get(ctx, &entity.BarkDevice{DeviceKey: key})
	if err != nil {
		return err
	}
	msg.DeviceToken = d.DeviceToken
	err = uc.repo.SaveMessage(ctx, msg)
	if err != nil {
		uc.log.Errorln("save apns message %w", err)
	}
	return uc.api.Push(ctx, msg)
}

func (uc *BarkUseCase) Register(ctx context.Context, device *entity.BarkDevice) error {
	if device.DeviceToken == "" {
		return fmt.Errorf("device token is empty")
	}

	if device.DeviceKey != "" {
		obj, err := uc.repo.Get(ctx, &entity.BarkDevice{DeviceKey: device.DeviceKey})
		if err != nil {
			return err
		}
		if obj.DeviceToken == device.DeviceToken {
			return nil
		}
	}

	obj, err := uc.repo.Get(ctx, &entity.BarkDevice{DeviceToken: device.DeviceToken})
	if errors.Is(err, sql.ErrNoRows) {
		device.DeviceKey = uuid.NewV4().String()
		return uc.repo.Store(ctx, device)
	}
	if err != nil {
		return err
	}
	device.DeviceKey = obj.DeviceKey
	return nil
}

func (uc *BarkUseCase) Pull(ctx context.Context, device *entity.BarkDevice, offset, limit int) ([]*entity.BarkHistory, error) {
	return nil, nil
}
