package routes

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/infrastructure/middlewares"
	"City-Pulse-API/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController, roleMiddleware middlewares.IAuthMiddleware) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", roleMiddleware.RequireRole(entities.Admin), userController.AllUsers)
		userGroup.GET("/:id", roleMiddleware.RequireRole(entities.NormalUser), userController.UserByID)
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.PUT("/:id", roleMiddleware.RequireRole(entities.NormalUser), userController.UpdateUser)
		userGroup.DELETE("/:id", roleMiddleware.RequireRole(entities.NormalUser), userController.DeleteUser)
	}
}
