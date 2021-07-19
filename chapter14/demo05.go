package main

import (
	"bufio"
	"fmt"
	"os"
)

//打开一个存在的文件，将原来的内容覆盖成新的内容
func main() {
	fileName := "d:/test1/b.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "你好，尚硅谷！\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
