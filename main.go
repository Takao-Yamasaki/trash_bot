package main

import (
	sqlite "trash_bot/config/database"

	"github.com/gin-gonic/gin"

	"trash_bot/controller"
	"trash_bot/domain/repository"
	"trash_bot/infrastructure/persistance"
)

func main() {
	db := sqlite.New()
	connect, _ := db.DB()
	defer connect.Close()

	// DI(Dependency Injection: オブジェクトの注入)
	var trashDayRepository repository.TrashDayRepository
	trashDayPersistance := persistance.NewTrashDayPersistance(db, trashDayRepository)
	trashDayController := controller.NewTrashDayController(trashDayPersistance)

	router := gin.Default()
	router.LoadHTMLGlob("view/**/*")

	// TrashDay
	router.GET("/trash-day/index", trashDayController.IndexTrashDay)
	router.GET("/trash-day/:id", trashDayController.DetailTrashDay)
	router.POST("/trash-day/create", trashDayController.CreateTrashDay)
	router.POST("/trash-day/update", trashDayController.UpdateTrashDay)
	router.POST("/trash-day/delete", trashDayController.DeleteTrashDay)

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
