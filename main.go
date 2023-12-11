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
	artistRepository := dataaccess.NewInMemoryArtistRepository()
	artistGenreRepository := dataaccess.NewInMemoryArtistGenreRepository()
	cityRepository := dataaccess.NewInMemoryCityRepository()
	locationRepository := dataaccess.NewInMemoryLocationRepository()
	userRepository := dataaccess.NewInMemoryUserRepository()

	genreService := services.GenreService{Repo: genreRepository, ArtistGenreRepo: artistGenreRepository}
	artistService := services.ArtistService{Repo: artistRepository, ArtistGenreRepo: artistGenreRepository}
	artistGenreService := services.ArtistGenreService{Repo: artistGenreRepository, GenreRepo: genreRepository, ArtistRepo: artistRepository}
	cityService := services.CityService{Repo: cityRepository, LocationRepo: locationRepository}
	locationService := services.LocationService{Repo: locationRepository, CityRepo: cityRepository}
	userService := services.UserService{Repo: userRepository}

	genreController := controllers.GenreController{Service: &genreService}
	artistController := controllers.ArtistController{Service: &artistService}
	artistGenreController := controllers.ArtistGenreController{Service: &artistGenreService}
	cityController := controllers.CityController{Service: &cityService}
	locationController := controllers.LocationController{Service: &locationService}
	userController := controllers.UserController{Service: &userService}

	routes.RegisterGenreRoutes(router, &genreController)
	routes.RegisterArtistRoutes(router, &artistController)
	routes.RegisterArtistGenreRoutes(router, &artistGenreController)
	routes.RegisterCityRoutes(router, &cityController)
	routes.RegisterLocationRoutes(router, &locationController)
	routes.RegisterUserRoutes(router, &userController)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
