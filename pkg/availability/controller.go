package availability

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

type AvailabilityConfig struct {
	*config.Config
}

func NewAvailibilityConfig(afg *config.Config) *AvailabilityConfig {
	return &AvailabilityConfig{Config: afg}
}

func (config *AvailabilityConfig) CreateAvailability(w http.ResponseWriter, r *http.Request) {
	var availabilityRequest models.AvailabilityRequest
	if err := render.Bind(r, &availabilityRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}
	availability := &dbmodel.Availability{
		BeginTime: availabilityRequest.DateBegin,
		EndTime:   availabilityRequest.DateEnd,
		UserID:    availabilityRequest.UserID,
	}
	createdAvailability, err := config.AvailabilityRepository.Create(availability)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create availability"})
		return
	}
	availabilityResponse := &models.AvailabilityResponse{
		ID:        createdAvailability.ID,
		DateBegin: createdAvailability.BeginTime,
		DateEnd:   createdAvailability.EndTime,
		UserID:    createdAvailability.UserID,
	}
	render.JSON(w, r, availabilityResponse)
}

func (config *AvailabilityConfig) GetAllAvailability(w http.ResponseWriter, r *http.Request) {
	availabilities, err := config.AvailabilityRepository.FindAll()
	if err != nil {
		fmt.Println("Failed to retrieve availabilities")
		return
	}
	availabilityResponse := make([]models.AvailabilityResponse, 0)
	for _, availability := range availabilities {
		availabilityResponse = append(availabilityResponse, models.AvailabilityResponse{
			ID:        availability.ID,
			DateBegin: availability.BeginTime,
			DateEnd:   availability.EndTime,
			UserID:    availability.UserID,
		})
	}
	render.JSON(w, r, availabilityResponse)
}

func (config *AvailabilityConfig) GetAvailabilityByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	availability, err := config.AvailabilityRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve availability"})
		return
	}
	availabilityResponse := &models.AvailabilityResponse{
		ID:        availability.ID,
		DateBegin: availability.BeginTime,
		DateEnd:   availability.EndTime,
		UserID:    availability.UserID,
	}
	render.JSON(w, r, availabilityResponse)
}

func (config *AvailabilityConfig) GetAvailabilitiesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		fmt.Println("Error during user_id convertion")
	}
	if userID < 1 {
		render.JSON(w, r, map[string]string{"error": "user_id must be >= 1"})
		return
	}
	availabilities, err := config.AvailabilityRepository.FindByUserID(uint(userID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve date"})
		return
	}
	availabilityResponse := make([]models.AvailabilityResponse, 0)
	for _, availability := range availabilities {
		availabilityResponse = append(availabilityResponse, models.AvailabilityResponse{
			ID:        availability.ID,
			DateBegin: availability.BeginTime,
			DateEnd:   availability.EndTime,
			UserID:    availability.UserID,
		})
	}
	render.JSON(w, r, availabilityResponse)
}

func (config *AvailabilityConfig) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	var dateRequest models.AvailabilityRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}
	availability := &dbmodel.Availability{
		BeginTime: dateRequest.DateBegin,
		EndTime:   dateRequest.DateEnd,
		UserID:    dateRequest.UserID,
	}
	err = config.AvailabilityRepository.UpdateByID(uint(id), availability)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update availability"})
		return
	}
	render.JSON(w, r, map[string]string{"message": "Availability updated successfully"})
}

func (config *AvailabilityConfig) DeleteAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.AvailabilityRepository.DeleteByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete availability"})
		return
	}
	render.JSON(w, r, map[string]string{"message": "Availability deleted successfully"})
}
