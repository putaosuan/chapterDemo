package main

import (
	"fmt"
	"os"
)

func main() {
	//打开一个文件
	file, err := os.Open("d:/test1/a.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	fmt.Printf("file=%v\n", file)
	err = file.Close()
	if err != nil {
		fmt.Println("open file err = ", err)
	}
}
