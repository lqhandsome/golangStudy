package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"net"
	//"../common/message"
)

/*
	登陆函数
 */
func login(userId int,pwd string)(err error) {
	//定义协议
	fmt.Println(userId,pwd)

	//连接服务端
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("客户端连接失败：",err)
		return
	}
	//延时关闭连接
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
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pakLen)
	n,err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	fmt.Println("客户端发送的长度",data,buf[0:4])
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("写入data失败",err)
		return
	}
	msg,err := readPkg(conn)
	if err != nil {
		fmt.Println("读取失败76",err)
	}

	//将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(msg.Data),&loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println("登录失败",loginResMes.Error)
	}
	return
}