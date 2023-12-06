package routes

import (
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userController.AllUsers)
		userGroup.GET("/:id", userController.UserByID)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
