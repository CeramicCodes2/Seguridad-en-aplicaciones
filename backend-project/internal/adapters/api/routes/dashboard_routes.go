package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterDashboardRoutes sets up routes for the dashboard.
func RegisterDashboardRoutes(router *gin.Engine) {
	dashboardGroup := router.Group("/dashboard")
	{
		dashboardGroup.GET("/", getDashboardHandler)
	}
}

func getDashboardHandler(c *gin.Context) {
	// TODO: Implement get dashboard logic
	c.JSON(200, gin.H{"message": "Dashboard retrieved"})
}
