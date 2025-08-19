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
type UpdateUserPayload struct {
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserPayload

	// assign body payload request to 'input' variable
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userPayload := models.User{Name: input.Name, Nickname: input.Nickname}
	models.DB.Create(&userPayload)

	c.JSON(http.StatusOK, gin.H{"data": userPayload})

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

func UpdateUser(c *gin.Context) {
	var user models.User
	var input UpdateUserPayload

	// find data with id
	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found!"})
		return
	}

	// assign body payload request to 'input' variable
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser := models.User{Name: input.Name, Nickname: input.Nickname}

	// update user
	models.DB.Model(&user).Updates(&updatedUser)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
