package controllers

import (
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserPayload struct {
	Name     string `json:"title" binding:"required"`
	Nickname string `json:"content" binding:"required"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
