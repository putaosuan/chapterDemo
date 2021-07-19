package model

import (
	"chapterDemo/chapter18/chatroom/common/message"
	"net"
)

//因为在客户端，我们在很多地方会使用curUser，我们将其作为一个全局变量

type CurUser struct {
	Conn net.Conn
	message.User
}
