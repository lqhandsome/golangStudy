package client

import (
	"fmt"
	"net"
)

func main() {
	conn,err := net.Dial("tcp","127.0.0.1:777")
	if err != nil {
		fmt.Println("拨号失败",err)
	}
	str := "hello world"
	conn.Write([]byte(str))
}
