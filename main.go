package main

import (
	"go-rest/controllers"
	"go-rest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
	router.PATCH("/user/:id", controllers.UpdateUser)
	router.Run("localhost:8000")
}
