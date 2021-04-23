package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func main() {

	conn, err := net.Dial("tcp","localhost:8888")
	if err != nil {
		fmt.Println("client dial err=",err)
		return
	}
	defer conn.Close()

	//功能一，客户端可以发送单行数据，然后就退出

	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入
	//从终端读取一行用户输入，并准备发送给服务器
	for  {
		fmt.Print("请输入：")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("reading err=",err)
		}
		line = strings.Trim(line,"\n\r")
		if line == "exit" {
			fmt.Print("客户端已经退出。。。")
			return
		}
		//再将line发送给服务器
		n,errr := conn.Write([]byte(line))
		fmt.Printf("客户端发送了 %d 字节的数据，数据： %v \n",n,line)
		if errr != nil {
			fmt.Println("conn write err=",errr)
			return
		}
	}


}
