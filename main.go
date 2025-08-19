package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Nickname string `json:"Nickname"`
}

var dummyUsers = []user{
	{ID: 1, Name: "wildan", Nickname: "dan"},
	{ID: 2, Name: "raniasfah", Nickname: "rani"},
	{ID: 3, Name: "fajrina", Nickname: "fajrin"},
	{ID: 3, Name: "grace", Nickname: "ges"},
}

func getUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, dummyUsers)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8000")
}
