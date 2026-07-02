package authentication

import (
	"net/http"
	"os"
	"strconv"

	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/errors"
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
// @Param user body models.CreateUserRequest true "User registration information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
func (config *AuthConfig) Register(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	_, err := config.UserRepository.FindByEmail(req.Email)
	if err == nil {
		errors.RenderError(w, r, http.StatusConflict, "Email already exists")
		return
	}
	_, err = config.UserRepository.FindByUsername(req.Username)
	if err == nil {
		errors.RenderError(w, r, http.StatusConflict, "Email or username already in use")
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	userEntry := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username, ColorID: req.ColorID}
	res, err := config.UserRepository.Create(userEntry)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}
	user := models.ToUserResponse(res)

	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET_KEY"), strconv.Itoa(int(user.ID)))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET_KEY"), strconv.Itoa(int(user.ID)))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to generate refresh token: "+err.Error())
		return
	}
	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
		ID:           user.ID,
	}

	render.JSON(w, r, tokens)
}

// @Summary User login
// @Description Authenticate a user and return access and refresh tokens
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "User login information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/login [post]
func (config *AuthConfig) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}
	user, err := config.UserRepository.FindByEmail(req.Email)
	if err != nil {
		user, err = config.UserRepository.FindByUsername(req.Username)
		if err != nil {
			errors.RenderError(w, r, http.StatusUnauthorized, "Invalid email or password")
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		errors.RenderError(w, r, http.StatusUnauthorized, "Invalid email or password")
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET_KEY"), strconv.Itoa(int(user.ID)))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET_KEY"), strconv.Itoa(int(user.ID)))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to generate refresh token: "+err.Error())
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
		ID:           user.ID,
	}

	render.JSON(w, r, tokens)
}

// @Summary Refresh access token
// @Description Refresh the access token using a valid refresh token
// @Tags authentication
// @Accept json
// @Produce json
// @Param token body models.TokenRequest true "Token for authentication"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/refresh [post]
func (config *AuthConfig) Refresh(w http.ResponseWriter, r *http.Request) {
	req := &models.TokenRequest{}
	if err := render.Bind(r, req); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	userID, err := ParseToken(os.Getenv("REFRESH_SECRET_KEY"), req.RefreshToken)
	if err != nil {
		errors.RenderError(w, r, http.StatusUnauthorized, "Invalid refresh token")
		return
	}
	ID, err := strconv.Atoi(userID)
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to parse user ID: "+err.Error())
		return
	}
	user, err := config.UserRepository.FindByID(uint(ID))
	if err != nil {
		errors.RenderError(w, r, http.StatusNotFound, "User not found")
		return
	}
	accessToken, err := GenerateToken(os.Getenv("JWT_SECRET_KEY"), strconv.Itoa(int(user.ID)))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to generate token: "+err.Error())
		return
	}
	refreshToken, err := GenerateRefreshToken(os.Getenv("REFRESH_SECRET_KEY"), strconv.Itoa(int(user.ID)))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to generate refresh token: "+err.Error())
		return
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
		ID:           user.ID,
	}

	render.JSON(w, r, tokens)
}
