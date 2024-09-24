package controllers

import (
	"RAT/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateToDo creates a new to-do item and associates it with a user
func CreateToDoHandler(c *gin.Context) {
	// Bind JSON data from the request body
	var requestBody struct {
		Username string `json:"username"`
		Task     string `json:"task"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if requestBody.Task == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task is empty"})
		return
	}
	// Call the function to create the ToDo
	todo, err := CreateToDo(requestBody.Username, requestBody.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created ToDo as a JSON response
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateToDo(username string, task string) (*models.ToDo, error) {
	var user models.User
	db := models.GetDb()
	// Find the user by their username
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error // Return an error if the user is not found
	}

	// Create a new ToDo and associate it with the user
	newToDo := models.ToDo{
		ToDo:   task,
		IsDone: false, // Default value
		UserID: user.ID,
	}

	// Save the ToDo in the database
	if err := db.Create(&newToDo).Error; err != nil {
		return nil, err // Return an error if the creation fails
	}

	return &newToDo, nil // Return the created ToDo
}
