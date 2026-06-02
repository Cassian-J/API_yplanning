package groupuser

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
group-user routes:
POST /group-user/ - Create a new group-user relationship
GET /group-user/groups-users - Get all group-user relationships (FOR TESTING PURPOSES ONLY, REMOVE LATER)
GET /group-user/user/{id} - Get all groups for a specific user
GET /group-user/group/{id} - Get all users for a specific group
GET /group-user/group/{groupID}/user/{userID} - Get a specific group-user relationship by user ID and group ID
PUT /group-user/group/{groupID}/user/{userID} - Update the color of a specific group-user relationship
DELETE /group-user/group/{groupID}/user/{userID} - Delete a specific group-user relationship
*/

func Routes(config *config.Config) chi.Router {
	GroupUserConfig := NewGroupUserConfig(config)
	router := chi.NewRouter()
	router.Post("/", GroupUserConfig.CreateGroupUser)
	router.Get("/groups-users", GroupUserConfig.GetAllGroupUsers) // FOR TESTING PURPOSES ONLY, REMOVE LATER
	router.Get("/user/{id}", GroupUserConfig.GetGroupsByUserID)
	router.Get("/group/{id}", GroupUserConfig.GetUsersByGroupID)
	router.Get("/group/{groupID}/user/{userID}", GroupUserConfig.GetGroupUserByUserIDAndGroupID)
	router.Put("/color", GroupUserConfig.UpdateGroupUserColor)
	router.Delete("/", GroupUserConfig.DeleteGroupUser)
	return router
}
