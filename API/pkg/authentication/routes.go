package authentication

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
auth routes:
POST /auth/login
POST /auth/refresh
POST /auth/register
*/

func Routes(configuration *config.Config) chi.Router {
	UserConfig := New(configuration)
	router := chi.NewRouter()
	router.Post("/login", UserConfig.Login)
	router.Post("/refresh", UserConfig.Refresh)
	router.Post("/register", UserConfig.Register)
	return router
}
