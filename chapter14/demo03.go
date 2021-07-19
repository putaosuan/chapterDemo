package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//一次性将文件内容读取到程序中
	fileName := "d:/test1/a.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
	}
	fmt.Println(content)
	fmt.Printf("%s", content)

}
