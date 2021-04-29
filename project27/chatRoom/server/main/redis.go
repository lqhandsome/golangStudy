package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
const (
	REDIS_IP_PORT = "localhost:6380"
)
var (
	pool *redis.Pool
)
func main() {
	//关闭连接，一旦关闭连接也不能在从连接池取连接
	defer pool.Close()
	//新建一个连接
	conn := pool.Get()
	defer 	conn.Close()
	conn.Do("set","lq","lqhandsome")
	conn.Do("auth","lqhandsome")
	fmt.Println(redis.String(conn.Do("get","lq")))

}

//分配一个连接
func init() {

	pool = &redis.Pool{
		MaxIdle: 8,//最大空闲连接数
		MaxActive: 0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 0,//最大空闲时间
		Dial: func()(redis.Conn,error){
			return  redis.Dial("tcp",REDIS_IP_PORT)
		},
	}

}