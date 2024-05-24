package users

import (
	"awesomeProject/domain"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeRegisterUserEndpoint(service UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		registrationRequest := request.(domain.AuthenticationRequest[domain.RegisterUserRequest])
		createdUser, err := service.Register(ctx, registrationRequest)
		if err != nil {
			return nil, err
		}

		return &domain.AuthenticationResponse{
			User: createdUser,
		}, nil
	}
}

func makeLoginUserEndpoint(service UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		loginRequest := request.(domain.AuthenticationRequest[domain.LoginUserRequest])
		verifiedUser, err := service.Login(ctx, loginRequest)
		if err != nil {
			return nil, err
		}

		return &domain.AuthenticationResponse{
			User: verifiedUser,
		}, nil
	}
}
