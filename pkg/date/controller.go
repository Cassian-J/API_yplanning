package date

import (
	"net/http"
	"strconv"

	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type DateConfig struct {
	*config.Config
}

func NewDateConfig(cfg *config.Config) *DateConfig {
	return &DateConfig{Config: cfg}
}

// @Summary Create a new date
// @Description Create a new date with the provided details
// @Tags dates
// @Accept json
// @Produce json
// @Param date body models.DateRequest true "Date details"
// @Success 200 {object} models.DateResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/ [post]
func (config *DateConfig) CreateDate(w http.ResponseWriter, r *http.Request) {
	var dateRequest models.DateRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	date := &dbmodel.Date{
		Title:        dateRequest.Title,
		Body:         dateRequest.Body,
		BeginTime:    dateRequest.DateBegin,
		EndTime:      dateRequest.DateEnd,
		UserID:       dateRequest.UserID,
		Private:      dateRequest.Private,
		RecurrenceID: dateRequest.RecurrenceID,
		ColorID:      dateRequest.ColorID,
	}
	createdDate, err := config.DateRepository.Create(date)
	if err != nil {
		http.Error(w, "Failed to create date", http.StatusInternalServerError)
		return
	}
	dateResponse := &models.DateResponse{
		ID:           createdDate.ID,
		Title:        createdDate.Title,
		Body:         createdDate.Body,
		DateBegin:    createdDate.BeginTime,
		DateEnd:      createdDate.EndTime,
		UserID:       createdDate.UserID,
		Private:      createdDate.Private,
		RecurrenceID: createdDate.RecurrenceID,
		ColorID:      createdDate.ColorID,
	}
	render.JSON(w, r, dateResponse)
}

// @Summary Get all dates
// @Description Retrieve a list of all dates
// @Tags dates
// @Accept json
// @Produce json
// @Success 200 {array} models.DateResponse
// @Failure 500 {object} http.Error
// @Router /date/dates [get]
func (config *DateConfig) GetAllDates(w http.ResponseWriter, r *http.Request) {
	dates, err := config.DateRepository.FindAll()
	if err != nil {
		http.Error(w, "Failed to retrieve dates", http.StatusInternalServerError)
		return
	}
	DateResponse := make([]models.DateResponse, 0)
	for _, date := range dates {
		DateResponse = append(DateResponse, models.DateResponse{
			ID:           date.ID,
			Title:        date.Title,
			Body:         date.Body,
			DateBegin:    date.BeginTime,
			DateEnd:      date.EndTime,
			UserID:       date.UserID,
			Private:      date.Private,
			RecurrenceID: date.RecurrenceID,
			ColorID:      date.ColorID,
		})
	}
	render.JSON(w, r, DateResponse)
}

// @Summary Get date by ID
// @Description Retrieve a date by its ID
// @Tags dates
// @Accept json
// @Produce json
// @Param id path int true "Date ID"
// @Success 200 {object} models.DateResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/{id} [get]
func (config *DateConfig) GetDateByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	date, err := config.DateRepository.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to retrieve date", http.StatusInternalServerError)
		return
	}
	dateResponse := &models.DateResponse{
		ID:           date.ID,
		Title:        date.Title,
		Body:         date.Body,
		DateBegin:    date.BeginTime,
		DateEnd:      date.EndTime,
		UserID:       date.UserID,
		Private:      date.Private,
		RecurrenceID: date.RecurrenceID,
		ColorID:      date.ColorID,
	}
	render.JSON(w, r, dateResponse)
}

// @Summary Get dates by user ID
// @Description Retrieve a list of dates associated with a specific user ID
// @Tags dates
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/user/{userID} [get]
func (config *DateConfig) GetDatesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, "Error during user_id convertion", http.StatusBadRequest)
		return
	}
	if userID < 1 {
		http.Error(w, "user_id must be >= 1", http.StatusBadRequest)
		return
	}
	dates, err := config.DateRepository.FindByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to retrieve date", http.StatusInternalServerError)
		return
	}
	dateResponse := make([]models.DateResponse, 0)
	for _, date := range dates {
		dateResponse = append(dateResponse, models.DateResponse{
			ID:           date.ID,
			Title:        date.Title,
			Body:         date.Body,
			DateBegin:    date.BeginTime,
			DateEnd:      date.EndTime,
			UserID:       date.UserID,
			Private:      date.Private,
			RecurrenceID: date.RecurrenceID,
			ColorID:      date.ColorID,
		})
	}
	render.JSON(w, r, dateResponse)
}

