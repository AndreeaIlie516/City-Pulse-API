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

type ArtistController struct {
	Service *services.ArtistService
}

func (controller *ArtistController) AllArtists(c *gin.Context) {
	artists, err := controller.Service.AllArtists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch artists"})
		return
	}
	c.JSON(http.StatusOK, artists)
}

func (controller *ArtistController) ArtistByID(c *gin.Context) {
	id := c.Param("id")
	artist, err := controller.Service.ArtistByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		}
		return
	}
	c.JSON(http.StatusOK, artist)
}

func (controller *ArtistController) CreateArtist(c *gin.Context) {
	var newArtist entities.Artist

	if err := c.BindJSON(&newArtist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	validators := map[string]validator.Func{
		"nameValidator": utils.NameValidator,
		"bandValidator": utils.BandValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

	err := validate.Struct(newArtist)

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

	artist, err := controller.Service.CreateArtist(newArtist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create artist"})
		return
	}

	c.JSON(http.StatusCreated, artist)
}

func (controller *ArtistController) DeleteArtist(c *gin.Context) {
	id := c.Param("id")

	artist, err := controller.Service.DeleteArtist(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		}
		return
	}

	c.JSON(http.StatusOK, artist)
}

func (controller *ArtistController) UpdateArtist(c *gin.Context) {
	id := c.Param("id")

	var updatedArtist entities.Artist

	if err := c.BindJSON(&updatedArtist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	validators := map[string]validator.Func{
		"nameValidator": utils.NameValidator,
		"bandValidator": utils.BandValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

	err := validate.Struct(updatedArtist)

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

	artist, err := controller.Service.UpdateArtist(id, updatedArtist)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		return
	}

	c.JSON(http.StatusOK, artist)
}
