package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/home",sayHello)
	r.POST("/home",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method" : "post",
		})
	})
	r.Any("/homea",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method" : "Any",
		})
	})


	r.PUT("/home",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method" : "PUT",
		})
	})

	r.DELETE("/home",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"method" : "DELETE",
		})
	})
	r.Run("localhost:9091")
}

func sayHello(c *gin.Context) {
	c.JSON(200,gin.H{
		"message" : "ok",
		"info" : "hello go",
	})
}
