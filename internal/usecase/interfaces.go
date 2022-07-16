package usecase

import (
	"aio/ent"
	"aio/internal/entity"
	"aio/internal/pb"
	"context"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase

type (
	Bark interface {
		Push(context.Context, string, *entity.APNsMessage) error
		Register(context.Context, *entity.BarkDevice) error
		Pull(ctx context.Context, key string, offset, limit int) ([]*pb.BarkHistory, error)
	}

	BarkRepo interface {
		Store(context.Context, *entity.BarkDevice) error
		Get(context.Context, *entity.BarkDevice) (*entity.BarkDevice, error)
		SaveMessage(ctx context.Context, device *entity.BarkDevice, message *entity.APNsMessage) error
		FetchHistory(ctx context.Context, device *entity.BarkDevice, offset, limit int) ([]*pb.BarkHistory, error)
	}

	BarkWebAPI interface {
		Push(context.Context, *entity.APNsMessage) error
	}
)

type (
	PushDeer interface {
		ValidateToken(ctx context.Context, token string) error

		Register(ctx context.Context, appleID, email, name string) (string, error)
		GetUser(ctx context.Context, token string) (*entity.User, error)

		RegisterDevice(ctx context.Context, info *entity.RegInfo) ([]*entity.Device, error)
		GetAllDevice(ctx context.Context, token string) ([]*entity.Device, error)
		RenameDevice(ctx context.Context, id, name string) error
		RemoveDevice(ctx context.Context, id string) error

		PushMessage(ctx context.Context, key string, msg *entity.Message) error
		GetMessage(ctx context.Context, token string, offset, limit uint64) ([]*entity.Message, error)
		RemoveMessage(ctx context.Context, token, msgID string) error

		GenNewKey(ctx context.Context, token, name string) ([]*entity.PushKey, error)
		RenameKey(ctx context.Context, token, id, name string) error
		RegenKey(ctx context.Context, token, id string) error
		GetAllKey(ctx context.Context, token string) ([]*entity.PushKey, error)
		RemoveKey(ctx context.Context, token, id string) error

		Upgrade(ctx context.Context, deviceID string, w http.ResponseWriter, r *http.Request, h http.Header) error
	}

	PushDeerRepo interface {
		GetUserID(ctx context.Context, token string) (string, error)
		StoreUser(ctx context.Context, appleID, email, name, uuid string) error
		GetUser(ctx context.Context, uuid, appleID string) (*entity.User, error)

		StoreDevice(context.Context, *entity.Device) error
		GetDevice(context.Context, string) (*entity.Device, error)
		GetAllDevice(context.Context, string) ([]*entity.Device, error)
		UpdateDeviceName(ctx context.Context, deviceID, name string) error
		RemoveDevice(context.Context, string) error

		StorePushKey(context.Context, *entity.PushKey) error
		GetPushKey(context.Context, int64, string, string) (*entity.PushKey, error)
		GetAllPushKey(context.Context, string) ([]*entity.PushKey, error)
		UpdatePushKey(context.Context, *entity.PushKey) error
		RemovePushKey(context.Context, string) error

		StoreMessage(context.Context, *entity.Message) error
		GetMessage(ctx context.Context, userID string, offset, count uint64) ([]*entity.Message, error)
		RemoveMessage(context.Context, string) error
	}

	PushDeerWebAPI interface {
		Push(context.Context, []*entity.Device, *entity.Message) error
		Register(ctx context.Context, key string, c *websocket.Conn)
	}
)

type (
	Extension interface {
		Register(ctx context.Context, name, extensionID string) (*ent.ExtensionClient, error)
		Pull(ctx context.Context, uid uuid.UUID) (*ent.ExtensionClient, error)
		Add(ctx context.Context, uid uuid.UUID, group ...*pb.Group) error
		RemoveGroup(ctx context.Context, uid, groupUid uuid.UUID) error
		RemoveTab(ctx context.Context, uid, groupUid, tabUid uuid.UUID) error

		SwapTab(ctx context.Context, uid, firstGroupUid, firstGroupTabUid, secondGroupUid, secondGroupTabUid uuid.UUID) error
		//MoveTab(ctx context.Context, uid, fromGroup, toGroup, tabUid uuid.UUID) error

		Connect(ctx context.Context, uid uuid.UUID, wsConn *websocket.Conn) error
	}
)
