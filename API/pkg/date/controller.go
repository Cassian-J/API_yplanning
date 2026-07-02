package date

import (
	"log"
	"net/http"
	"strconv"

	"yplanning/config"
	"yplanning/database/dbmodel"
	"yplanning/pkg/errors"
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
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/ [post]
func (config *DateConfig) CreateDate(w http.ResponseWriter, r *http.Request) {
	var dateRequest models.DateRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
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
		log.Println(err)
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to create date")
		return
	}

	render.JSON(w, r, models.ToDateResponse(createdDate))
}

// @Summary Get all dates
// @Description Retrieve a list of all dates
// @Tags dates
// @Accept json
// @Produce json
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/dates [get]
func (config *DateConfig) GetAllDates(w http.ResponseWriter, r *http.Request) {
	dates, err := config.DateRepository.FindAll()
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve dates")
		return
	}

	responses := make([]models.DateResponse, 0, len(dates))

	for _, d := range dates {
		responses = append(responses, *models.ToDateResponse(&d))
	}

	render.JSON(w, r, responses)
}

// @Summary Get date by ID
// @Description Retrieve a date by its ID
// @Tags dates
// @Accept json
// @Produce json
// @Param id path int true "Date ID"
// @Success 200 {object} models.DateResponse
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/{id} [get]
func (config *DateConfig) GetDateByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Error during id convertion")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}
	date, err := config.DateRepository.FindByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve date")
		return
	}

	render.JSON(w, r, models.ToDateResponse(date))
}

// @Summary Get dates by user ID
// @Description Retrieve a list of dates associated with a specific user ID
// @Tags dates
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/user/{userID} [get]
func (config *DateConfig) GetDatesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Error during user_id convertion")
		return
	}
	if userID < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "user_id must be >= 1")
		return
	}
	dates, err := config.DateRepository.FindByUserID(uint(userID))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve dates")
		return
	}

	responses := make([]models.DateResponse, 0, len(dates))

	for _, d := range dates {
		responses = append(responses, *models.ToDateResponse(&d))
	}

	render.JSON(w, r, responses)
}

// @Summary Get dates by recurrence ID
// @Description Retrieve a list of dates associated with a specific recurrence ID
// @Tags dates
// @Accept json
// @Produce json
// @Param recurrenceID path int true "Recurrence ID"
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/recurrence/{recurrenceID} [get]
func (config *DateConfig) GetDatesByRecurrenceID(w http.ResponseWriter, r *http.Request) {
	recurrenceID, err := strconv.Atoi(chi.URLParam(r, "recurrenceID"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Error during recurrence_id convertion")
		return
	}
	if recurrenceID < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "recurrence_id must be >= 1")
		return
	}
	date, err := config.DateRepository.FindByRecurrenceID(uint(recurrenceID))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve date")
		return
	}

	render.JSON(w, r, models.ToDateResponse(date))
}

// @Summary Get dates by day range
// @Description Retrieve a list of dates that fall within a specified day range
// @Tags dates
// @Accept json
// @Produce json
// @Param start query int true "Start date as Unix timestamp"
// @Param end query int true "End date as Unix timestamp"
// @Param userID query int false "User ID to filter dates (optional)"
// @Success 200 {array} models.DateResponse
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/range [post]
func (config *DateConfig) GetDateByDayRange(w http.ResponseWriter, r *http.Request) {
	var rangeRequest models.AvailabilityRequest

	if err := render.Bind(r, &rangeRequest); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}
	date := &models.AvailabilityRequest{
		DateBegin: rangeRequest.DateBegin,
		DateEnd:   rangeRequest.DateEnd,
		UserID:    rangeRequest.UserID,
	}

	dates, err := config.DateRepository.FindByDayRange(date.DateBegin, date.DateEnd, uint(date.UserID))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve dates")
		return
	}

	responses := make([]models.DateResponse, 0, len(dates))

	for _, d := range dates {
		responses = append(responses, *models.ToDateResponse(&d))
	}

	render.JSON(w, r, responses)
}

// @Summary Update a date by ID
// @Description Update the details of a date identified by its ID
// @Tags dates
// @Accept json
// @Produce json
// @Param id path int true "Date ID"
// @Param date body models.DateRequest true "Updated date details"
// @Success 200 {object} models.DateResponse
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/{id} [put]
func (config *DateConfig) UpdateDate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Error during id convertion")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}
	var dateRequest models.DateRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Invalid request payload")
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
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to update date")
		return
	}

	updated, err := config.DateRepository.FindByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to retrieve updated date")
		return
	}

	render.JSON(w, r, models.ToDateResponse(updated))
}

// @Summary Delete a date by ID
// @Description Delete a date identified by its ID
// @Tags dates
// @Accept json
// @Produce json
// @Param id path int true "Date ID"
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {object} models.ErrorResponse
// @Security 	BearerAuth
// @Router /date/{id} [delete]
func (config *DateConfig) DeleteDate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		errors.RenderError(w, r, http.StatusBadRequest, "Error during id convertion")
		return
	}
	if id < 1 {
		errors.RenderError(w, r, http.StatusBadRequest, "id must be >= 1")
		return
	}
	err = config.DateRepository.DeleteByID(uint(id))
	if err != nil {
		errors.RenderError(w, r, http.StatusInternalServerError, "Failed to delete date")
		return
	}
	render.JSON(w, r, map[string]string{"message": "Date deleted successfully"})
}
