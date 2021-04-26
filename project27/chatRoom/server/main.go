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
	fmt.Printf("服务器在 %v 监听...\n",listen.Addr())
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

	//读取客户端发送的信息
	for {
		//获取数据
		buf := make([]byte,1024 * 8)
		//conn.Read 是一个阻塞进程
		fmt.Println("正在读取文件...")
		n, err := conn.Read(buf[:4])
		if n != 4 && err != nil {
			fmt.Println("读取错误:",err)
			return
		}
		fmt.Println("读取到的buf=",buf[:4])
	}
}
