package controllers

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	Service *services.UserService
}

func (controller *UserController) AllUsers(c *gin.Context) {
	users, err := controller.Service.AllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (controller *UserController) UserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := controller.Service.UserByID(id)
	if err != nil {
        if err.Error() == "invalid ID format" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
        } else {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        }
        return
    }
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var newUser entities.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newUser)

	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
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

	user, err := controller.Service.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	user, err := controller.Service.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser entities.User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(updatedUser)

	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
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

	user, err := controller.Service.UpdateUser(id, updatedUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
