package groupuser

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

type GroupUserConfig struct {
	*config.Config
}

func NewGroupUserConfig(cfg *config.Config) *GroupUserConfig {
	return &GroupUserConfig{Config: cfg}
}

// @Summary		Create a new group-user relationship
// @Description	Create a new group-user relationship with the provided user ID, group ID, and optional color ID
// @Tags		group-users
// @Accept		json
// @Produce		json
// @Param		groupUser	body		models.GroupUserRequest	true	"GroupUserRequest"
// @Success		200	{object}	models.GroupUserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/ [post]
func (config *GroupUserConfig) CreateGroupUser(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupUserRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	groupUser := &dbmodel.UserGroup{UserID: req.UserID, GroupID: req.GroupID}
	created, err := config.UserGroupRepository.Create(groupUser)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to create group-user relationship")
		return
	}
	groupUserResponse := &models.GroupUserResponse{UserID: created.UserID, GroupID: created.GroupID, ColorID: created.ColorID}
	render.JSON(w, r, groupUserResponse)
}

// @Summary		Get all group-user relationships
// @Description	Retrieve a list of all group-user relationships
// @Tags		group-users
// @Accept		json
// @Produce		json
// @Success		200	{array}	models.GroupUserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/groups-users/ [get]
func (config *GroupUserConfig) GetAllGroupUsers(w http.ResponseWriter, r *http.Request) {
	groupUsers, err := config.UserGroupRepository.FindAll()
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve group-user relationships")
		return
	}
	groupUserResponse := make([]models.GroupUserResponse, 0)
	for _, groupUser := range groupUsers {
		groupUserResponse = append(groupUserResponse, models.GroupUserResponse{UserID: groupUser.UserID, GroupID: groupUser.GroupID, ColorID: groupUser.ColorID})
	}
	render.JSON(w, r, groupUserResponse)
}

// @Summary		Get groups by user ID
// @Description	Retrieve a list of groups that a user belongs to by the user's ID
// @Tags		groups
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"User ID"
// @Success		200	{array}	models.GroupResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/user/{id} [get]
func (config *GroupUserConfig) GetGroupsByUserID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "User ID must be greater than 0")
		return
	}
	groupUsers, err := config.UserGroupRepository.FindByUserID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve group-user relationships")
		return
	}
	groupResponse := make([]models.GroupResponse, 0)
	for _, groupUser := range groupUsers {
		group, err := config.GroupRepository.FindByID(groupUser.GroupID)
		if err != nil {
			errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve group")
			return
		}
		groupResponse = append(groupResponse, models.GroupResponse{ID: group.ID, Name: group.Name, CreatorID: group.CreatorID})
	}
	render.JSON(w, r, groupResponse)
}

// @Summary		Get users by group ID
// @Description	Retrieve a list of users that belong to a group by the group's ID
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"Group ID"
// @Success		200	{array}	models.UserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/group/{id} [get]
func (config *GroupUserConfig) GetUsersByGroupID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "Group ID must be greater than 0")
		return
	}
	userGroups, err := config.UserGroupRepository.FindByGroupID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve user-group relationships")
		return
	}
	userResponse := make([]models.UserResponse, 0)
	for _, userGroup := range userGroups {
		user, err := config.UserRepository.FindByID(userGroup.UserID)
		if err != nil {
			errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve user")
			return
		}
		userResponse = append(userResponse, models.UserResponse{ID: user.ID, Username: user.Username, Email: user.Email, Name: user.Name, Surname: user.Surname, ColorID: user.ColorID})
	}
	render.JSON(w, r, userResponse)
}

// @Summary		Get group-user relationship by user ID and group ID
// @Description	Retrieve a group-user relationship by the user's ID and the group's ID
// @Tags		group-users
// @Accept		json
// @Produce		json
// @Param		user_id	path		int	true	"User ID"
// @Param		group_id	path		int	true	"Group ID"
// @Success		200	{object}	models.GroupUserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/group/{groupID}/user/{userID} [get]
func (config *GroupUserConfig) GetGroupUserByUserIDAndGroupID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}
	groupID, err := strconv.Atoi(chi.URLParam(r, "group_id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}
	if userID < 1 || groupID < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "User ID and Group ID must be greater than 0")
		return
	}
	groupUser, err := config.UserGroupRepository.FindByUserIDAndGroupID(uint(userID), uint(groupID))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve group-user relationship")
		return
	}
	groupUserResponse := &models.GroupUserResponse{UserID: groupUser.UserID, GroupID: groupUser.GroupID, ColorID: groupUser.ColorID}
	render.JSON(w, r, groupUserResponse)
}

// @Summary		Update group-user color
// @Description	Update the color of a group-user relationship
// @Tags		group-users
// @Accept		json
// @Produce		json
// @Param		groupUser	body		models.GroupUserRequest	true	"GroupUserRequest"
// @Success		200	{object}	map[string]string
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/color [put]
func (config *GroupUserConfig) UpdateGroupUserColor(w http.ResponseWriter, r *http.Request) {
	req := &models.GroupUserRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := config.UserGroupRepository.UpdateColorByUserIDAndGroupID(req.UserID, req.GroupID, req.ColorID)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to update group-user color")
		return
	}
	render.JSON(w, r, map[string]string{"message": "Group-user color updated successfully"})
}

// @Summary		Delete group-user relationship
// @Description	Delete a group-user relationship by user ID and group ID
// @Tags		group-users
// @Accept		json
// @Produce		json
// @Param		groupUser	body		models.GroupUserRequest	true	"GroupUserRequest"
// @Success		200	{object}	map[string]string
// @Failure 	400 {object} models.ErrorResponse
// @Failure 	500 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/group-user/group/{groupID}/user/{userID} [delete]
func (config *GroupUserConfig) DeleteGroupUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}
	groupID, err := strconv.Atoi(chi.URLParam(r, "group_id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid group ID")
		return
	}
	if userID < 1 || groupID < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "User ID and Group ID must be greater than 0")
		return
	}
	err = config.UserGroupRepository.DeleteByUserIDAndGroupID(uint(userID), uint(groupID))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to delete group-user relationship")
		return
	}
	render.JSON(w, r, map[string]string{"message": "Group-user relationship deleted successfully"})
}
