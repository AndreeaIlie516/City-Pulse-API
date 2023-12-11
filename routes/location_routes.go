package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterLocationRoutes(router *gin.Engine, locationController *controllers.LocationController) {
	locationGroup := router.Group("/locations")
	{
		locationGroup.GET("/", locationController.AllLocations)
		locationGroup.GET("/:id", locationController.LocationByID)
		locationGroup.GET("/city/:cityId", locationController.LocationsByCityID)
		locationGroup.POST("/", locationController.CreateLocation)
		locationGroup.PUT("/:id", locationController.UpdateLocation)
		locationGroup.DELETE("/:id", locationController.DeleteLocation)
	}
}
