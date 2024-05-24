package awesomeProject

import (
	"awesomeProject/shared"
	"awesomeProject/users"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger, container *ServiceContainer) http.Handler {
	router := mux.NewRouter()
	router.Use(shared.CorsPolicy)
	router.Use(shared.JsonContentType)

	router = users.MakeUserRoutes(logger, router, container.UsersService)

	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()
	apiRouter.PathPrefix("/").Handler(router)
	router.PathPrefix("/api").Handler(apiRouter)

	return router
}
