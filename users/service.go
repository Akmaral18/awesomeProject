package users

import (
	"awesomeProject/domain"
	"awesomeProject/repositories"
	"awesomeProject/utilities"
	"context"
	"golang.org/x/exp/slog"
)

type (
	UsersService interface {
		Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (*domain.User, error)
		Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (*domain.User, error)
	}

	userService struct {
		logger          *slog.Logger
		repository      repositories.UsersRepository
		tokenService    utilities.TokenService
		securityService utilities.SecurityService
	}

	UsersServiceMiddleware func(service UsersService) UsersService
)

func NewUsersService(logger *slog.Logger, repository repositories.UsersRepository, tokenService utilities.TokenService, securityService utilities.SecurityService) UsersService {
	return &userService{
		logger:          logger,
		repository:      repository,
		tokenService:    tokenService,
		securityService: securityService,
	}
}
