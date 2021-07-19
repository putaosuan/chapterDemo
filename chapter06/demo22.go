package main

import (
	"errors"
	"fmt"
)

func test22() {
	defer func() {
		err := recover() //recover()内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err=", err)
			fmt.Println("错误发送给管理员")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
func test02() {
	err := readConf("confi.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02()继续执行")

}
func main() {
	//test22()
	//for{
	//	fmt.Println("main下面的代码...")
	//	time.Sleep(time.Second)
	//}
	test02()
}
