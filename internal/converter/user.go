package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/olezhek28/clean-architecture/internal/model"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return &desc.User{
		Uuid:      user.UUID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Age:       info.Age,
	}
}

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Age:       info.Age,
	}
}
