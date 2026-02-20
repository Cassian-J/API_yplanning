package color

import (
	"net/http"
	"strconv"

	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ColorConfig struct {
	*config.Config
}

func NewColorConfig(cfg *config.Config) *ColorConfig {
	return &ColorConfig{Config: cfg}
}

// @Summary Create a new color
// @Description Create a new color with the provided hex code and name
// @Tags colors
// @Accept json
// @Produce json
// @Param color body models.ColorRequest true "Color information"
// @Success 200 {object} models.ColorResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /color/ [post]
func (config *ColorConfig) CreateColor(w http.ResponseWriter, r *http.Request) {
	colorRequest := &models.ColorRequest{}
	if err := render.Bind(r, colorRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	color := &dbmodel.Color{
		HexCode: colorRequest.HexCode,
		Name:    colorRequest.Name,
	}
	createdColor, err := config.ColorRepository.Create(color)
	if err != nil {
		http.Error(w, "Failed to create color", http.StatusInternalServerError)
		return
	}
	colorResponse := &models.ColorResponse{ID: createdColor.ID, HexCode: createdColor.HexCode, Name: createdColor.Name}
	render.JSON(w, r, colorResponse)
}

// @Summary Get all colors
// @Description Retrieve a list of all colors
// @Tags colors
// @Accept json
// @Produce json
// @Success 200 {array} models.ColorResponse
// @Failure 500 {object} http.Error
// @Router /color/colors [get]
func (config *ColorConfig) GetAllColors(w http.ResponseWriter, r *http.Request) {
	colors, err := config.ColorRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve colors", http.StatusInternalServerError)
		return
	}
	ColorResponse := make([]models.ColorResponse, 0)
	for _, color := range colors {
		ColorResponse = append(ColorResponse, models.ColorResponse{
			ID:      color.ID,
			HexCode: color.HexCode,
			Name:    color.Name,
		})
	}
	render.JSON(w, r, ColorResponse)
}

// @Summary Get color by ID
// @Description Retrieve a color by its ID
// @Tags colors
// @Accept json
// @Produce json
// @Param id path int true "Color ID"
// @Success 200 {object} models.ColorResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /color/{id} [get]
func (config *ColorConfig) GetColorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	color, err := config.ColorRepository.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve color", http.StatusInternalServerError)
		return
	}
	colorResponse := &models.ColorResponse{ID: color.ID, HexCode: color.HexCode, Name: color.Name}
	render.JSON(w, r, colorResponse)
}

// @Summary Get color by hex code
// @Description Retrieve a color by its hex code
// @Tags colors
// @Accept json
// @Produce json
// @Param color body models.ColorRequest true "Color information"
// @Success 200 {object} models.ColorResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /color/hex [post]
func (config *ColorConfig) GetByHexCode(w http.ResponseWriter, r *http.Request) {
	colorRequest := &models.ColorRequest{}
	if err := render.Bind(r, colorRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	color := &dbmodel.Color{
		HexCode: colorRequest.HexCode,
		Name:    colorRequest.Name,
	}
	color, err := config.ColorRepository.FindByHexCode(color.HexCode)
	if err != nil {
		http.Error(w, "Failed to retrieve color", http.StatusInternalServerError)
		return
	}
	colorResponse := &models.ColorResponse{ID: color.ID, HexCode: color.HexCode, Name: color.Name}
	render.JSON(w, r, colorResponse)
}

// @Summary Update a color by ID
// @Description Update a color's hex code and name by its ID
// @Tags colors
// @Accept json
// @Produce json
// @Param id path int true "Color ID"
// @Param color body models.ColorRequest true "Updated color information"
// @Success 200 {object} models.ColorResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /color/{id} [put]
func (config *ColorConfig) UpdateColor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion", http.StatusBadRequest)
		return
	}
	colorRequest := &models.ColorRequest{}
	if err := render.Bind(r, colorRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	color := &dbmodel.Color{
		HexCode: colorRequest.HexCode,
		Name:    colorRequest.Name,
	}
	err = config.ColorRepository.UpdateByID(uint(id), color)
	if err != nil {
		http.Error(w, "Failed to update color", http.StatusInternalServerError)
		return
	}
	colorResponse := &models.ColorResponse{ID: uint(id), HexCode: color.HexCode, Name: color.Name}
	render.JSON(w, r, colorResponse)
}

// @Summary Update a color by ID
// @Description Update a color's hex code and name by its ID
// @Tags colors
// @Accept json
// @Produce json
// @Param id path int true "Color ID"
// @Param color body models.ColorRequest true "Updated color information"
// @Success 200 {object} models.ColorResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /color/{id} [put]
func (config *ColorConfig) DeleteColor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion: "+err.Error(), http.StatusBadRequest)
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	err = config.ColorRepository.DeleteByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete color", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Color deleted successfully"})
}
