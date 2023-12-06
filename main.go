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

	routes.RegisterGenreRoutes(router, &genreController)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
