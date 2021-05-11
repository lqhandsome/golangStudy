package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//1.普通的，没有值则为空
		name := c.Query("name")

		//2.没有值则默认值
		age := c.DefaultQuery("age", "88")

		//3.接收array
		like := c.QueryArray("like")

		//4.取不到值返回false
		sex, ok := c.GetQuery("sex")
		if !ok {
			sex = "未知"
		}
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
			"like":  like,
			"sex":  sex,
		})
	})
	r.Run(":9000")
}
