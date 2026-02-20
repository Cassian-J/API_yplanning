package group

import (
	"fmt"
	"net/http"
	"strconv"

	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type GroupConfig struct {
	*config.Config
}

func NewGroupConfig(cfg *config.Config) *GroupConfig {
	return &GroupConfig{Config: cfg}
}

func (config *GroupConfig) CreateGroup(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}
	group := &dbmodel.Group{Name: req.Name, CreatorID: req.CreatorID}
	created, err := config.GroupRepository.Create(group)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create group"})
		return
	}
	groupResponse := &models.GroupResponse{ID: created.ID, Name: created.Name, CreatorID: created.CreatorID}
	render.JSON(w, r, groupResponse)
}

func (config *GroupConfig) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := config.GroupRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve groups", http.StatusInternalServerError)
		return
	}

	GroupResponse := make([]models.GroupResponse, 0)
	for _, group := range groups {
		GroupResponse = append(GroupResponse, models.GroupResponse{
			ID:        group.ID,
			Name:      group.Name,
			CreatorID: group.CreatorID,
		})
	}
	render.JSON(w, r, GroupResponse)
}

func (config *GroupConfig) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	group, err := config.GroupRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve group"})
		return
	}
	groupResponse := &models.GroupResponse{ID: group.ID, Name: group.Name, CreatorID: group.CreatorID}
	render.JSON(w, r, groupResponse)
}

func (config *GroupConfig) GetGroupByCreatorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	group, err := config.GroupRepository.FindByCreatorID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve group"})
		return
	}
	groupResponse := &models.GroupResponse{ID: group.ID, Name: group.Name, CreatorID: group.CreatorID}
	render.JSON(w, r, groupResponse)
}

func (config *GroupConfig) Updategroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}

	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	group := &dbmodel.Group{Name: req.Name, CreatorID: req.CreatorID}
	updated, err := config.GroupRepository.UpdateByID(uint(id), group)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update group"})
		return
	}

	groupResponse := &models.GroupResponse{ID: uint(id), Name: updated.Name, CreatorID: updated.CreatorID}
	render.JSON(w, r, groupResponse)
}

func (config *GroupConfig) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.GroupRepository.DeleteByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete group"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
