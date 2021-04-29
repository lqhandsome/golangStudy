package main

import "github.com/garyburd/redigo/redis"

//分配一个连接
func initPool(address string) {

	pool = &redis.Pool{
		MaxIdle: 8,//最大空闲连接数
		MaxActive: 0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 0,//最大空闲时间
		Dial: func()(redis.Conn,error){
			return  redis.Dial("tcp",address)
		},
	}

}