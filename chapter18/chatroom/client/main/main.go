package main

import (
	"chapterDemo/chapter18/chatroom/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户ID，一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {
	//接受用户的选择
	var key int
	//判断是否继续显示菜单
	//var loop = true

	for true {
		fmt.Println("------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1.登录聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcessor{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的ID")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的昵称(nickname)")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcessor{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入:")
		}
	}
	//if key ==1{
	//	fmt.Println("请输入用户的ID")
	//	fmt.Scanf("%d\n",&userId)
	//	fmt.Println("请输入用户的密码")
	//	fmt.Scanf("%s\n",&userPwd)
	//	//把登录的函数写到另一个文件中
	//	login(userId,userPwd)
	//	//if err!=nil{
	//	//	fmt.Println("登陆失败")
	//	//}else {
	//	//	fmt.Println("登录成功")
	//	//}
	//}else if key==2{
	//	fmt.Println("jinxinzhuce")
	//}
}
