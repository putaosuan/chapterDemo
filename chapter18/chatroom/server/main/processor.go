package main

import (
	"chapterDemo/chapter18/chatroom/common/message"
	"chapterDemo/chapter18/chatroom/server/process"
	"chapterDemo/chapter18/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能：根据客户端发送消息种类的不同，决定调用哪个函数来处理
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登录函数
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册函数
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)

	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
func (this *Processor) process2() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return err
			} else {
				fmt.Println("readPkg err007=", err)
				return err
			}
		}
		//fmt.Println("mes=",mes)

		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}

	}
}
