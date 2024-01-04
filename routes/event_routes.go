package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(router *gin.Engine, eventController *controllers.EventController) {
	eventGroup := router.Group("/events")
	{
		eventGroup.GET("/", eventController.AllEvents)
		eventGroup.GET("/:id", eventController.EventByID)
		eventGroup.POST("/", eventController.CreateEvent)
		eventGroup.PUT("/:id", eventController.UpdateEvent)
		eventGroup.DELETE("/:id", eventController.DeleteEvent)
		eventGroup.GET("/favourites", eventController.FavouriteEvents)
		eventGroup.PATCH("/favourites/add/:id", eventController.AddEventToFavourites)
		eventGroup.PATCH("/favourites/delete/:id", eventController.DeleteEventFromFavourites)
	}
}
