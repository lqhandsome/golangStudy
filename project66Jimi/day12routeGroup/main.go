package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./404.html")
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method": "get",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method": "POST",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method": "DELETE",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		name := c.Query("name")
		name = strings.Trim(name," ")
		c.JSON(http.StatusOK,gin.H{
			"method": "PUT",
			"name":name,
		})
	})
	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound,"404.html",nil)
	})

	r.Any("/hello", func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			context.JSON(http.StatusOK,gin.H{
				"method": "Get",
			})
		case http.MethodPost:
			context.JSON(http.StatusOK,gin.H{
				"method": "post",
			})
		}
	})
	mobile := r.Group("/mobile")
	{
		mobile.GET("/1", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"method": "/mobile/1",
			})
		})
		mobile.GET("/2", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"method": "/mobile/2",
			})
		})
	}
	r.Run(":7777")
}
