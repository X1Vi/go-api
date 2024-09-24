package controllers

import (
	"RAT/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var new_user models.User

	// Bind JSON input to the User struct
	if err := c.ShouldBindJSON(&new_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var db *gorm.DB = models.GetDb()

	db.Create(&new_user)
	// Perform any necessary logic like database insertion here
	// For now, we're just returning the input as a response
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    new_user,
	})
}
