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
		eventGroup.GET("/location/:locationId", eventController.EventsByLocationID)
		eventGroup.GET("/city/:cityId", eventController.EventsByCityID)
		eventGroup.POST("/", eventController.CreateEvent)
		eventGroup.PUT("/:id", eventController.UpdateEvent)
		eventGroup.DELETE("/:id", eventController.DeleteEvent)
	}
}
