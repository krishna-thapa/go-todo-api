package main

import (
	"log"
	"os"
	"todoAPI/config"
	"todoAPI/migration"
	"todoAPI/route"

	"github.com/gin-gonic/gin"
)

// gets the Database instance and creates the
// Migration on it (if not exist)
func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// setup function for deploy the routes
	router := route.SetupRoutes()

	// finally creates a server with the indicated port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
