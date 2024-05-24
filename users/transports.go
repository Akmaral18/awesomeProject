package users

import (
	"awesomeProject/domain"
	"awesomeProject/shared"
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
)

func MakeUserRoutes(logger *slog.Logger, router *mux.Router, service UsersService) *mux.Router {
	registerUserHandler := httptransport.NewServer(
		makeRegisterUserEndpoint(service),
		decodeRegisterUserRequest,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)
	loginUserHandler := httptransport.NewServer(
		makeLoginUserEndpoint(service),
		decodeLoginUserRequest,
		shared.EncodeSuccessfulOkResponse,
		shared.HandlerOptions(logger)...,
	)

	usersRouter := router.PathPrefix("/users").Subrouter()
	usersRouter.Methods("POST").Path("signup").HandlerFunc(registerUserHandler.ServeHTTP)
	usersRouter.Methods("POST").Path("/login").HandlerFunc(loginUserHandler.ServeHTTP)

	return router
}

func decodeRegisterUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeUserRequest[domain.RegisterUserRequest](ctx, r)
}

func decodeLoginUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return decodeUserRequest[domain.LoginUserRequest](ctx, r)
}

func decodeUserRequest[T domain.RegisterUserRequest | domain.LoginUserRequest](_ context.Context, r *http.Request) (interface{}, error) {
	var request domain.AuthenticationRequest[T]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, shared.ErrorWithContext("error while attempting to decode the authentication request", shared.ErrInvalidRequestBody)
	}

	return request, nil
}
