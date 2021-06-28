package main

import (
	"net"
	"net/rpc"
)

type HelloServer struct {

}
func main() {
	rpc.RegisterName("HelloServer","hello")
	listen,err := net.Listen("tcp",":1234")
	if err != nil {
		return
	}
	conn,err := listen.Accept()
	if err != nil {
		return
	}
	rpc.ServeConn(conn)
}

func (p *HelloServer) Hello (request string,reply *string) error {
	str := "aaa"
	reply = &str
	return nil
}
