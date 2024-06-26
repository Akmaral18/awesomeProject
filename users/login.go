package users

import (
	"awesomeProject/domain"
	"awesomeProject/shared"
	"context"
	"net/http"
)

func (us *userService) Login(ctx context.Context, request domain.AuthenticationRequest[domain.LoginUserRequest]) (*domain.User, error) {
	us.logger.InfoContext(ctx, "attempting to login user, checking for existing user", "email", request.User.Email)
	existingUsers, err := us.repository.SearchUsers(ctx, "", request.User.Email)

	if len(existingUsers) == 0 {
		us.logger.ErrorContext(ctx, "user was not found", "email", request.User.Email)
		return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusNotFound, shared.ErrUserNotFound)
	} else if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting to query for existing users", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	existingUser := existingUsers[0]
	us.logger.InfoContext(ctx, "user found, attempting to verify password", "username", existingUser.Username, "email", existingUser.Email)
	validLoginAttempt := us.securityService.IsValidPassword(existingUser.Password, request.User.Password)

	if !validLoginAttempt {
		us.logger.ErrorContext(ctx, "unauthorized attempt to login", "username", existingUser.Username, "email", existingUser.Email)
		return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusBadRequest, shared.ErrInvalidLoginAttempt)
	}

	us.logger.InfoContext(ctx, "user successfully verified, generating token", "username", existingUser.Username, "email", existingUser.Email, "user_id", existingUser.ID.String())
	token, err := us.tokenService.GenerateUserToken(existingUser.ID, existingUser.Email)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting generate user token", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	us.logger.InfoContext(ctx, "token successfully generated", "username", existingUser.Username, "email", existingUser.Email, "user_id", existingUser.ID.String())

	return existingUser.ToUser(token), nil
}
