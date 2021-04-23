package main

import (
	"fmt"
	"net" //做网络socket开发时，net包含有我们需要所有的方法和函数
)
func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp","0.0.0.0:8888")

	if err != nil {
		fmt.Println("监听失败：err=",err)
		return
	}
	defer listen.Close() //延时关闭listen
	//循环等待客户端链接
	for {
		fmt.Println("等待客户端链接我")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept失败：err=",err)
		} else {
			//fmt.Printf("listen conn=%v,客户端IP=%v\n",conn,conn.RemoteAddr())
			go process(conn)
		}

	}
}

func process(conn net.Conn) {
	//这里我们循环接收
	defer conn.Close() //关闭conn

	for  {
		//创建一个新的切片
		buf := make([]byte,1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write，那么携程就阻塞在这里
		//fmt.Printf("等待客户端:%v 发送信息\n" , conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务端的read错误 err=  协程关闭",err)
			return
		}
		//3.显示客户端发送的内容
		fmt.Printf("客户端%v发送的信息:%v\n",conn.RemoteAddr().String(),string(buf[:n]))
	}
}
