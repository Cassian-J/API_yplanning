package user

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
// @Failure 	400 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/user/users [get]
func (config *UserConfig) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := config.UserRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
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
// @Failure 	400 {object}	http.Error
// @Security 	BearerAuth
// @Router		/user/{id} [get]
func (config *UserConfig) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "User ID must be greater than 0", http.StatusBadRequest)
		return
	}

	user, err := config.UserRepository.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
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
// @Success		200	{object}	models.UserResponse
// @Failure 	400 {object}	http.Error
// @Security 	BearerAuth
// @Router		/user/ [get]
func (config *UserConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	var req models.GetUserRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user *dbmodel.User
	var err error

	if req.Username != "" {
		user, err = config.UserRepository.FindByUsername(req.Username)
	} else if req.Email != "" {
		user, err = config.UserRepository.FindByEmail(req.Email)
	} else {
		http.Error(w, "Either username or email must be provided", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
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
// @Failure 	400 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/user/{id} [put]
func (config *UserConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if id < 1 {
		http.Error(w, "User ID must be greater than 0", http.StatusBadRequest)
		return
	}

	user := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username}
	updated, err := config.UserRepository.UpdateByID(uint(id), user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	userResponse := &models.UserResponse{ID: uint(id), Email: updated.Email, Username: updated.Username}
	render.JSON(w, r, userResponse)
}

// @Summary		Delete a user
// @Description	Delete a user by its ID
// @Tags		users
// @Accept		json
// @Produce		json
// @Param		id	path		int		true	"User ID"
// @Success		200	{string}	string	"Successfully deleted entry"
// @Failure 	400 {object} 	http.Error
// @Security 	BearerAuth
// @Router		/user/{id} [delete]
func (config *UserConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "User ID must be greater than 0", http.StatusBadRequest)
		return
	}
	err = config.UserRepository.DeleteByID(uint(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
