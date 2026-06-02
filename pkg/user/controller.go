package user

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

type UserConfig struct {
	*config.Config
}

func NewUserConfig(cfg *config.Config) *UserConfig {
	return &UserConfig{Config: cfg}
}

// @Summary		Get all users
// @Description	Retrieve a list of all users
// @Tags		users
// @Accept		json
// @Produce		json
// @Success		200	{array}		models.UserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router		/user/users [get]
func (config *UserConfig) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := config.UserRepository.FindAll()
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	UserResponse := make([]models.UserResponse, 0)
	for _, user := range users {
		UserResponse = append(UserResponse, models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Name:     user.Name,
			Surname:  user.Surname,
			ColorID:  user.ColorID,
		})
	}
	render.JSON(w, r, UserResponse)
}

// @Summary		Get user by ID
// @Description	Retrieve a user by its ID
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		id	path		int	true	"User ID"
// @Success		200	{object}	models.UserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router		/user/{id} [get]
func (config *UserConfig) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "User ID must be greater than 0")
		return
	}

	user, err := config.UserRepository.FindByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}
	userResponse := &models.UserResponse{ID: user.ID, Email: user.Email, Username: user.Username}
	render.JSON(w, r, userResponse)
}

// @Summary		Get user by email or username
// @Description	Retrieve a user by its username or email
// @Tags		users
// @Accept		json
// @Produce		json
// @Param username query string false "Username of the user"
// @Param email    query string false "Email of the user"
// @Success		200	{object}	models.UserResponse
// @Failure 	400 {object} 	models.ErrorResponse
// @Security 	BearerAuth
// @Router		/user/ [post]
func (config *UserConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	req := &models.GetUserRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request parameters")
		return
	}
	username := req.Username
	email := req.Email

	var user *dbmodel.User
	var err error
	if username != "" {
		user, err = config.UserRepository.FindByUsername(username)
	} else {
		user, err = config.UserRepository.FindByEmail(email)
	}

	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	userResponse := &models.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		ColorID:  user.ColorID,
	}

	render.JSON(w, r, userResponse)
}

// @Summary		Update a user
// @Description	Update a user by its ID
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		id		path	int					true	"User ID"
// @Param		request	body	models.UserRequest	true	"Updated user data"
// @Success		200	{object}	models.UserResponse
// @Failure 	400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router		/user/{id} [put]
func (config *UserConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}

	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request body")
		return
	}

	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "User ID must be greater than 0")
		return
	}

	user := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username, Name: req.Name, Surname: req.Surname, ColorID: req.ColorID}
	updated, err := config.UserRepository.UpdateByID(uint(id), user)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to update user")
		return
	}

	userResponse := &models.UserResponse{ID: uint(id), Email: updated.Email, Username: updated.Username, Name: updated.Name, Surname: updated.Surname, ColorID: updated.ColorID}
	render.JSON(w, r, userResponse)
}

// @Summary		Delete a user
// @Description	Delete a user by its ID
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		id	path		int		true	"User ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Failure 	400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router		/user/{id} [delete]
func (config *UserConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "User ID must be greater than 0")
		return
	}
	err = config.UserRepository.DeleteByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	render.JSON(w, r, "Successfully deleted entry")
}
