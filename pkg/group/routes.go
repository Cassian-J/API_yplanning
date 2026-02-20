package group

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
group routes:
POST /groups - Create a new group
GET /groups - Get all groups (for testing purposes, remove later)
GET /groups/{id} - Get a group by ID
GET /groups/creator/{id} - Get groups by creator ID
PUT /groups/{id} - Update a group by ID
DELETE /groups/{id} - Delete a group by ID
*/

func Routes(config *config.Config) chi.Router {
	GroupConfig := NewGroupConfig(config)
	router := chi.NewRouter()
	router.Post("/", GroupConfig.CreateGroup)
	router.Get("/groups", GroupConfig.GetAllGroups) // FOR TESTING PURPOSES ONLY, REMOVE LATER
	router.Get("/{id}", GroupConfig.GetGroupByID)
	router.Get("/creator/{id}", GroupConfig.GetGroupByCreatorID)
	router.Put("/{id}", GroupConfig.Updategroup)
	router.Delete("/{id}", GroupConfig.DeleteGroupHandler)
	return router
}
