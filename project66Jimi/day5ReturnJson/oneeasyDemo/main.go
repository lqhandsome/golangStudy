package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}
	Mgs := msg{
		Name:    "lq",
		Message: "fuck u",
		Age:     24,
	}
	arr := [3]int{1, 2, 3}
	r.GET("/json", func(r *gin.Context) {

		data := gin.H{
			"name":    "lqhandsome",
			"age":     23,
			"message": "hello json",
			"success": "true",
			"mgs":     Mgs,
			"arr":     arr,
		}
		r.JSON(http.StatusOK, data)
	})
	r.Run(":9001")
	//{"age":23,"message":"hello json","mgs":{"name":"lq","Message":"fuck u","Age":24},"name":"lqhandsome","success":"true"}
}
