package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port string

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.StringVar(&port, "port", "3306", "端口号，默认为3306")
	//必须调用该方法
	flag.Parse()
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n", user, pwd, host, port)
}
