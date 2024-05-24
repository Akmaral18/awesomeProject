package awesomeProject

import (
	"awesomeProject/repositories"
	"awesomeProject/users"
	"awesomeProject/utilities"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
)

type ServiceContainer struct {
	UsersService users.UsersService
}

// NewServiceContainer builds the downstream services used throughout the application.
func NewServiceContainer(logger *slog.Logger, db *sql.DB) *ServiceContainer {
	validation := validator.New()
	usersRepository := repositories.NewUsersRepository(db)

	var usersService users.UsersService
	{
		tokenService := utilities.NewTokenService()
		securityService := utilities.NewSecurityService()
		usersService = users.NewUsersService(logger, usersRepository, tokenService, securityService)
		usersService = users.NewUsersServiceLoggingMiddleware(logger)(usersService)
		usersService = users.NewUsersServiceValidationMiddleware(validation)(usersService)
	}

	return &ServiceContainer{
		UsersService: usersService,
	}
}
