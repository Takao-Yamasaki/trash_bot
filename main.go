package main

import (
	"github.com/gin-gonic/gin"
	sqlite "trash_bot/config/database"


	"trash_bot/controller"
	"trash_bot/domain/repository"
	"trash_bot/infrastructure/persistance"
	"trash_bot/usecase"
)

func main() {
	db := sqlite.New()
	connect, _ := db.DB()
	defer connect.Close()

	// trash DI(Dependency Injection: オブジェクトの注入)
	trashDayRepository := persistance.NewTrashDayPersistance(db)
	trashDayUseCase := usecase.NewTrashDayUseCase(trashDayRepository)
	trashDayController := controller.NewTrashDayController(trashDayUseCase)

	// admin DI
	var adminRepository repository.AdminRepository
	adminPersistance := persistance.NewAdminPersistance(db, adminRepository)
	adminUseCase := usecase.NewAdminUseCase(adminPersistance)
	adminController := controller.NewAdminController(adminUseCase)

	// comment DI
	var commentRepository repository.CommentRepository
	commentPersistance  := persistance.NewCommentPersistance(db, commentRepository)
	commentUseCase := usecase.NewCommentUseCase(commentPersistance)
	commentController := controller.NewCommentController(commentUseCase)

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
	router.GET("/comment/:id", commentController.DetailComment)
	router.POST("/comment/create", commentController.CreateComment)
	router.POST("/comment/update", commentController.UpdateComment)
	router.POST("/comment/delete", commentController.DeleteComment)

	// サーバーの起動
	router.Run(":8080")
}
