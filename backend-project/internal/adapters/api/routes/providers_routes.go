package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterProvidersRoutes sets up routes for providers.
func RegisterProvidersRoutes(router *gin.Engine) {
	providersGroup := router.Group("/providers")
	{
		providersGroup.GET("/", listProvidersHandler)
	}
}

func listProvidersHandler(c *gin.Context) {
	// TODO: Implement list providers logic
	c.JSON(200, gin.H{"message": "Providers listed"})
}
