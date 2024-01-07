package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterSimpleEventRoutes(router *gin.Engine, eventController *controllers.SimpleEventController) {
	eventGroup := router.Group("/simple-events")
	{
		eventGroup.GET("/", eventController.AllEvents)
		eventGroup.GET("/:id", eventController.EventByID)
		eventGroup.POST("/", eventController.CreateEvent)
		eventGroup.PUT("/:id", eventController.UpdateEvent)
		eventGroup.DELETE("/:id", eventController.DeleteEvent)
		eventGroup.GET("/favourites", eventController.FavouriteEvents)
		eventGroup.GET("/privates", eventController.PrivateEvents)
		eventGroup.PATCH("/favourites/add/:id", eventController.AddEventToFavourites)
		eventGroup.PATCH("/favourites/delete/:id", eventController.DeleteEventFromFavourites)
	}
}
