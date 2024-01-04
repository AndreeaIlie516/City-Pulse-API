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
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
		}
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

	validate := validator.New()

	validators := map[string]validator.Func{
		"nameValidator":    utils.NameValidator,
		"countryValidator": utils.CountryValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

	err := validate.Struct(newCity)

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
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
		}
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

	validate := validator.New()

	validators := map[string]validator.Func{
		"nameValidator":    utils.NameValidator,
		"countryValidator": utils.CountryValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

	err := validate.Struct(updatedCity)

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

	city, err := controller.Service.UpdateCity(id, updatedCity)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, city)
}
