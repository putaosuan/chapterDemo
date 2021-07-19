package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//将原先的内容输出到屏幕上再追加写入三句话
func main() {
	fileName := "d:/test1/b.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str1, err := reader.ReadString('\n')
		print(str1)
		if err == io.EOF {
			fmt.Println()
			break
		}

	}

	str2 := "ABC！\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str2)
	}
	writer.Flush()

}
