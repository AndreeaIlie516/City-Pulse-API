package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterEventArtistRoutes(router *gin.Engine, eventArtistController *controllers.EventArtistController) {
	eventArtistGroup := router.Group("/event-artist")
	{
		eventArtistGroup.GET("/", eventArtistController.AllEventArtistAssociations)
		eventArtistGroup.GET("/:id", eventArtistController.EventArtistAssociationByID)
		eventArtistGroup.GET("/associationByEventAndArtist", eventArtistController.EventArtistAssociation)
		eventArtistGroup.GET("/event/:eventId", eventArtistController.EventWithArtists)
		eventArtistGroup.GET("/artist/:artistId", eventArtistController.ArtistWithEvents)
		eventArtistGroup.POST("/", eventArtistController.CreateEventArtistAssociation)
		eventArtistGroup.DELETE("/:id", eventArtistController.DeleteEventArtistAssociation)
		eventArtistGroup.PUT("/:id", eventArtistController.UpdateEventArtistAssociation)
	}
}
