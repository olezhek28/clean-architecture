package user

import (
	"context"
	"sync"
	"time"

	"github.com/olezhek28/clean-architecture/internal/model"
	def "github.com/olezhek28/clean-architecture/internal/repository"
	"github.com/olezhek28/clean-architecture/internal/repository/user/converter"
	repoModel "github.com/olezhek28/clean-architecture/internal/repository/user/model"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.User
	m    sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.User),
	}
}

func (r *repository) Create(_ context.Context, userUUID string, info *model.UserInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[userUUID] = &repoModel.User{
		UUID:      userUUID,
		Info:      converter.ToUserInfoFromService(info),
		CreatedAt: time.Now(),
	}

	return nil
}
