package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/query", func(c *gin.Context) {
		//1.取不到为其添加默认值
		name := c.DefaultQuery("name", "none")

		//2.取不到为空
		age := c.Query("age")

		//3.取不到第二个值返回false
		sex, ok := c.GetPostForm("sex")
		if !ok {
			sex = "未知"
		}

		//4.用数组的形势来取数据，把取到的数据放入数组中
		score,ok:= c.GetQueryArray("score")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
			"sex":  sex,
			"score": score,
		})
	})
	r.Run(":9003")
}
