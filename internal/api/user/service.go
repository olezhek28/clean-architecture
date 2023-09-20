package user

import (
	"github.com/olezhek28/clean-architecture/internal/service"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
