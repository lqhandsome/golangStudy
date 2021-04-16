package main

import (
	"fmt"
	"flag"
)
func main() {
	var user string
	var pwd string
	var host string
	var port string

	flag.StringVar(&user,"u","","用户名,默认为空")
	flag.StringVar(&pwd,"pwd","","密码,默认为空")
	flag.StringVar(&host,"host","localhost","用户名,默认为空")
	flag.StringVar(&port,"port","3306","端口号")
	flag.Parse()
	fmt.Println(user,pwd,host,port)
	fmt.Printf("%v",user)
}
