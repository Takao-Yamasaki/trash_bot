package main

import (
	"github.com/gin-gonic/gin"

	"trash_bot/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/trashday/*html")
	router.LoadHTMLGlob("view/admin/*html")
	router.LoadHTMLGlob("view/comment/*html")

	// trashday
	router.GET("/trash-days", controller.IndexTrashDay)
	router.GET("/trash-day/:id", controller.DetailsTrashDay)
	router.POST("trash-day/create", controller.CreateTrashDay)
	router.POST("/trash-day/update", controller.UpdateTrashDay)
	router.POST("/trash-day/delete", controller.DeleteTrashDay)

	// admin
	router.GET("/admins", controller.IndexAdmin)
	router.GET("/admin/:id", controller.DetailsAdmin)
	router.POST("admin/create", controller.CreateAdmin)
	router.POST("/admin/update", controller.UpdateAdmin)
	router.POST("/admin/delete", controller.DeleteAdmin)

	// commets
	router.GET("/comments", controller.IndexComment)
	router.GET("/comment/create", controller.CreateComment)

	// サーバーの起動
	router.Run(":8080")
}
