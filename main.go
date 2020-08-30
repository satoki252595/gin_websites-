package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// エンジンインスタンスの生成
	router := gin.Default()

	// テンプレートをロード（相対パス）
	router.LoadHTMLGlob("templates/*.html")

	router.Static("css", "./css")

	// 変数の宣言
	data := "Hello Variables!"

	// ハンドラの宣言 第1引数にアクセスされたら第2引数の関数が実行される。
	router.GET("/", func(ctx *gin.Context) {
		// HTMLのレンダリング 1.ステータスコード 2.読むテンプレート名 3.渡す変数
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.POST("/name", func(ctx *gin.Context) {
		// HTMLのレンダリング 1.ステータスコード 2.読むテンプレート名 3.渡す変数
		firstName := ctx.PostForm("firstName")
		ctx.HTML(200, "submit1.html", gin.H{"firstName": firstName})
	})

	// エンジンの実行
	router.Run()
}
