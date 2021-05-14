package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/home",sayHello)
}

func sayHello(c *gin.Context) {

}