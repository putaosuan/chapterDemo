package process2

import (
	"chapterDemo/chapter18/chatroom/common/message"
	"chapterDemo/chapter18/chatroom/server/model"
	"chapterDemo/chapter18/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	UserId int
	Conn   net.Conn
}

//这里我们要编写通知所有在线的用户方法
//userId 要通知其他在线用户，我上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历 onlineUsers，然后一个一个的发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知，单独写一个方法
		up.NotifyMeOnline(userId)
	}
}
func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的UserStatusNotifyMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	//将。。序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2.声明一个loginResMes，并赋值
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_ESISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_ESISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//6.发送data，我们将其封装到writePkg函数中
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

//编写一个函数serverProcessLogin,专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.声明一个loginResMes，并赋值
	var loginResMes message.LoginResMes

	//我们需要到redis数据库完成验证
	//1.使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTESISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}

	} else {
		loginResMes.Code = 200
		//这里，因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		//将登录成功的用户的userId赋值给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的Id，放入到loginResMes.UsersId
		//遍历
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}
		fmt.Println(user, "登录成功")
	}
	//if loginMes.UserId==100&&loginMes.UserPwd=="123456"{
	//	loginResMes.Code=200
	//}else {
	//	loginResMes.Code=500
	//	loginResMes.Error="该用户不存在，请注册再使用"
	//}
	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//6.发送data，我们将其封装到writePkg函数中
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
