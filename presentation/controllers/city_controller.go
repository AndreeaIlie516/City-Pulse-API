package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CityController struct {
	Service *services.CityService
}

func (controller *CityController) AllCities(c *gin.Context) {
	cities, err := controller.Service.AllCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cities"})
		return
	}
	c.JSON(http.StatusOK, cities)
}

func (controller *CityController) CityByID(c *gin.Context) {
	id := c.Param("id")
	city, err := controller.Service.CityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
		return
	}
	c.JSON(http.StatusOK, city)
}

func (controller *CityController) CreateCity(c *gin.Context) {
	var newCity entities.City

	if err := c.BindJSON(&newCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city, err := controller.Service.CreateCity(newCity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create city"})
		return
	}

	c.JSON(http.StatusCreated, city)
}

func (controller *CityController) DeleteCity(c *gin.Context) {
	id := c.Param("id")

	city, err := controller.Service.DeleteCity(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
		return
	}

	c.JSON(http.StatusOK, city)
}

func (controller *CityController) UpdateCity(c *gin.Context) {
	id := c.Param("id")

	var updatedCity entities.City

	if err := c.BindJSON(&updatedCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city, err := controller.Service.UpdateCity(id, updatedCity)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
		return
	}

	c.JSON(http.StatusOK, city)
}
