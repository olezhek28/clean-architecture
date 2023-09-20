package service

import (
	"context"

	"github.com/olezhek28/clean-architecture/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (string, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
}
