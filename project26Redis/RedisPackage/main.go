package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
func main() {

	//options := redis.DialOption{
	//
	//}
	c, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败",err)
		return
	}
	defer c.Close()
	res ,err := c.Do("set","lq","帅气")
	c.Do("set","age",23)
	if err != nil {
		fmt.Println("操作redis错误：",err)
	}
	parseRes :=res.(string)
	data,err := redis.Strings(c.Do("keys","*")) //获取数据需要转换
	//data,err := redis.String(c.Do("keys","*")) //获取数据需要转换
	fmt.Println(parseRes)
	fmt.Println(data)

}
