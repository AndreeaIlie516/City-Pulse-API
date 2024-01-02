package main

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"City-Pulse-API/infrastructure/dataaccess"
	"City-Pulse-API/presentation/controllers"
	"City-Pulse-API/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	"log"
)

func main() {
	router := gin.Default()

	dsn := "host=localhost user=postgres password=postgres dbname=CityPulse port=5433 sslmode=disable TimeZone=Europe/Bucharest"
    

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %s", err.Error())
    }

	err = db.AutoMigrate(&entities.User{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

	genreRepository := dataaccess.NewInMemoryGenreRepository()
	artistRepository := dataaccess.NewInMemoryArtistRepository()
	artistGenreRepository := dataaccess.NewInMemoryArtistGenreRepository()
	eventRepository := dataaccess.NewInMemoryEventRepository()
	eventArtistRepository := dataaccess.NewInMemoryEventArtistRepository()
	cityRepository := dataaccess.NewInMemoryCityRepository()
	locationRepository := dataaccess.NewInMemoryLocationRepository()
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
	routes.RegisterUserRoutes(router, &userController)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
