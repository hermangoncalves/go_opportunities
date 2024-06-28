package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router
	router := gin.Default()
	
	// Initialize routes
	initializeRoutes(router)

	// Get PORT from enviroment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Running the server
	router.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
