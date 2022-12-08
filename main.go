package main

import (
	"github.com/gin-gonic/gin"

	"trash_bot/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")
	
	router.GET("/trash-days", controller.Index)
	router.GET("/trash-day/:id", controller.DetailsTrashDay)	
	router.POST("trash-day/create", controller.CreateTrashDay)
	router.POST("/trash-day/update", controller.UpdateTrashDay)
	router.POST("/trash-day/delete", controller.DeleteTrashDay)
	
	// サーバーの起動
	router.Run(":8080")
}