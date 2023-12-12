package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventArtistController struct {
	Service *services.EventArtistService
}

func (controller *EventArtistController) AllEventArtistAssociations(c *gin.Context) {
	eventArtistAssociations, err := controller.Service.AllEventArtistAssociations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch event artist associations"})
		return
	}
	c.JSON(http.StatusOK, eventArtistAssociations)
}

func (controller *EventArtistController) EventArtistAssociationByID(c *gin.Context) {
	id := c.Param("id")
	eventArtistAssociation, err := controller.Service.EventArtistAssociationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event artist association not found"})
		return
	}
	c.JSON(http.StatusOK, eventArtistAssociation)
}

func (controller *EventArtistController) EventArtistAssociation(c *gin.Context) {
	eventID := c.Query("eventId")
	artistID := c.Query("artistId")

	if eventID == "" || artistID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "eventId and artistId query parameters are required"})
		return
	}

	eventArtistAssociation, err := controller.Service.EventArtistAssociation(eventID, artistID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event artist association not found"})
		return
	}
	c.JSON(http.StatusOK, eventArtistAssociation)
}

func (controller *EventArtistController) EventWithArtists(c *gin.Context) {
	eventID := c.Param("eventId")
	eventWithArtists, err := controller.Service.EventWithArtists(eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, eventWithArtists)
}

func (controller *EventArtistController) ArtistWithEvents(c *gin.Context) {
	artistID := c.Param("artistId")
	artistWithEvents, err := controller.Service.ArtistWithEvents(artistID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist id not found"})
		return
	}
	c.JSON(http.StatusOK, artistWithEvents)
}

func (controller *EventArtistController) CreateEventArtistAssociation(c *gin.Context) {
	var newEventArtistAssociation entities.EventArtist

	if err := c.BindJSON(&newEventArtistAssociation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventArtistAssociation, err := controller.Service.CreateEventArtistAssociation(newEventArtistAssociation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, eventArtistAssociation)
}

func (controller *EventArtistController) DeleteEventArtistAssociation(c *gin.Context) {
	id := c.Param("id")

	eventArtistAssociation, err := controller.Service.DeleteEventArtistAssociation(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event artist association not found"})
		return
	}

	c.JSON(http.StatusOK, eventArtistAssociation)
}

func (controller *EventArtistController) UpdateEventArtistAssociation(c *gin.Context) {
	id := c.Param("id")

	var updatedEventArtistAssociation entities.EventArtist

	if err := c.BindJSON(&updatedEventArtistAssociation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventArtistAssociation, err := controller.Service.UpdateEventArtistAssociation(id, updatedEventArtistAssociation)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist event association not found"})
		return
	}

	c.JSON(http.StatusOK, eventArtistAssociation)
}
