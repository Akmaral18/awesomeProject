package users

import (
	"awesomeProject/domain"
	"awesomeProject/shared"
	"awesomeProject/utilities"
	"context"
	"net/http"
)

func (us *userService) Register(ctx context.Context, request domain.AuthenticationRequest[domain.RegisterUserRequest]) (*domain.User, error) {
	us.logger.InfoContext(ctx, "attempting to register new user, checking for existing users", "username", request.User.Username, "email", request.User.Email)
	existingUsers, err := us.repository.SearchUsers(ctx, request.User.Username, request.User.Email)

	if len(existingUsers) != 0 {
		// Technically, there could be at least two entries here - only log out the first one
		existingUser := existingUsers[0]
		us.logger.ErrorContext(ctx, "username or email is taken", "username", existingUser.Username, "email", existingUser.Email)
		return &domain.User{}, shared.MakeApiErrorWithStatus(http.StatusConflict, shared.ErrUsernameOrEmailTaken)
	} else if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting to query for existing users", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	us.logger.InfoContext(ctx, "no user clashes found, hashing user password", "username", request.User.Username, "email", request.User.Email)
	hashedPassword, err := us.securityService.HashPassword(request.User.Password)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting to hash user password", "err", err, "username", request.User.Username, "email", request.User.Email)
		return &domain.User{}, shared.MakeApiError(err)
	}

	us.logger.InfoContext(ctx, "password successfully hashed, creating user", "username", request.User.Username, "email", request.User.Email)
	createdUser, err := us.repository.CreateUser(ctx, request.User.Username, request.User.Email, hashedPassword)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting create user", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	us.tokenService = utilities.NewTokenService()
	us.logger.InfoContext(ctx, "user successfully created, generating token", "username", createdUser.Username, "email", createdUser.Email, "user_id", createdUser.ID)
	token, err := us.tokenService.GenerateUserToken(createdUser.ID, createdUser.Email)
	if err != nil {
		us.logger.ErrorContext(ctx, "error while attempting generate user token", "err", err)
		return &domain.User{}, shared.MakeApiError(err)
	}

	us.logger.InfoContext(ctx, "token successfully generated", "username", createdUser.Username, "email", createdUser.Email, "user_id", createdUser.ID)

	return createdUser.ToUser(token), nil
}
