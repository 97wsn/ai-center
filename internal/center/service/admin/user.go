package admin

import (
	"context"

	center "github.com/97wsn/ai-center/api/center/admin/v1"
)

var _ center.UserHTTPServer = (*User)(nil)

type User struct {
}

func (u User) Login(ctx context.Context, request *center.UserLoginRequest) (*center.UserLoginResponse, error) {
	//TODO implement me
	panic("implement me")
}
