package main

import (
	"log"
	"net/http"
	"os"
	"yplanning/config"
	"yplanning/pkg/authentication"
	"yplanning/pkg/availability"
	"yplanning/pkg/color"
	"yplanning/pkg/date"
	"yplanning/pkg/group"
	"yplanning/pkg/user"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			LocateThis API
// @version			1.0
// @host			localhost:8080
// @BasePath		/api
// @securityDefinitions.apikey	BearerAuth
// @in				header
// @name			Authorization
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	router.Mount("/api/auth", authentication.Routes(configuration))

	router.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware(os.Getenv("JWT_SECRET")))
		r.Mount("/api/groups", group.Routes(configuration))
		r.Mount("/api/date", date.Routes(configuration))
		r.Mount("/api/availability", availability.Routes(configuration))
		r.Mount("/api/color", color.Routes(configuration))
		r.Mount("/api/users", user.Routes(configuration))
	})

	return router
}

func main() {
	// Initialisation de la configuration
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}
	godotenv.Load()
	// Initialisation des routes
	router := Routes(configuration)

	log.Println("Server running on http://localhost:" + os.Getenv("PORT"))
	log.Println("Swagger UI available at http://localhost:" + os.Getenv("PORT") + "/swagger/index.html")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
