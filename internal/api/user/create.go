package user

import (
	"context"
	"log"

	"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	uuid, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with uuid: %s", uuid)

	return &desc.CreateResponse{
		Uuid: uuid,
	}, nil
}
