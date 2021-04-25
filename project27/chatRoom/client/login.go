package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"go_code/project27/chatRoom/common/message"
)

/*
	登陆函数
 */
func login(userId int,pwd string)(err error) {
	//定义协议
	fmt.Println(userId,pwd)
	if userId != 100 || pwd != "123456" {
		return errors.New("账号或者密码不正确")
	}

	//连接服务端
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("客户端连接失败：",err)
		return
	}
	defer conn.Close()

	//准备发送信息
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建一个LoginMes结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = pwd

	//将数据LoginMes json格式化
	jsonData,err := json.Marshal(LoginMes)
	if err != nil{
		fmt.Println("LoginMes格式化失败",err)
		return
	}
	//将格式化的数据转为字符复制给data
	mes.Data = string(jsonData)

	//将整个mes json格式化
	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes格式化失败",err)
		return
	}

	dataLen := len(data)
	var pakLen uint32
	pakLen = uint32(dataLen)
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4],pakLen)
	n,err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	//获取键盘输入
	reader := bufio.NewReader(os.Stdin)
	line,_ := reader.ReadString('\n')
	conn.Write([]byte(line))
	return
}