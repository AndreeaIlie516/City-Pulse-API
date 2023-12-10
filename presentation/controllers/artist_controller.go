package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
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

	artist, err := controller.Service.UpdateArtist(id, updatedArtist)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "artist not found"})
		return
	}

	c.JSON(http.StatusOK, artist)
}
