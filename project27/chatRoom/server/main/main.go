package main

import (
	"fmt"
	"net"
)

func main() {
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
	err := process.mainProcess()
	if err != nil {
		 fmt.Println("客户端和服务器协程退出")
		return
	}
}

