package repository

import (
	"context"

	"github.com/olezhek28/clean-architecture/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, userUUID string, info *model.UserInfo) error
	Get(ctx context.Context, uuid string) (*model.User, error)
}
