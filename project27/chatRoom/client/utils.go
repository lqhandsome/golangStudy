package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"net"
)

func readPkg(conn net.Conn)(message message.Message,err error) {
	//获取数据

	buf := make([]byte,1024 * 8)
	fmt.Println("正在读取文件...")
	//conn.Read 是一个阻塞进程 必须发送端和接收端都不能关闭
	_, err = conn.Read(buf[:4])
	if  err != nil {
		fmt.Println("读取错误:",err)
		return
	}
	fmt.Println("读取到的buf=",buf[:4])

	//根据buf[0:4]转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen 读取消息内容
	n ,err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("conn.Read fail err=",err)
		return
	}
	err = json.Unmarshal(buf[:pkgLen],&message)
	if err != nil {
		fmt.Println("Unmarshal fail err=",err)
		return
	}
	fmt.Println(string(buf[:pkgLen]))
	return
}
