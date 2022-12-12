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

	// trash DI(Dependency Injection: オブジェクトの注入)
	var trashDayRepository repository.TrashDayRepository
	trashDayPersistance := persistance.NewTrashDayPersistance(db, trashDayRepository)
	trashDayController := controller.NewTrashDayController(trashDayPersistance)

	// admin DI
	var adminRepository repository.AdminRepository
	adminPersistance := persistance.NewAdminPersistance(db, adminRepository)
	adminController := controller.NewAdminController(adminPersistance)

	// comment DI
	var commentRepository repository.CommentRepository
	commentPersistance  := persistance.NewCommentPersistance(db, commentRepository)
	commentController := controller.NewCommentController(commentPersistance)

	router := gin.Default()
	router.LoadHTMLGlob("view/**/*")

	// TrashDay
	router.GET("/trash-day/index", trashDayController.IndexTrashDay)
	router.GET("/trash-day/:id", trashDayController.DetailTrashDay)
	router.POST("/trash-day/create", trashDayController.CreateTrashDay)
	router.POST("/trash-day/update", trashDayController.UpdateTrashDay)
	router.POST("/trash-day/delete", trashDayController.DeleteTrashDay)

	// Admin
	router.GET("/admin/index", adminController.IndexAdmin)
	router.GET("/admin/:id", adminController.DetailAdmin)
	router.POST("/admin/create", adminController.CreateAdmin)
	router.POST("/admin/update", adminController.UpdateAdmin)
	router.POST("/admin/delete", adminController.DeleteAdmin)

	// Comment
	router.GET("/comment/index", commentController.IndexComment)
	router.GET("/comment/:id", commentController.DetailsComment)
	router.POST("/comment/create", commentController.CreateComment)
	router.POST("/comment/update", commentController.UpdateComment)
	router.POST("/comment/delete", commentController.DeleteComment)

	// サーバーの起動
	router.Run(":8080")
}
