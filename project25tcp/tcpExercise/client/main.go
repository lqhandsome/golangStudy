package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main()  {
	conn, err := net.Dial("tcp","localhost:8888")
	if err != nil {
		fmt.Println("进程创建失败")
	}
	defer conn.Close()
	//循环写入
	for true {
		fmt.Print("请输入要发送的数据：")
		//获取用户的输入数据
		reader := bufio.NewReader(os.Stdin)
		line,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("获取用户输入失败",err)
		}
		//删除多余的字符
		line = strings.Trim(line,"\n\r")
		if line == "exit" {
			fmt.Printf("您已经退出...退出时间: %v",time.Now().Format("2006-01-02 15:04:05"))
			return
		}
		conn.Write([]byte(line))
	}
}
