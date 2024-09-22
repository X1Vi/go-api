package main

import (
	"RAT/models"
	"RAT/routes" // Adjust the import path as necessary

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	models.MigrateModels()
	r.Run() // listen and serve on 0.0.0.0:8080
}
