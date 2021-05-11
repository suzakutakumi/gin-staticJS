package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Static("/js", "src/js/") // ここを追加するだけ!
	// router.Static("HTMLで読み込むURL","読み込むファイルがあるディレクトリの場所")
	router.LoadHTMLGlob("src/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.Run()
}
