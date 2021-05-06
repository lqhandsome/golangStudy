package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"net"
)

//将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8086]byte
}
func (this *Transfer)ReadPkg()(message message.Message,err error) {
	//获取数据

	//conn.Read 是一个阻塞进程 必须发送端和接收端都不能关闭
	_, err = this.Conn.Read(this.Buf[:4])
	if  err != nil {
		fmt.Println("读取错误:",err)
		return
	}
	fmt.Println("读取到的buf=",this.Buf[:4])

	//根据buf[0:4]转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen 读取消息内容
	n ,err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("conn.Read fail err=",err)
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen],&message)
	if err != nil {
		fmt.Println("Unmarshal fail err=",err)
		return
	}
	fmt.Println(string(this.Buf[:pkgLen]))
	return
}

//服务器发送给客户端信息
func(this *Transfer)WritePkg(data []byte)(err error) {
	//先发送一个长度给对方
	var pakLen uint32
	pakLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pakLen)
	//发送长度
	n,err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	//发送data
	n, err = this.Conn.Write(data)
	if err != nil || n != int(pakLen){
		fmt.Println("写入data失败",err)
		return
	}
	fmt.Println("发送给服务端的数据：",string(data))
	return
}