package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `form:"username" json:"user_name"`
	PassWord string	`form:"password" json:"pass_word"`
}
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//
		//u := UserInfo{
		//	username,
		//	password,
		//}
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "false",
				"data": u,
			})
		}
		fmt.Printf("%#v\n",u)
		c.JSON(http.StatusOK, gin.H{
			"message": "true",
			"data": u,
		})

	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)

		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK, gin.H{
				"message": "true",
				"data": u,
			})
		}
	})
	r.Run(":9008")
}
