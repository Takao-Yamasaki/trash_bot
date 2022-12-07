package main

import (
	"github.com/gin-gonic/gin"

	"trash_bot/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")
	// 一覧の取得
	router.GET("/", controller.Index)
	// データ1件取得
	router.GET("/trash-day/:id", controller.DetailsTrashDay)	
	// 登録
	router.POST("trash-day/create", controller.CreateTrashDay)
	// 更新
	router.PUT("/trash-day/update", controller.UpdateTrashDay)
	// 削除
	router.DELETE("/trash-day/delete", controller.DeleteTrashDay)
	// サーバーの起動
	router.Run(":8080")
}