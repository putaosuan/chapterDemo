package process

import (
	"chapterDemo/chapter18/chatroom/client/model"
	"chapterDemo/chapter18/chatroom/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUSer model.CurUser //我们在登录成功后，完成对他的初始化
//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把onlineusers
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Printf("用户id:%d\n", id)
	}

}

//编写一个方法，处理返回的NotifyUserStatusMes
func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
