package usecase_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"aio/internal/entity"
	"aio/internal/usecase"
	"testing"
)

//go:generate mockgen -source=interfaces.go -package=usecase_test -destination=mocks_test.go

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

var (
	errEmptyDevice = fmt.Errorf("empty device")
)

func bark(t *testing.T) (*usecase.BarkUseCase, *MockBarkRepo) {
	t.Helper()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := NewMockBarkRepo(ctl)
	b := usecase.NewBark(repo, nil, nil)
	return b, repo
}

func TestBark(t *testing.T) {
	b, r := bark(t)
	tests := []test{
		{
			name: "empty",
			mock: func() {
				r.EXPECT().Store(context.Background(), &entity.BarkDevice{}).Return(errEmptyDevice)
			},
			res: &entity.BarkDevice{},
			err: errEmptyDevice,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := b.Register(context.Background(), &entity.BarkDevice{})
			require.ErrorIs(t, err, errEmptyDevice)
		})
	}
}
