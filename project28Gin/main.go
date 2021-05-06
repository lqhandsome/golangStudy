package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/aa", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action

		a := map[string]string{"name":"lq","data":message}
		json , _  := json.Marshal(a)
		c.String(http.StatusOK, string(json))
	})

	router.Run(":8081")
}