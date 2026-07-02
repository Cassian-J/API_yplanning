package main

import (
	"log"
	"net/http"
	"yplanning/config"
	"yplanning/pkg/authentication"
	"yplanning/pkg/availability"
	"yplanning/pkg/color"
	"yplanning/pkg/date"
	"yplanning/pkg/group"
	"yplanning/pkg/groupuser"
	"yplanning/pkg/user"

	_ "yplanning/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Yplanning API
// @version			1.0
// @BasePath		/api
// @securityDefinitions.apikey	BearerAuth
// @in				header
// @name			Authorization
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	router.Mount("/api/auth", authentication.Routes(configuration))

	router.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware(configuration.JWTSecret))
		r.Mount("/api/group", group.Routes(configuration))
		r.Mount("/api/date", date.Routes(configuration))
		r.Mount("/api/availability", availability.Routes(configuration))
		r.Mount("/api/color", color.Routes(configuration))
		r.Mount("/api/user", user.Routes(configuration))
		r.Mount("/api/group-user", groupuser.Routes(configuration))
	})

	return router
}

func main() {
	// Initialisation de la configuration
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}
	// Initialisation des routes
	router := Routes(configuration)
	log.Println("Server running on http://" + configuration.ADDRESS_IP + ":" + configuration.Port)
	log.Println("Swagger UI available at http://" + configuration.ADDRESS_IP + ":" + configuration.Port + "/swagger/index.html")
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}