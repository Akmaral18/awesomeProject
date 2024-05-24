package users

import (
	"awesomeProject/domain"
	"context"
	"golang.org/x/exp/slog"
	"time"
)

type usersServiceLoggingMiddleware struct {
	logger *slog.Logger
	next   UsersService
}

func NewUsersServiceLoggingMiddleware(logger *slog.Logger) UsersServiceMiddleware {
	return func(next UsersService) UsersService {
		return &usersServiceLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (mw *usersServiceLoggingMiddleware) Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (user *domain.User, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoContext(ctx,
			"Register",
			"request_time", time.Since(begin),
			"error", err,
			"user_created", user != nil,
		)
	}(time.Now())

	mw.logger.InfoContext(ctx,
		"Register",
		"request", request,
	)

	return mw.next.Register(ctx, request)
}

func (mw *usersServiceLoggingMiddleware) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (user *domain.User, err error) {
	defer func(begin time.Time) {
		mw.logger.InfoContext(ctx,
			"Login",
			"request_time", time.Since(begin),
			"error", err,
			"user_verified", user != nil,
		)
	}(time.Now())

	mw.logger.InfoContext(ctx,
		"Login",
		"request", request,
	)

	return mw.next.Login(ctx, request)
}
