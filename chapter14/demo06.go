package main

import (
	"bufio"
	"fmt"
	"os"
)

//追加
func main() {
	fileName := "d:/test1/b.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "zhuijia！\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
