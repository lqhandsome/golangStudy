package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLFiles("./index.html", "./manyFiles.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/manyFiles", func(c *gin.Context) {
		c.HTML(http.StatusOK, "manyFiles.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		//上传至指定目录
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Println("保存失败", err)
		}
		c.String(http.StatusOK, "ok")
	})

	r.POST("/uploadMany", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files[]"]
		for _,file := range files{
			fmt.Println(file.Filename)
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				c.String(http.StatusOK, "bad")
			}
		}

		//上传至指定目录

		c.String(http.StatusOK, "ok")
	})

	r.Run(":9009")
}
