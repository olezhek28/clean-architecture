package note

import (
	"github.com/olezhek28/clean-architecture/internal/repository"
	def "github.com/olezhek28/clean-architecture/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *service {
	return &service{
		userRepository: userRepository,
	}
}
