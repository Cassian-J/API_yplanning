package group

import (
	"net/http"
	"strconv"

	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/errors"
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
// @Failure 	400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group/ [post]
func (config *GroupConfig) CreateGroup(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if req.CreatorID == 0 {
		errors.RenderError(w, r, http.StatusBadRequest, "Creator ID is required")
		return
	}

	group := &dbmodel.Group{Name: req.Name, CreatorID: req.CreatorID}
	created, err := config.GroupRepository.Create(group)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to create group")
		return
	}

	userGroup := &dbmodel.UserGroup{UserID: req.CreatorID, GroupID: created.ID}
	if _, err := config.UserGroupRepository.Create(userGroup); err != nil {
		_ = config.GroupRepository.DeleteByID(created.ID)
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to link creator to group")
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
// @Failure 	400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group/groups [get]
func (config *GroupConfig) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := config.GroupRepository.FindAll()
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve groups")
		return
	}

	GroupResponse := make([]models.GroupResponse, 0)
	for _, group := range groups {
		GroupResponse = append(GroupResponse, *models.ToGroupResponse(&group))
	}
	render.JSON(w, r, GroupResponse)
}

// @Summary		Get group by ID
// @Description	Retrieve a group by its ID
// @Tags		groups
// @Produce		json
// @Param		id	path	int	true	"Group ID"
// @Success		200	{object}	models.GroupResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group/{id} [get]
func (config *GroupConfig) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}

	group, err := config.GroupRepository.FindByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve group")
		return
	}
	groupResponse := models.ToGroupResponse(group)
	render.JSON(w, r, groupResponse)
}

// @Summary		Get group by creator ID
// @Description	Retrieve a group by its creator ID
// @Tags		groups
// @Produce		json
// @Param		id	path	int	true	"Creator ID"
// @Success		200	{object}	models.GroupResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group/creator/{id} [get]
func (config *GroupConfig) GetGroupByCreatorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}

	group, err := config.GroupRepository.FindByCreatorID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve group")
		return
	}
	groupResponse := models.ToGroupResponse(group)
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
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group/{id} [put]
func (config *GroupConfig) Updategroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}

	req := &models.GroupRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}

	group := &dbmodel.Group{Name: req.Name, CreatorID: req.CreatorID}
	updated, err := config.GroupRepository.UpdateByID(uint(id), group)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to update group")
		return
	}

	groupResponse := models.ToGroupResponse(updated)
	render.JSON(w, r, groupResponse)
}

// @Summary		Delete a group
// @Description	Delete a group by its ID
// @Tags		groups
// @Produce		json
// @Param		id	path	int	true	"Group ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group/{id} [delete]
func (config *GroupConfig) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}
	err = config.GroupRepository.DeleteByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to delete group")
		return
	}
	render.JSON(w, r, "Successfully deleted entry")
}
