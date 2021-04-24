package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

type Monster struct {
	Name string
	Age int
	Skill string
}
func main() {
	var monster Monster = Monster{}
	//fmt.Scanf("%v",&monster.Name)
	//fmt.Scanf("%v",&monster.Age)
	//fmt.Scanf("%v",&monster.Skill)
	fmt.Scanln(&monster.Name)
	fmt.Scanln(&monster.Age)
	fmt.Scanln(&monster.Skill)
	//获取跟字段值相关的

	//获取跟字段属性相关的
	typ := reflect.TypeOf(monster)
	val := reflect.ValueOf(monster)
	//val = val.Elem()
	count := val.NumField()
	fmt.Println(typ.Field(0).Name,val.Field(0).String())
	conn := connectRedis()
	defer conn.Close()
	for i := 0;i < count;i++{
		key := typ.Field(i).Name
		var value interface{}
		if typ.Field(i).Type.Name() == "int" {
			value = val.Field(i).Int()
			value = value.(int64)
		} else {
			value = val.Field(i).String()
			value = value.(string)
		}
		conn.Do("hset","monster",key,value)
	}

	data,err := redis.StringMap(conn.Do("hgetall","monster"))
	if err != nil {
		fmt.Println("获取数据失败，",err)
	}
	fmt.Printf("%T",data["Age"])

}
func connectRedis() redis.Conn {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis服务器失败：",err)
	}
	fmt.Println("连接redis服务器成功：",err)

	return conn
}
