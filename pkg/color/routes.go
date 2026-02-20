package color

import (
	"yplanning/config"

	"github.com/go-chi/chi/v5"
)

/*
color routes:
POST /color/ - create a new color
GET /color/colors - get all colors (for testing purposes only)
GET /color/{id} - get a color by ID
GET /color/hexcode/{hexcode} - get a color by hex code
PUT /color/{id} - update a color by ID
DELETE /color/{id} - delete a color by ID
*/

func Routes(config *config.Config) chi.Router {
	router := chi.NewRouter()
	colorConfig := NewColorConfig(config)
	router.Post("/", colorConfig.CreateColor)
	router.Get("/colors", colorConfig.GetAllColors) // FOR TESTING PURPOSES ONLY
	router.Get("/{id}", colorConfig.GetColorByID)
	router.Get("/hexcode", colorConfig.GetByHexCode)
	router.Put("/{id}", colorConfig.UpdateColor)
	router.Delete("/{id}", colorConfig.DeleteColor)
	return router
}
