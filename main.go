package main

import (
	"github.com/gin-gonic/gin"

	"trash_bot/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/**/*")

	// TrashDay
	td := router.Group("/trash-day")
	{
		ctrl := controller.TrashDayController{}
		td.GET("/index", ctrl.IndexTrashDay)
		td.GET("/:id", ctrl.DetailsTrashDay)
		td.POST("/create", ctrl.CreateTrashDay)
		td.POST("/update", ctrl.UpdateTrashDay)
		td.POST("/delete", ctrl.DeleteTrashDay)
	}
	

	// Admin
	admin := router.Group("/admin")
	{
		ctrl := controller.AdminController{}
		admin.GET("/index", ctrl.IndexAdmin)
		admin.GET("/:id", ctrl.DetailsAdmin)
		admin.POST("/create", ctrl.CreateAdmin)
		admin.POST("/update", ctrl.UpdateAdmin)
		admin.POST("/delete", ctrl.DeleteAdmin)
	}

	// Comment
	comment := router.Group("/comment")
	{
		ctrl := controller.CommentController{}
		comment.GET("/index", ctrl.IndexComment)
		comment.GET("/:id", ctrl.DetailsComment)
		comment.POST("/create", ctrl.CreateComment)
		comment.POST("/update", ctrl.UpdateComment)
		comment.POST("/delete", ctrl.DeleteComment)
	}

	// サーバーの起動
	router.Run(":8080")
}