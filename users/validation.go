package users

import (
	"awesomeProject/domain"
	"awesomeProject/shared"
	"context"
	"github.com/go-playground/validator/v10"
)

type usersServiceValidationMiddleware struct {
	validation *validator.Validate
	next       UsersService
}

func NewUsersServiceValidationMiddleware(validation *validator.Validate) UsersServiceMiddleware {
	return func(next UsersService) UsersService {
		return &usersServiceValidationMiddleware{
			validation: validation,
			next:       next,
		}
	}
}

func (mw *usersServiceValidationMiddleware) Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (*domain.User, error) {
	if err := mw.validation.StructCtx(ctx, request); err != nil {
		return &domain.User{}, shared.MakeValidationError(err)
	}

	return mw.next.Register(ctx, request)
}

func (mw *usersServiceValidationMiddleware) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (*domain.User, error) {
	if err := mw.validation.StructCtx(ctx, request); err != nil {
		return &domain.User{}, shared.MakeValidationError(err)
	}

	return mw.next.Login(ctx, request)
}
