package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
type user struct {
	name string
}

func main() {

	r := gin.Default()
	r.Use(m1,m2)
	r.GET("/index", indexHandle)
	r.Run(":7777")

}

func indexHandle(c *gin.Context) {
	fmt.Println("index")
	time.Sleep(time.Nanosecond * 10)
	c.JSON(200, gin.H{
		"name": "lq",
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 in")
	start := time.Now().UnixNano()
	c.Next()
	fmt.Println(time.Now().UnixNano() - start)
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in")
	c.Abort()
	c.Redirect(301,"http://www.sogo.com")
	fmt.Println("m2 out")
}
