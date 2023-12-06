package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterGenreRoutes(router *gin.Engine, genreController *controllers.GenreController) {
	genreGroup := router.Group("/genres")
	{
		genreGroup.GET("/", genreController.AllGenres)
		genreGroup.GET("/:id", genreController.GenreByID)
		genreGroup.POST("/", genreController.CreateGenre)
		genreGroup.PUT("/:id", genreController.UpdateGenre)
		genreGroup.DELETE("/:id", genreController.DeleteGenre)
	}
}
