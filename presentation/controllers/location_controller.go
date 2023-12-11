package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LocationController struct {
	Service *services.LocationService
}

func (controller *LocationController) AllLocations(c *gin.Context) {
	locations, err := controller.Service.AllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch locations"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func (controller *LocationController) LocationByID(c *gin.Context) {
	id := c.Param("id")
	location, err := controller.Service.LocationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}
	c.JSON(http.StatusOK, location)
}

func (controller *LocationController) LocationsByCityID(c *gin.Context) {
	cityID := c.Param("cityId")
	locations, err := controller.Service.LocationsByCityID(cityID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "city id not found"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func (controller *LocationController) CreateLocation(c *gin.Context) {
	var newLocation entities.Location

	if err := c.BindJSON(&newLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location, err := controller.Service.CreateLocation(newLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create location"})
		return
	}

	c.JSON(http.StatusCreated, location)
}

func (controller *LocationController) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	location, err := controller.Service.DeleteLocation(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

func (controller *LocationController) UpdateLocation(c *gin.Context) {
	id := c.Param("id")

	var updatedLocation entities.Location

	if err := c.BindJSON(&updatedLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location, err := controller.Service.UpdateLocation(id, updatedLocation)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}
