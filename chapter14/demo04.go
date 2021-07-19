package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "d:/test1/b.txt"
	//1.打开文件
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	//及时关闭句柄，防止内存泄露
	defer file.Close()
	str := "hello Gardon\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用WriterString方法时，需要调用flush方法，将缓存的数据真正写入到文件中，
	//否则文件中
	writer.Flush()

}
