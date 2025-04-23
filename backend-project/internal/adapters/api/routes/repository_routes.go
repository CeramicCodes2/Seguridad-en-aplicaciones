package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRepositoryRoutes sets up routes for repositories.
func RegisterRepositoryRoutes(router *gin.Engine) {
	repositoryGroup := router.Group("/repository")
	{
		repositoryGroup.GET("/", listRepositoriesHandler)
	}
}

func listRepositoriesHandler(c *gin.Context) {
	// TODO: Implement list repositories logic
	c.JSON(200, gin.H{"message": "Repositories listed"})
}
