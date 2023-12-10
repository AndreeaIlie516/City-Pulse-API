package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterArtistRoutes(router *gin.Engine, artistController *controllers.ArtistController) {
	artistGroup := router.Group("/artists")
	{
		artistGroup.GET("/", artistController.AllArtists)
		artistGroup.GET("/:id", artistController.ArtistByID)
		artistGroup.POST("/", artistController.CreateArtist)
		artistGroup.PUT("/:id", artistController.UpdateArtist)
		artistGroup.DELETE("/:id", artistController.DeleteArtist)
	}
}
