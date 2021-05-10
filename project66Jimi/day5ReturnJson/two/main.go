package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		data := gin.H{
			"name": "lq",
			"age":  233,
		}
		context.JSON(200,data)
	})
	r.Run(":9002")
}
