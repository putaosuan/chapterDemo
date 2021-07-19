package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("d:/test1/a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //要及时关闭句柄，否则会有内存泄漏
	//创建一个*reader 带缓冲
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF { //io.EOF表示文件的末尾
			fmt.Println()
			break
		}

	}
	fmt.Println("文件读取结束")

}
