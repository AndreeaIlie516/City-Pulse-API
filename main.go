package main

import (
	"City-Pulse-API/domain/services"
	"City-Pulse-API/infrastructure/dataaccess"
	"City-Pulse-API/presentation/controllers"
	"City-Pulse-API/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	genreRepository := dataaccess.NewInMemoryGenreRepository()
	genreService := services.GenreService{Repo: genreRepository}
	genreController := controllers.GenreController{Service: &genreService}

	cityRepository := dataaccess.NewInMemoryCityRepository()
	cityService := services.CityService{Repo: cityRepository}
	cityController := controllers.CityController{Service: &cityService}

	routes.RegisterGenreRoutes(router, &genreController)
	routes.RegisterCityRoutes(router, &cityController)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
