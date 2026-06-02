package errors

import (
	"net/http"
	"yplanning/pkg/models"

	"github.com/go-chi/render"
)

func RenderError(w http.ResponseWriter, r *http.Request, status int, msg string) {
	render.Status(r, status)
	render.JSON(w, r, models.ErrorResponse{
		Error: msg,
	})
}
