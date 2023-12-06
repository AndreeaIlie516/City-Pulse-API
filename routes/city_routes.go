package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCityRoutes(router *gin.Engine, cityController *controllers.CityController) {
	cityGroup := router.Group("/cities")
	{
		cityGroup.GET("/", cityController.AllCities)
		cityGroup.GET("/:id", cityController.CityByID)
		cityGroup.POST("/", cityController.CreateCity)
		cityGroup.PUT("/:id", cityController.UpdateCity)
		cityGroup.DELETE("/:id", cityController.DeleteCity)
	}
}
