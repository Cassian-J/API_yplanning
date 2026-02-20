package availability

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
availability routes:
POST /availability/
GET /availability/availabilities
GET /availability/{id}
GET /availability/user/{userID}
PUT /availability/{id}
DELETE /availability/{id}
*/

func Routes(config *config.Config) chi.Router {
	AvailabilityConfig := NewAvailibilityConfig(config)
	router := chi.NewRouter()
	router.Post("/", AvailabilityConfig.CreateAvailability)
	router.Get("/availabilities", AvailabilityConfig.GetAllAvailability) // FOR TESTING PURPOSES ONLY
	router.Get("/{id}", AvailabilityConfig.GetAvailabilityByID)
	router.Get("/user/{userID}", AvailabilityConfig.GetAvailabilitiesByUserID)
	router.Put("/{id}", AvailabilityConfig.UpdateAvailability)
	router.Delete("/{id}", AvailabilityConfig.DeleteAvailability)
	return router
}