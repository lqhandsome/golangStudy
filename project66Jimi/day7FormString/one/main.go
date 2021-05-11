package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})
	r.POST("/login", func(c *gin.Context) {

		//1.第一种获取
		//name := c.PostForm("name")
		//pwd := c.PostForm("pwd")

		//2.第二种获取方式没有这个键值（不同于值为空）
		//name := c.DefaultPostForm("name","none")
		//pwd := c.DefaultPostForm("pwd","***")

		//3.取值，取不到返回false
		name, ok := c.GetPostForm("name")
		if !ok {
			name = "none"
		}
		pwd, ok := c.GetPostForm("pwd")
		if !ok {
			pwd = "null"
		}
		sex, ok := c.GetPostForm("sex")
		if !ok {
			sex = "defaule"
		}
		c.HTML(http.StatusOK,"index.html",gin.H{
				"name":name,
				"pwd":pwd,
				"sex":sex,
		})
	})
	r.Run(":9002")
}
