package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	//创建服务器监听
	lister, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("服务器监听失败！",err)
		return
	}
	//延迟关闭
	defer lister.Close()

	for  {
		//创建等待链接事件
		fmt.Println("服务器监听完成，正在等待客户端链接")
		conn,err := lister.Accept()
		if err != nil {
			fmt.Println("失败！",err)
		}
		go process(conn)

	}
}

func process(conn net.Conn) {
	//协程开启
	defer conn.Close()

	for  {
		//声明切片保存用户输入
		buf := make([]byte,1024)

		n,err := conn.Read(buf)

		if err != nil{
			fmt.Println("一个客户端失去了链接")
			return
		}
		//现实客户端发送的信息
		message := string(buf[:n])
		message = strings.Trim(message,"\n\r")
		fmt.Printf("%v客户端%v:%v\n",time.Now().Format("2006:1:02 15:04:05"),conn.RemoteAddr().String(),message)
	}

}


