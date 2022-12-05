package main

import (
	"github.com/gin-gonic/gin"

	"trash_bot/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")

	router.GET("/", controller.Index)

	router.Run(":8080")
}