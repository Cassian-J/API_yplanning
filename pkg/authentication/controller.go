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

func (config *AuthConfig) Register(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request body"})
		return
	}

	_, err := config.UserRepository.FindByEmail(req.Email)
	if err == nil {
		render.JSON(w, r, map[string]string{"error": "Email already exists"})
		return
	}
	_, err = config.UserRepository.FindByUsername(req.Username)
	if err == nil {
		render.JSON(w, r, map[string]string{"error": " email or pseudo already in use"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	userEntry := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username}
	res, err := config.UserRepository.Create(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create user"})
		return
	}
	user := &models.UserResponse{ID: res.ID, Email: res.Email, Username: res.Username}

	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}
	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

func (config *AuthConfig) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request body"})
		return
	}
	user, err := config.UserRepository.FindByEmail(req.Email)
	if err != nil {
		user, err = config.UserRepository.FindByUsername(req.Username)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

func (config *AuthConfig) Refresh(w http.ResponseWriter, r *http.Request) {
	req := &models.TokenRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}

	email, err := ParseToken(os.Getenv("REFRESH_SECRET"), req.RefreshToken)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid refresh token"})
		return
	}

	user, err := config.UserRepository.FindByEmail(email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET"), user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}
