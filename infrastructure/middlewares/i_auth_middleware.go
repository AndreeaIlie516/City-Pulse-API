package middlewares

import (
	"City-Pulse-API/domain/entities"
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	RequireRole(role entities.AccessType) gin.HandlerFunc
}
