package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		}
		return
	}
	c.JSON(http.StatusOK, event)
}

func (controller *EventController) EventsByLocationID(c *gin.Context) {
	locationID := c.Param("locationId")
	events, err := controller.Service.EventsByLocationID(locationID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, events)
}

func (controller *EventController) EventsByCityID(c *gin.Context) {
	cityID := c.Param("cityId")
	events, err := controller.Service.EventsByCityID(cityID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, events)
}

func (controller *EventController) CreateEvent(c *gin.Context) {
	var newEvent entities.Event

	if err := c.BindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newEvent)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
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
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		}
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

	validate := validator.New()

	err := validate.Struct(updatedEvent)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	event, err := controller.Service.UpdateEvent(id, updatedEvent)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (controller *EventController) FavouriteEvents(c *gin.Context) {
	favouriteEvents, err := controller.Service.FavouriteEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch favourite events"})
		return
	}
	c.JSON(http.StatusOK, favouriteEvents)
}

func (controller *EventController) AddEventToFavourites(c *gin.Context) {
	id := c.Param("id")
	event, err := controller.Service.AddEventToFavourites(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		}
		return
	}
	c.JSON(http.StatusOK, event)
}

func (controller *EventController) DeleteEventFromFavourites(c *gin.Context) {
	id := c.Param("id")
	event, err := controller.Service.DeleteEventFromFavourites(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		}
		return
	}
	c.JSON(http.StatusOK, event)
}
