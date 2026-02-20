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

func (config *UserConfig) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	user, err := config.UserRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve user"})
		return
	}
	userResponse := &models.UserResponse{ID: user.ID, Email: user.Email, Username: user.Username}
	render.JSON(w, r, userResponse)
}

func (config *UserConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	var req models.GetUserRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request body"})
		return
	}

	var user *dbmodel.User
	var err error

	if req.Username != "" {
		user, err = config.UserRepository.FindByUsername(req.Username)
	} else if req.Email != "" {
		user, err = config.UserRepository.FindByEmail(req.Email)
	} else {
		render.JSON(w, r, map[string]string{"error": "Username or email is required"})
		return
	}

	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve user"})
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

func (config *UserConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}

	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}

	user := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username}
	updated, err := config.UserRepository.UpdateByID(uint(id), user)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	userResponse := &models.UserResponse{ID: uint(id), Email: updated.Email, Username: updated.Username}
	render.JSON(w, r, userResponse)
}

func (config *UserConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.UserRepository.DeleteByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}
	render.JSON(w, r, "Succefully deleted entry")
}
