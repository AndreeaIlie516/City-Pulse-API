package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"City-Pulse-API/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SimpleEventController struct {
	Service *services.SimpleEventService
}

func (controller *SimpleEventController) AllEvents(c *gin.Context) {
	events, err := controller.Service.AllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (controller *SimpleEventController) EventByID(c *gin.Context) {
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

func (controller *SimpleEventController) CreateEvent(c *gin.Context) {
	var newEvent entities.SimpleEvent

	if err := c.BindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	validators := map[string]validator.Func{
		"bandValidator": utils.BandValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

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

func (controller *SimpleEventController) DeleteEvent(c *gin.Context) {
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

func (controller *SimpleEventController) UpdateEvent(c *gin.Context) {
	id := c.Param("id")

	var updatedEvent entities.SimpleEvent

	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	validators := map[string]validator.Func{
		"bandValidator": utils.BandValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

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

func (controller *SimpleEventController) FavouriteEvents(c *gin.Context) {
	favouriteEvents, err := controller.Service.FavouriteEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch favourite events"})
		return
	}
	c.JSON(http.StatusOK, favouriteEvents)
}

func (controller *SimpleEventController) PrivateEvents(c *gin.Context) {
	privateEvents, err := controller.Service.PrivateEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch private events"})
		return
	}
	c.JSON(http.StatusOK, privateEvents)
}

func (controller *SimpleEventController) AddEventToFavourites(c *gin.Context) {
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

func (controller *SimpleEventController) DeleteEventFromFavourites(c *gin.Context) {
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
