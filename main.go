package main

import (
	"City-Pulse-API/database"
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"City-Pulse-API/infrastructure/dataaccess"
	"City-Pulse-API/infrastructure/middlewares"
	"City-Pulse-API/presentation/controllers"
	"City-Pulse-API/routes"
	"github.com/joho/godotenv"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectDB()

	entitiesToMigrate := []interface{}{
		&entities.Genre{},
		&entities.Artist{},
		&entities.ArtistGenre{},
		&entities.Event{},
		&entities.EventArtist{},
		&entities.City{},
		&entities.Location{},
		&entities.User{},
	}

	for _, entity := range entitiesToMigrate {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	authMiddleware := middlewares.AuthMiddleware{}

	genreRepository := dataaccess.NewGormGenreRepository(db)
	artistRepository := dataaccess.NewGormArtistRepository(db)
	artistGenreRepository := dataaccess.NewGormArtistGenreRepository(db)
	eventRepository := dataaccess.NewGormEventRepository(db)
	eventArtistRepository := dataaccess.NewGormEventArtistRepository(db)
	cityRepository := dataaccess.NewGormCityRepository(db)
	locationRepository := dataaccess.NewGormLocationRepository(db)
	userRepository := dataaccess.NewGormUserRepository(db)

	genreService := services.GenreService{Repo: genreRepository, ArtistGenreRepo: artistGenreRepository}
	artistService := services.ArtistService{Repo: artistRepository, ArtistGenreRepo: artistGenreRepository, EventArtistRepo: eventArtistRepository}
	artistGenreService := services.ArtistGenreService{Repo: artistGenreRepository, GenreRepo: genreRepository, ArtistRepo: artistRepository}
	eventService := services.EventService{Repo: eventRepository, LocationRepo: locationRepository, CityRepo: cityRepository, EventArtistRepo: eventArtistRepository}
	eventArtistService := services.EventArtistService{Repo: eventArtistRepository, EventRepo: eventRepository, ArtistRepo: artistRepository}
	cityService := services.CityService{Repo: cityRepository, LocationRepo: locationRepository}
	locationService := services.LocationService{Repo: locationRepository, CityRepo: cityRepository}
	userService := services.UserService{Repo: userRepository}

	genreController := controllers.GenreController{Service: &genreService}
	artistController := controllers.ArtistController{Service: &artistService}
	artistGenreController := controllers.ArtistGenreController{Service: &artistGenreService}
	eventController := controllers.EventController{Service: &eventService}
	eventArtistController := controllers.EventArtistController{Service: &eventArtistService}
	cityController := controllers.CityController{Service: &cityService}
	locationController := controllers.LocationController{Service: &locationService}
	userController := controllers.UserController{Service: &userService}
	routes.RegisterGenreRoutes(router, &genreController)
	routes.RegisterArtistRoutes(router, &artistController)
	routes.RegisterArtistGenreRoutes(router, &artistGenreController)
	routes.RegisterEventRoutes(router, &eventController)
	routes.RegisterEventArtistRoutes(router, &eventArtistController)
	routes.RegisterCityRoutes(router, &cityController)
	routes.RegisterLocationRoutes(router, &locationController)
	routes.RegisterUserRoutes(router, &userController, authMiddleware)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
