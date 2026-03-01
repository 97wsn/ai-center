package rpc

import (
	"context"

	"app/internal/center/config"

	center "github.com/97wsn/ai-center/api/center/rpc/v1"
)

var _ center.UserServer = (*User)(nil)

type User struct {
	Conf *config.Config
}

func (u User) Login(ctx context.Context, request *center.UserLoginRequest) (*center.UserLoginResponse, error) {
	return &center.UserLoginResponse{
		UserId: 1234,
	}, nil
}
