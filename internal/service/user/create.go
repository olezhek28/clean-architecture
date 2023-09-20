package user

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/olezhek28/clean-architecture/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (string, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("ошибка генерации user UUID: %v\n", err)
		return "", err
	}

	err = s.userRepository.Create(ctx, userUUID.String(), info)
	if err != nil {
		log.Printf("ошибка создания пользователя: %v\n", err)
		return "", err
	}

	return userUUID.String(), nil
}
