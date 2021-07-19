package main

import (
	"fmt"
	"io/ioutil"
)

//编写一个程序 ，将一个文件的内容写入到另外一个文件中
func main() {
	//将b.txt文件内容，导入到c.txt中去
	fileName1 := "d:/test1/b.txt"
	fileName2 := "d:/test1/c.txt"
	date, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(fileName2, date, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
