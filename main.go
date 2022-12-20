package main

import (
	sqlite "trash_bot/config/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"trash_bot/controller"
	"trash_bot/infrastructure/persistance"
	"trash_bot/middleware"
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
	adminRepository := persistance.NewAdminPersistance(db)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminController := controller.NewAdminController(adminUseCase)

	// comment DI
	commentRepository := persistance.NewCommentPersistance(db)
	commentUseCase := usecase.NewCommentUseCase(commentRepository)
	commentController := controller.NewCommentController(commentUseCase)

	router := gin.Default()
	router.LoadHTMLGlob("view/**/*")

	store := cookie.NewStore([]byte("kokonisecretkeyiretene"))
	router.Use(sessions.Sessions("session", store))

	// TrashDay
	router.GET("/trash-day/index", trashDayController.IndexTrashDay)
	router.GET("/trash-day/:id", trashDayController.DetailTrashDay)
	router.POST("/trash-day/create", trashDayController.CreateTrashDay)
	router.POST("/trash-day/update", trashDayController.UpdateTrashDay)
	router.POST("/trash-day/delete", trashDayController.DeleteTrashDay)

	// Admin
	router.GET("/admin/index", adminController.IndexAdmin)
	router.POST("/admin/create", adminController.CreateAdmin)
	router.POST("/admin/update", adminController.UpdateAdmin)
	router.POST("/admin/delete", adminController.DeleteAdmin)

	// Comment
	router.GET("/comment/index", commentController.IndexComment)
	router.GET("/comment/:id", commentController.DetailComment)
	router.POST("/comment/create", commentController.CreateComment)
	router.POST("/comment/update", commentController.UpdateComment)
	router.POST("/comment/delete", commentController.DeleteComment)

	// login
	router.GET("/login", adminController.Login)
	router.POST("/login", adminController.AuthLogin)
	router.GET("/logout", adminController.Logout)

	authAdminGroup := router.Group("/")
	authAdminGroup.Use(middleware.IsLogin())
	{
		authAdminGroup.GET("/admin/:id", adminController.DetailAdmin)
	}

	// サーバーの起動
	router.Run(":8080")
}
