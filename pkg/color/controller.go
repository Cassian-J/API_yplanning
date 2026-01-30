package color

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

type ColorConfig struct {
	*config.Config
}

func NewColorConfig(cfg *config.Config) *ColorConfig {
	return &ColorConfig{Config: cfg}
}

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

func (config *ColorConfig) GetColorByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	color, err := config.ColorRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve color"})
		return
	}
	colorResponse := &models.ColorResponse{ID: color.ID, HexCode: color.HexCode, Name: color.Name}
	render.JSON(w, r, colorResponse)
}

func (config *ColorConfig) GetByHexCode(w http.ResponseWriter, r *http.Request) {
	colorRequest := &models.ColorRequest{}
	if err := render.Bind(r, colorRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	color := &dbmodel.Color{
		HexCode: colorRequest.HexCode,
		Name:    colorRequest.Name,
	}
	color, err := config.ColorRepository.FindByHexCode(color.HexCode)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve color"})
		return
	}
	colorResponse := &models.ColorResponse{ID: color.ID, HexCode: color.HexCode, Name: color.Name}
	render.JSON(w, r, colorResponse)
}

func (config *ColorConfig) CreateColor(w http.ResponseWriter, r *http.Request) {
	colorRequest := &models.ColorRequest{}
	if err := render.Bind(r, colorRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	color := &dbmodel.Color{
		HexCode: colorRequest.HexCode,
		Name:    colorRequest.Name,
	}
	createdColor, err := config.ColorRepository.Create(color)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create color"})
		return
	}
	colorResponse := &models.ColorResponse{ID: createdColor.ID, HexCode: createdColor.HexCode, Name: createdColor.Name}
	render.JSON(w, r, colorResponse)
}

func (config *ColorConfig) UpdateColor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	colorRequest := &models.ColorRequest{}
	if err := render.Bind(r, colorRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	color := &dbmodel.Color{
		HexCode: colorRequest.HexCode,
		Name:    colorRequest.Name,
	}
	err = config.ColorRepository.UpdateByID(uint(id), color)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update color"})
		return
	}
	colorResponse := &models.ColorResponse{ID: uint(id), HexCode: color.HexCode, Name: color.Name}
	render.JSON(w, r, colorResponse)
}

func (config *ColorConfig) DeleteColor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.ColorRepository.DeleteByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete color"})
		return
	}
	render.JSON(w, r, map[string]string{"message": "Color deleted successfully"})
}
