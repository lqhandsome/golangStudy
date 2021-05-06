package model

import (
	"go_code/project27/chatRoom/common/message"
	"net"
)

type  CurUser struct {
	Conn net.Conn
	message.User
}
