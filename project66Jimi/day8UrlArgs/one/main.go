package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {

		//获取路径参数
		name := c.Param("name")
		id := c.Param("id")
		ii,_:= strconv.Atoi(id)

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"id":ii,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {

		//获取路径参数
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year": year,
			"month":month,
		})
	})
	r.Run(":9006")
}
