package date

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

type DateConfig struct {
	*config.Config
}

func NewDateConfig(cfg *config.Config) *DateConfig {
	return &DateConfig{Config: cfg}
}

func (config *DateConfig) GetAllDates(w http.ResponseWriter, r *http.Request) {
	dates, err := config.DateRepository.FindAll()
	if err != nil {
		fmt.Println("Failed to retrieve dates")
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

func (config *DateConfig) GetDateByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	date, err := config.DateRepository.FindByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve date"})
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

func (config *DateConfig) GetDatesByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		fmt.Println("Error during user_id convertion")
	}
	if userID < 1 {
		render.JSON(w, r, map[string]string{"error": "user_id must be >= 1"})
		return
	}
	dates, err := config.DateRepository.FindByUserID(uint(userID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve date"})
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

func (config *DateConfig) GetDatesByRecurrenceID(w http.ResponseWriter, r *http.Request) {
	recurrenceID, err := strconv.Atoi(chi.URLParam(r, "recurrence_id"))
	if err != nil {
		fmt.Println("Error during recurrence_id convertion")
	}
	if recurrenceID < 1 {
		render.JSON(w, r, map[string]string{"error": "recurrence_id must be >= 1"})
		return
	}
	date, err := config.DateRepository.FindByRecurrenceID(uint(recurrenceID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve date"})
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

func (config *DateConfig) GetDateByDayRange(w http.ResponseWriter, r *http.Request) {
	var rangeRequest models.AvailabilityRequest
	if err := render.Bind(r, &rangeRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}
	date := &models.AvailabilityRequest{
		DateBegin: rangeRequest.DateBegin,
		DateEnd:   rangeRequest.DateEnd,
		UserID:    rangeRequest.UserID,
	}

	dates, err := config.DateRepository.FindByDayRange(date.DateBegin, date.DateEnd, uint(date.UserID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve dates"})
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

func (config *DateConfig) CreateDate(w http.ResponseWriter, r *http.Request) {
	var dateRequest models.DateRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
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
		render.JSON(w, r, map[string]string{"error": "Failed to create date"})
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

func (config *DateConfig) UpdateDate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	var dateRequest models.DateRequest
	if err := render.Bind(r, &dateRequest); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
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
		render.JSON(w, r, map[string]string{"error": "Failed to update date"})
		return
	}
	render.JSON(w, r, map[string]string{"message": "Date updated successfully"})
}

func (config *DateConfig) DeleteDate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("Error during id convertion")
	}
	if id < 1 {
		render.JSON(w, r, map[string]string{"error": "id must be >= 1"})
		return
	}
	err = config.DateRepository.DeleteByID(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete date"})
		return
	}
	render.JSON(w, r, map[string]string{"message": "Date deleted successfully"})
}
