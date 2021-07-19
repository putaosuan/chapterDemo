package process

import (
	"chapterDemo/chapter18/chatroom/common/message"
	"chapterDemo/chapter18/chatroom/server/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcessor struct {
}

//登录
func (this *UserProcessor) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)
	//return nil
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("ner.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("son.Marshal err=", err)
		return
	}
	//5.把mes.Data赋给
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("son.Marshal err=", err)
		return
	}
	//7.这个时候，data就是我们要发送的数据
	//先获取到一个长度->转换成一个表示长度的byte切片
	var pakLen uint32
	pakLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pakLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	fmt.Printf("客户端，发送消息的长度%d，内容=%s\n", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	//休眠
	//time.Sleep(10*time.Second)
	//fmt.Println("休眠了一会")
	//这里还需要处理服务器端返回的信息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}
	//将mes的Data部分发序列化成LgoinResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		curUSer.Conn = conn
		curUSer.UserId = userId
		curUSer.UserStatus = message.UserOnline
		//fmt.Println("登录成功")
		//显示我们登录成功的界面
		//可以显示当前在线的用户列表，遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UserId {
			//如果我们要求不显示自己在线，下面我们增加一行代码
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端onlineUsers 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return

}

//注册
func (this *UserProcessor) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("ner.Dial err=", err)
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("son.Marshal err=", err)
		return
	}
	//5.把mes.Data赋给
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("son.Marshal err=", err)
		return
	}
	//7.这个时候，data就是我们要发送的数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册信息发送错误err=", err)
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}
	//将mes的Data部分发序列化成registerResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)

	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}
