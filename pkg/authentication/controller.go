package authentication

import (
	"net/http"
	"os"
	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/models"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AuthConfig {
	return &AuthConfig{configuration}
}

// @Summary Register a new user
// @Description Register a new user with email, username, and password
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "User registration information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} http.Error
// @Failure 409 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /auth/register [post]
func (config *AuthConfig) Register(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := config.UserRepository.FindByEmail(req.Email)
	if err == nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}
	_, err = config.UserRepository.FindByUsername(req.Username)
	if err == nil {
		http.Error(w, "Email or username already in use", http.StatusConflict)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	userEntry := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username}
	res, err := config.UserRepository.Create(userEntry)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	user := &models.UserResponse{ID: res.ID, Email: res.Email, Username: res.Username}

	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		http.Error(w, "Failed to generate refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

// @Summary User login
// @Description Authenticate a user and return access and refresh tokens
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "User login information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} http.Error
// @Failure 401 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /auth/login [post]
func (config *AuthConfig) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	user, err := config.UserRepository.FindByEmail(req.Email)
	if err != nil {
		user, err = config.UserRepository.FindByUsername(req.Username)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		http.Error(w, "Failed to generate refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

// @Summary Refresh access token
// @Description Refresh the access token using a valid refresh token
// @Tags authentication
// @Accept json
// @Produce json
// @Param token body models.Token
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} http.Error
// @Failure 401 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /auth/refresh [post]
func (config *AuthConfig) Refresh(w http.ResponseWriter, r *http.Request) {
	req := &models.TokenRequest{}
	if err := render.Bind(r, req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	email, err := ParseToken(os.Getenv("REFRESH_SECRET"), req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	user, err := config.UserRepository.FindByEmail(email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		http.Error(w, "Failed to generate refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}
