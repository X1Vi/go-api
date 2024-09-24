package routes

import (
	"RAT/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Ping)
	r.POST("/createUser/", controllers.CreateUser)
	r.POST("/createTodo/", controllers.CreateToDoHandler)
}
