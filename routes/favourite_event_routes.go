package routes

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/infrastructure/middlewares"
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterFavouriteEventRoutes(router *gin.Engine, favouriteEventController *controllers.FavouriteEventController, roleMiddleware middlewares.IAuthMiddleware) {
	eventArtistGroup := router.Group("/favourite-events")
	{
		eventArtistGroup.GET("/", roleMiddleware.RequireRole(entities.Admin), favouriteEventController.AllFavouriteEventAssociations)
		eventArtistGroup.GET("/:id", roleMiddleware.RequireRole(entities.Admin), favouriteEventController.FavouriteEventAssociationByID)
		eventArtistGroup.GET("/associationByEventAndUser", roleMiddleware.RequireRole(entities.Admin), favouriteEventController.FavouriteEventAssociation)
		eventArtistGroup.GET("/event/:userId", roleMiddleware.RequireRole(entities.Admin), favouriteEventController.EventWithUsers)
		eventArtistGroup.GET("/user/:userId", roleMiddleware.RequireRole(entities.NormalUser), favouriteEventController.UserWithEvents)
		eventArtistGroup.POST("/", roleMiddleware.RequireRole(entities.NormalUser), favouriteEventController.AddEventToFavourites)
		eventArtistGroup.DELETE("/:id", roleMiddleware.RequireRole(entities.NormalUser), favouriteEventController.DeleteEventFromFavourites)
	}
}