// @Summary Get dates by recurrence ID
// @Description Retrieve a list of dates associated with a specific recurrence ID
// @Tags dates
// @Accept json
// @Produce json
// @Param recurrenceID path int true "Recurrence ID"
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/recurrence/{recurrenceID} [get]
func (config *DateConfig) GetDatesByRecurrenceID(w http.ResponseWriter, r *http.Request) {
	recurrenceID, err := strconv.Atoi(chi.URLParam(r, "recurrenceID"))
	if err != nil {
		http.Error(w, "Error during recurrence_id convertion", http.StatusBadRequest)
		return
	}
	if recurrenceID < 1 {
		http.Error(w, "recurrence_id must be >= 1", http.StatusBadRequest)
		return
	}
	date, err := config.DateRepository.FindByRecurrenceID(uint(recurrenceID))
	if err != nil {
		http.Error(w, "Failed to retrieve date", http.StatusInternalServerError)
		return
	}
	dateResponse := &models.DateResponse{
		ID:           date.ID,
		Title:        date.Title,
		Body:         date.Body,
		DateBegin:    date.BeginTime,
		DateEnd:      date.EndTime,
		UserID:       date.UserID,
		Private:      date.Private,
		RecurrenceID: date.RecurrenceID,
		ColorID:      date.ColorID,
	}
	render.JSON(w, r, dateResponse)
}

// @Summary Get dates by day range
// @Description Retrieve a list of dates that fall within a specified day range
// @Tags dates
// @Accept json
// @Produce json
// @Param start query string true "Start date in ISO format (e.g., 2024-01-01T00:00:00Z)"
// @Param end query string true "End date in ISO format (e.g., 2024-01-31T23:59:59Z)"
// @Param userID query int false "User ID to filter dates (optional)"
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/range [get]
func (config *DateConfig) GetDateByDayRange(w http.ResponseWriter, r *http.Request) {
	var rangeRequest models.AvailabilityRequest
	if err := render.Bind(r, &rangeRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	date := &models.AvailabilityRequest{
		DateBegin: rangeRequest.DateBegin,
		DateEnd:   rangeRequest.DateEnd,
		UserID:    rangeRequest.UserID,
	}

	dates, err := config.DateRepository.FindByDayRange(date.DateBegin, date.DateEnd, uint(date.UserID))
	if err != nil {
		http.Error(w, "Failed to retrieve dates", http.StatusInternalServerError)
		return
	}
	DateResponse := make([]models.DateResponse, 0)
	for _, date := range dates {
		DateResponse = append(DateResponse, models.DateResponse{
			ID:           date.ID,
			Title:        date.Title,
			Body:         date.Body,
			DateBegin:    date.BeginTime,
			DateEnd:      date.EndTime,
			UserID:       date.UserID,
			Private:      date.Private,
			RecurrenceID: date.RecurrenceID,
			ColorID:      date.ColorID,
		})
	}
	render.JSON(w, r, DateResponse)
}

// @Summary Update a date by ID
// @Description Update the details of a date identified by its ID
// @Tags dates
// @Accept json
// @Produce json
// @Param id path int true "Date ID"
// @Param date body models.DateRequest true "Updated date details"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/{id} [put]
func (config *DateConfig) UpdateDate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	var dateRequest models.DateRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	date := &dbmodel.Date{
		Title:        dateRequest.Title,
		Body:         dateRequest.Body,
		BeginTime:    dateRequest.DateBegin,
		EndTime:      dateRequest.DateEnd,
		UserID:       dateRequest.UserID,
		Private:      dateRequest.Private,
		RecurrenceID: dateRequest.RecurrenceID,
		ColorID:      dateRequest.ColorID,
	}
	err = config.DateRepository.UpdateByID(uint(id), date)
	if err != nil {
		http.Error(w, "Failed to update date", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Date updated successfully"})
}

// @Summary Delete a date by ID
// @Description Delete a date identified by its ID
// @Tags dates
// @Accept json
// @Produce json
// @Param id path int true "Date ID"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} http.Error
// @Failure 500 {object} http.Error
// @Router /date/{id} [delete]
func (config *DateConfig) DeleteDate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error during id convertion", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "id must be >= 1", http.StatusBadRequest)
		return
	}
	err = config.DateRepository.DeleteByID(uint(id))
	if err != nil {
		http.Error(w, "Failed to delete date", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]string{"message": "Date deleted successfully"})
}
