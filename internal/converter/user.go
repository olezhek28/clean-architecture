package converter

import (
	"github.com/olezhek28/clean-architecture/internal/model"
	desc "github.com/olezhek28/clean-architecture/pkg/user_v1"
)

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
