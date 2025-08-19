package controllers

import (
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserPayload struct {
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.User{Name: input.Name, Nickname: input.Nickname}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})

}
