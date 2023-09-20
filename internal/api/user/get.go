package user

import (
	"context"

	"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
