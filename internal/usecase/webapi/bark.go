package webapi

import (
	"aio/internal/entity"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/samber/lo"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"github.com/sideshow/apns2/token"
	"golang.org/x/net/http2"
	"net/http"

	"strings"
	"time"
)

var _ usecase.BarkWebAPI = (*BarkAPNsAPI)(nil)

type BarkAPNsAPI struct {
	Client *apns2.Client
	l      logger.Interface
}

func NewBarkAPNs(l logger.Interface) (usecase.BarkWebAPI, error) {
	key, err := token.AuthKeyFromBytes([]byte(apnsPrivateKey))
	if err != nil {
		return nil, err
	}

	CAs, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	lo.ForEach[string](apnsCAs, func(t string, _ int) {
		CAs.AppendCertsFromPEM([]byte(t))
	})

	client := &apns2.Client{
		Token: &token.Token{
			AuthKey: key,
			KeyID:   keyID,
			TeamID:  teamID,
		},
		HTTPClient: &http.Client{
			Transport: &http2.Transport{
				DialTLS: apns2.DialTLS,
				TLSClientConfig: &tls.Config{
					RootCAs: CAs,
				},
			},
			Timeout: apns2.HTTPClientTimeout,
		},
		Host: apns2.HostProduction,
	}
	return &BarkAPNsAPI{client, l}, nil
}

func (api *BarkAPNsAPI) Push(ctx context.Context, msg *entity.APNsMessage) error {
	pl := payload.NewPayload().
		AlertTitle(msg.Title).
		AlertBody(msg.Content).
		Sound(msg.Sound).
		Category(msg.Category)

	group, ok := msg.Params["group"]
	if ok {
		pl = pl.ThreadID(group)
	}

	for k, v := range msg.Params {
		pl.Custom(strings.ToLower(k), fmt.Sprintf("%v", v))
	}

	rsp, err := api.Client.Push(&apns2.Notification{
		DeviceToken: msg.DeviceToken,
		Topic:       topic,
		Payload:     pl.MutableContent(),
		Expiration:  time.Now().Add(time.Hour * 24),
	})

	if err != nil {
		return err
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("bark push %v error", rsp.Reason)
	}
	return nil
}
