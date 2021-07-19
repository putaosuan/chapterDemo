package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//将一个文件拷贝到另外一个目录下或者文件不存在，就创建文件
func CopyFile(dstFileName, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}
func main() {
	src := "D:/test/tupian/tupian.jpg"
	dst := "d:/test1/12.jpg"
	num, err := CopyFile(dst, src)
	if err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
	fmt.Println(num)
}
