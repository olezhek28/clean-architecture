package user

import (
	"context"
	"log"

	"github.com/olezhek28/clean-architecture/internal/model"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, uuid)
	if err != nil {
		log.Printf("ошибка получения пользователя: %v\n", err)
		return nil, err
	}
	if user == nil {
		log.Printf("пользователь с uuid %s не найден\n", uuid)
		return nil, model.ErrorUserNotFound
	}

	return user, nil
}
