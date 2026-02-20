package availability

import (
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

// @Summary Create a new availability
// @Description Create a new availability with the provided begin and end times, and user ID
// @Tags availabilities
// @Accept json
// @Produce json
// @Param availability body models.AvailabilityRequest true "Availability information"
// @Success 200 {object} models.AvailabilityResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /availability/ [post]
func (config *AvailabilityConfig) CreateAvailability(w http.ResponseWriter, r *http.Request) {
	var availabilityRequest models.AvailabilityRequest
	if err := render.Bind(r, &availabilityRequest); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	availability := &dbmodel.Availability{
		BeginTime: availabilityRequest.DateBegin,
		EndTime:   availabilityRequest.DateEnd,
		UserID:    availabilityRequest.UserID,
	}
	createdAvailability, err := config.AvailabilityRepository.Create(availability)
	if err != nil {
		http.Error(w, "Failed to create availability: "+err.Error(), http.StatusInternalServerError)
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

// @Summary Get all availabilities
// @Description Retrieve a list of all availabilities
// @Tags availabilities
// @Accept json
// @Produce json
// @Success 200 {array} models.AvailabilityResponse
// @Failure 500 {object} http.Error
// @Router /availability/availabilities [get]
func (config *AvailabilityConfig) GetAllAvailability(w http.ResponseWriter, r *http.Request) {
	availabilities, err := config.AvailabilityRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve availabilities: "+err.Error(), http.StatusInternalServerError)
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

// @Summary Get availability by ID
// @Description Retrieve an availability by its ID
// @Tags availabilities
// @Accept json
// @Produce json
// @Param id path int true "Availability ID"
// @Success 200 {object} models.AvailabilityResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /availability/{id} [get]
func (config *AvailabilityConfig) GetAvailabilityByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion: "+err.Error(), http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	availability, err := config.AvailabilityRepository.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve availability: "+err.Error(), http.StatusInternalServerError)
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

// @Summary Get availabilities by user ID
// @Description Retrieve a list of availabilities associated with a specific user ID
// @Tags availabilities
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.AvailabilityResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /availability/user/{userID} [get]
func (config *AvailabilityConfig) GetAvailabilitiesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, "Error during user_id convertion: "+err.Error(), http.StatusBadRequest)
		return
	}
	if userID < 1 {
		http.Error(w, "user_id must be >= 1", http.StatusBadRequest)
		return
	}
	availabilities, err := config.AvailabilityRepository.FindByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to retrieve availabilities: "+err.Error(), http.StatusInternalServerError)
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

// @Summary Update an availability by ID
// @Description Update an availability identified by its ID with the provided begin and end times, and user ID
// @Tags availabilities
// @Accept json
// @Produce json
// @Param id path int true "Availability ID"
// @Param availability body models.AvailabilityRequest true "Availability information"
// @Success 200 {object} map[string]string
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /availability/{id} [put]
func (config *AvailabilityConfig) UpdateAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion: "+err.Error(), http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	var dateRequest models.AvailabilityRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	availability := &dbmodel.Availability{
		BeginTime: dateRequest.DateBegin,
		EndTime:   dateRequest.DateEnd,
		UserID:    dateRequest.UserID,
	}
	err = config.AvailabilityRepository.UpdateByID(uint(id), availability)
	if err != nil {
		http.Error(w, "Failed to update availability: "+err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Availability updated successfully"})
}

// @Summary Delete an availability by ID
// @Description Delete an availability identified by its ID
// @Tags availabilities
// @Accept json
// @Produce json
// @Param id path int true "Availability ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /availability/{id} [delete]
func (config *AvailabilityConfig) DeleteAvailability(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion: "+err.Error(), http.StatusBadRequest)
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	err = config.AvailabilityRepository.DeleteByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete availability: "+err.Error(), http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Availability deleted successfully"})
}
