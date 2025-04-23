package main

import (
	routes "exlora/internal/adapters/api/routes"
	hexagon "exlora/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {

	// drivens

	// drivers

	hexagon.Exlora()
	router := gin.Default()

	// Register routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterProfileRoutes(router)
	routes.RegisterEventsRoutes(router)
	routes.RegisterDashboardRoutes(router)
	routes.RegisterProvidersRoutes(router)
	routes.RegisterRepositoryRoutes(router)

	// Start server
	router.Run(":8080")
}
