package user

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
User routes:
GET /user/users - Get all users (for testing purposes only)
GET /user/{id} - Get a user by ID
GET /user/ - Get a user by email
PUT /user/{id} - Update a user by ID
DELETE /user/{id} - Delete a user by ID
*/

func Routes(config *config.Config) chi.Router {
	UserConfig := NewUserConfig(config)
	router := chi.NewRouter()
	router.Get("/users", UserConfig.GetAllUsers) // FOR TESTING PURPOSES ONLY
	router.Get("/{id}", UserConfig.GetUserByID)
	router.Get("/", UserConfig.GetUser)
	router.Put("/{id}", UserConfig.UpdateUser)
	router.Delete("/{id}", UserConfig.DeleteUser)
	return router
}
