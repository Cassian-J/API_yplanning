package date

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
date routes:
POST /dates - Create a new date
GET /dates - Get all dates (for testing purposes only)
GET /dates/{id} - Get a date by ID
GET /dates/user/{userID} - Get dates by user ID
GET /dates/recurrence/{recurrenceID} - Get dates by recurrence ID
GET /dates/range?start={startDate}&end={endDate} - Get dates within a specific day range
PUT /dates/{id} - Update a date by ID
DELETE /dates/{id} - Delete a date by ID
*/

func Routes(config *config.Config) chi.Router {
	router := chi.NewRouter()
	dateConfig := NewDateConfig(config)
	router.Post("/", dateConfig.CreateDate)
	router.Get("/dates", dateConfig.GetAllDates) //FOR TESTING PURPOSES ONLY
	router.Get("/{id}", dateConfig.GetDateByID)
	router.Get("/user/{userID}", dateConfig.GetDatesByUserID)
	router.Get("/recurrence/{recurrenceID}", dateConfig.GetDatesByRecurrenceID)
	router.Get("/range", dateConfig.GetDateByDayRange)
	router.Put("/{id}", dateConfig.UpdateDate)
	router.Delete("/{id}", dateConfig.DeleteDate)
	return router
}
