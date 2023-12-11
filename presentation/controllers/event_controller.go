package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventController struct {
	Service *services.EventService
}

func (controller *EventController) AllEvents(c *gin.Context) {
	events, err := controller.Service.AllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (controller *EventController) EventByID(c *gin.Context) {
	id := c.Param("id")
	event, err := controller.Service.EventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (controller *EventController) CreateEvent(c *gin.Context) {
	var newEvent entities.Event

	if err := c.BindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := controller.Service.CreateEvent(newEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create event"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (controller *EventController) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := controller.Service.DeleteEvent(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (controller *EventController) UpdateEvent(c *gin.Context) {
	id := c.Param("id")

	var updatedEvent entities.Event

	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := controller.Service.UpdateEvent(id, updatedEvent)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}
