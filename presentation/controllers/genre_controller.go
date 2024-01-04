package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GenreController struct {
	Service *services.GenreService
}

func (controller *GenreController) AllGenres(c *gin.Context) {
	genres, err := controller.Service.AllGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch genres"})
		return
	}
	c.JSON(http.StatusOK, genres)
}

func (controller *GenreController) GenreByID(c *gin.Context) {
	id := c.Param("id")
	genre, err := controller.Service.GenreByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"})
		}
		return
	}
	c.JSON(http.StatusOK, genre)
}

func (controller *GenreController) CreateGenre(c *gin.Context) {
	var newGenre entities.Genre

	if err := c.BindJSON(&newGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newGenre)

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

	genre, err := controller.Service.CreateGenre(newGenre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create genre"})
		return
	}

	c.JSON(http.StatusCreated, genre)
}

func (controller *GenreController) DeleteGenre(c *gin.Context) {
	id := c.Param("id")

	genre, err := controller.Service.DeleteGenre(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func (controller *GenreController) UpdateGenre(c *gin.Context) {
	id := c.Param("id")

	var updatedGenre entities.Genre

	if err := c.BindJSON(&updatedGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(updatedGenre)

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

	genre, err := controller.Service.UpdateGenre(id, updatedGenre)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "genre not found"})
		return
	}

	c.JSON(http.StatusOK, genre)
}
