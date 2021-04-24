package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {

	redisClient,err := redis.Dial("tcp","localhost:6377")
	if err != nil {
		fmt.Println("连接失败:",err)
	}
	defer redisClient.Close()
	_,err = redisClient.Do("hset","user1","name","lq","sex","man")
	res,err := redis.StringMap(redisClient.Do("hgetall","user1"))
	if err != nil {
		fmt.Println("获取失败",err)
	}
	fmt.Println(res)
}
