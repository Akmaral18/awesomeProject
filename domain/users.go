package domain

import "github.com/gofrs/uuid"

type (
	User struct {
		ID       uuid.UUID `json:"-"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Token    string    `json:"token"`
	}

	RegisterUserRequest struct {
		Username string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	LoginUserRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	AuthenticationRequest[T RegisterUserRequest | LoginUserRequest] struct {
		User *T `json:"user" validate:"required"`
	}

	AuthenticationResponse struct {
		User *User `json:"user"`
	}
)
