package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `form:"username"`
	PassWord string	`form:"password"`
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
	r.Run(":9008")
}
