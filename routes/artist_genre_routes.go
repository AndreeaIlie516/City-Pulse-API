package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterArtistGenreRoutes(router *gin.Engine, artistGenreController *controllers.ArtistGenreController) {
	artistGenreGroup := router.Group("/artist-genre")
	{
		artistGenreGroup.GET("/", artistGenreController.AllArtistGenreAssociations)
		artistGenreGroup.GET("/:id", artistGenreController.ArtistGenreAssociationByID)
		artistGenreGroup.GET("/associationByArtistAndGenre", artistGenreController.ArtistGenreAssociation)
		artistGenreGroup.GET("/artist/:artistId", artistGenreController.ArtistWithGenre)
		artistGenreGroup.GET("/genre/:genreId", artistGenreController.GenreWithArtist)
		artistGenreGroup.POST("/", artistGenreController.CreateArtistGenreAssociation)
		artistGenreGroup.DELETE("/:id", artistGenreController.DeleteArtistGenreAssociation)
		artistGenreGroup.PUT("/:id", artistGenreController.UpdateArtistGenreAssociation)
	}
}
