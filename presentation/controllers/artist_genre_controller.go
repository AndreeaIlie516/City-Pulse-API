package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArtistGenreController struct {
	Service *services.ArtistGenreService
}

func (controller *ArtistGenreController) AllArtistGenreAssociations(c *gin.Context) {
	artistGenreAssociations, err := controller.Service.AllArtistGenreAssociations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch artist genre associations"})
		return
	}
	c.JSON(http.StatusOK, artistGenreAssociations)
}

func (controller *ArtistGenreController) ArtistGenreAssociationByID(c *gin.Context) {
	id := c.Param("id")
	artistGenreAssociation, err := controller.Service.ArtistGenreAssociationByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "artistGenreAssociation not found"})
		}
		return
	}
	c.JSON(http.StatusOK, artistGenreAssociation)
}

func (controller *ArtistGenreController) ArtistGenreAssociation(c *gin.Context) {
	artistID := c.Query("artistId")
	genreID := c.Query("genreId")

	if artistID == "" || genreID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "artistId and genreId query parameters are required"})
		return
	}

	artistGenreAssociation, err := controller.Service.ArtistGenreAssociation(artistID, genreID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "artistGenreAssociation not found"})
		}
		return
	}
	c.JSON(http.StatusOK, artistGenreAssociation)
}

func (controller *ArtistGenreController) ArtistWithGenre(c *gin.Context) {
	artistID := c.Param("artistId")
	artistWithGenres, err := controller.Service.ArtistWithGenres(artistID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist id not found"})
		}
		return
	}
	c.JSON(http.StatusOK, artistWithGenres)
}

func (controller *ArtistGenreController) GenreWithArtist(c *gin.Context) {
	genreID := c.Param("genreId")
	genreWithArtists, err := controller.Service.GenreWithArtists(genreID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "genre id not found"})
		}
		return
	}
	c.JSON(http.StatusOK, genreWithArtists)
}

func (controller *ArtistGenreController) CreateArtistGenreAssociation(c *gin.Context) {
	var newArtistGenreAssociation entities.ArtistGenre

	if err := c.BindJSON(&newArtistGenreAssociation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newArtistGenreAssociation)

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

	artistGenreAssociation, err := controller.Service.CreateArtistGenreAssociation(newArtistGenreAssociation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, artistGenreAssociation)
}

func (controller *ArtistGenreController) DeleteArtistGenreAssociation(c *gin.Context) {
	id := c.Param("id")

	artistGenreAssociation, err := controller.Service.DeleteArtistGenreAssociation(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "artist genre association not found"})
		}
		return
	}

	c.JSON(http.StatusOK, artistGenreAssociation)
}

func (controller *ArtistGenreController) UpdateArtistGenreAssociation(c *gin.Context) {
	id := c.Param("id")

	var updatedArtistGenreAssociation entities.ArtistGenre

	if err := c.BindJSON(&updatedArtistGenreAssociation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(updatedArtistGenreAssociation)

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

	artistGenreAssociation, err := controller.Service.UpdateArtistGenreAssociation(id, updatedArtistGenreAssociation)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist genre association not found"})
		return
	}

	c.JSON(http.StatusOK, artistGenreAssociation)
}
