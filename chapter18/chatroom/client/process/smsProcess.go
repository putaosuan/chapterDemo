package process

import (
	"chapterDemo/chapter18/chatroom/common/message"
	"chapterDemo/chapter18/chatroom/server/utils"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建一个message
	var mes message.Message
	mes.Type = message.SmsMesType
	//2.
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUSer.UserId
	smsMes.UserStatus = curUSer.UserStatus
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=", err)
		return
	}
	tf := utils.Transfer{
		Conn: curUSer.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes write err=", err)
		return
	}
	return
}
