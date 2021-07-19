package testcase2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {

	date, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	fileName := "d:/test1/xuliehua.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	_, err = write.Write(date)
	write.Flush()
	if err != nil {
		fmt.Println("写入失败", err)
		return false
	}
	return true
}
func (m *Monster) ReStore() bool {

	fileName := "d:/test1/xuliehua.txt"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	var date []byte
	for {
		date, err = reader.ReadBytes('\n')
		fmt.Println(date)
		if err == io.EOF {
			break
		}
	}
	err = json.Unmarshal(date, m)
	fmt.Println(m)
	if err != nil {
		return false
	}
	return true
}
