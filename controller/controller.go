package controller

import (
	"net/http"
	"todoAPI/config"
	"todoAPI/model"

	jwtapple2 "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Auth is not needed for registering a new user in the system database
func RegisterNewUser(c *gin.Context) {

	var user model.User

	// checks if a JsonObject with User values is inside the request and validates
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if the user already exists and send a response with the
	// message “User already exists”
	var userCheck model.User
	config.GetDB().First(&userCheck, "username = ?", user.Username)

	if userCheck.ID > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exist"})
		return
	}

	// if not, the user is saved in the database and send the
	// message “User created successfully!”
	config.GetDB().Save(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

// This function creates a task associated with a User.
// User have to authenticated first
func CreateTask(c *gin.Context) {

	// ses the ExtractClaims() method from the jwtapple2 plugin to get the user’s id 
	// (saved in the payload function) and ensure that the user exists.
	claims := jwtapple2.ExtractClaims(c)

	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
		return
	}

	// find in the request a JSON object compatible with a Task struct 
	var todo model.Task
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// sets the user’s id reference and finally inserts the todo task.
	todo.UserId = user.ID
	config.GetDB().Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully!", "task": todo})
}
