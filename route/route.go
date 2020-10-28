package route

import (
	"todoAPI/controller"
	"log"
	"todoAPI/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	// use the Default router from Gin and add the endpoints by indicating the 
	// HTTP Method to use with the corresponding Function
	router := gin.Default()
	authMiddleware, err := auth.SetupAuth()

	// if JWT produces an error
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	// The route “/” is welcome and only returns a simple message.
	router.GET("/", func(c * gin.Context) {
		c.String(http.StatusOK, "Welcome to the home page")
	})

	// grouping them by using the Group Function.
	v1 := router.Group("/v1")
	{
		// calling the LoginHandler from the authMiddleware
		v1.POST("/login", authMiddleware.LoginHandler)
		
		v1.POST("/register", controller.RegisterNewUser)

		// subgroups: Organize all the endpoints related to Tasks operations.
		todo := v1.Group("todo")
		{
			// calls the CreateTask function controller, 
			// only if the MiddlewareFunc from authMiddleware allows to proceed
			todo.POST("/create", authMiddleware.MiddlewareFunc(), controller.CreateTask)
		}
	}
	return router
}
