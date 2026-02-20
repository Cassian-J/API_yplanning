package group

import (
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

// @Summary		Create a new group
// @Description	Create a new group with the provided name and creator ID
// @Tags		groups
// @Accept		json
// @Produce		json
// @Param		request	body	models.GroupRequest	true	"Group creation data"
// @Success		200	{object}	models.GroupResponse
// @Failure 	400 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/group/ [post]
func (config *GroupConfig) CreateGroup(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	group := &dbmodel.Group{Name: req.Name, CreatorID: req.CreatorID}
	created, err := config.GroupRepository.Create(group)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}
	groupResponse := &models.GroupResponse{ID: created.ID, Name: created.Name, CreatorID: created.CreatorID}
	render.JSON(w, r, groupResponse)
}

// @Summary		Get all groups
// @Description	Retrieve a list of all groups
// @Tags		groups
// @Produce		json
// @Success		200	{array}	models.GroupResponse
// @Failure 	500 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/group/groups [get]
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

// @Summary		Get group by ID
// @Description	Retrieve a group by its ID
// @Tags		groups
// @Produce		json
// @Param		id	path	int	true	"Group ID"
// @Success		200	{object}	models.GroupResponse
// @Failure 	400 {object} 	http.Error
// @Failure 	500 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/group/{id} [get]
func (config *GroupConfig) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}

	group, err := config.GroupRepository.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve group", http.StatusInternalServerError)
		return
	}
	groupResponse := &models.GroupResponse{ID: group.ID, Name: group.Name, CreatorID: group.CreatorID}
	render.JSON(w, r, groupResponse)
}

// @Summary		Get group by creator ID
// @Description	Retrieve a group by its creator ID
// @Tags		groups
// @Produce		json
// @Param		id	path	int	true	"Creator ID"
// @Success		200	{object}	models.GroupResponse
// @Failure 	400 {object} 	http.Error
// @Failure 	500 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/group/creator/{id} [get]
func (config *GroupConfig) GetGroupByCreatorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}

	group, err := config.GroupRepository.FindByCreatorID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve group", http.StatusInternalServerError)
		return
	}
	groupResponse := &models.GroupResponse{ID: group.ID, Name: group.Name, CreatorID: group.CreatorID}
	render.JSON(w, r, groupResponse)
}

// @Summary		Update a group
// @Description	Update a group by its ID with the provided name and creator ID
// @Tags		groups
// @Accept		json
// @Produce		json
// @Param		id	path	int	true	"Group ID"
// @Param		request	body	models.GroupRequest	true	"Group update data"
// @Success		200	{object}	models.GroupResponse
// @Failure 	400 {object} 	http.Error
// @Failure 	500 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/group/{id} [put]
func (config *GroupConfig) Updategroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}

	group := &dbmodel.Group{Name: req.Name, CreatorID: req.CreatorID}
	updated, err := config.GroupRepository.UpdateByID(uint(id), group)
	if err != nil {
		http.Error(w, "Failed to update group", http.StatusInternalServerError)
		return
	}

	groupResponse := &models.GroupResponse{ID: uint(id), Name: updated.Name, CreatorID: updated.CreatorID}
	render.JSON(w, r, groupResponse)
}

// @Summary		Delete a group
// @Description	Delete a group by its ID
// @Tags		groups
// @Produce		json
// @Param		id	path	int	true	"Group ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Failure 	400 {object} 	http.Error
// @Failure 	500 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/group/{id} [delete]
func (config *GroupConfig) DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	err = config.GroupRepository.DeleteByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete group", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
