package admin

import (
	"github.com/google/wire"
)

type Service struct {
	User *User
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Service), "*"),
	wire.Struct(new(User), "*"),
)
