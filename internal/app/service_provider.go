package app

import (
	"log"

	"github.com/olezhek28/clean-architecture/internal/api/user"
	"github.com/olezhek28/clean-architecture/internal/config"
	"github.com/olezhek28/clean-architecture/internal/repository"
	userRepository "github.com/olezhek28/clean-architecture/internal/repository/user"
	"github.com/olezhek28/clean-architecture/internal/service"
	userService "github.com/olezhek28/clean-architecture/internal/service/user"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository()
	}

	return s.userRepository
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService())
	}

	return s.userImpl
}
