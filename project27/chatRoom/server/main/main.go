package main

import (
	"fmt"
	"go_code/project27/chatRoom/server/model"
	"net"
	"github.com/garyburd/redigo/redis"
)

const (
	REDIS_IP_PORT = "localhost:6380"
)
var (
	pool *redis.Pool
)

//初始化UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool(REDIS_IP_PORT)
	initUserDao()
	listen,err := net.Listen("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("创建监听失败,",err)
		return
	}
	defer listen.Close()
	//提示信息
	fmt.Printf("[新的结构]服务器在 %v 监听...\n",listen.Addr())
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept失败",err)
		}
		//将连接信息发给协程
		go process(conn)
	}

}

func process(conn net.Conn){
	defer conn.Close()

	//这里调用总控
	process := &Processor{
		Conn: conn,
	}
	//子协程循环调用
	for  {
		err := process.mainProcess()
		if err != nil {
			fmt.Println("客户端和服务器协程退出")
			return
		}
	}
}

