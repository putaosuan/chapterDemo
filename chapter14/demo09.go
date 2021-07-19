package main

import (
	"fmt"
	"os"
)

//判断文件是否存在（重要）
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true, nil
	}
	if os.IsNotExist(err) { //如果返回的错误类型使用os.IsNotExist(err)判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果返回的错误为其他类型，则不确定是否存在
}
func main() {
	b, err := PathExists("d:/test1/b.txt")
	fmt.Println(b, err)

}
