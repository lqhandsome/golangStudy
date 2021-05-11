package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//重定向
	r.GET("/index", func(context *gin.Context) {
		//context.JSON(http.StatusOK,gin.H{
		//	"status": "ok",
		//})
		context.Redirect(301, "http://www.sogo.com")
	})

	//隐式修改
	r.GET("/a", func(context *gin.Context) {
		//context.JSON(http.StatusOK,gin.H{
		//	"status": "ok",
		//})
		context.Request.URL.Path = "/b"
		r.HandleContext(context)
	})
	r.GET("/b", func(context *gin.Context) {
		//context.JSON(http.StatusOK,gin.H{
		//	"status": "ok",
		//})
		context.JSON(200, gin.H{
			"message": "b",
		})
	})
	r.Run(":7777")
}
