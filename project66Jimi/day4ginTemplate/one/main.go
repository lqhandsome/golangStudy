package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
)

func main() {
	r := gin.Default()
	//gin框架中给模板添加自定义标签
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//加载静态文件
	r.Static("/xxx","./statics")

	//r.LoadHTMLFiles("template/index.tmpl")

	//加载目录下的所有文件
	//一个*是文件，**是目录
	r.LoadHTMLGlob("template/**/*")

	r.GET("/posts/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title" : "posts",
		})
	})

	r.GET("/users/index",func(c *gin.Context){
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title" : "<a href='baidu.com'>点我</a>",
		})
	})
	err := r.Run(":9000")
	if err != nil {
		fmt.Println("启动失败",err)
	}
}
