package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func writeDataToFile(syncChan chan bool) {
	fileName := "d:/test1/2021.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	rand.Seed(time.Now().UnixNano())
	str := ""
	for i := 0; i < 10; i++ {
		str = str + strconv.Itoa(rand.Intn(100)) + "英雄" + "\n"
	}
	n, err := write.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	write.Flush()
	syncChan <- true
	close(syncChan)
}
func sort(syncChan chan bool) {
	<-syncChan
	fileName := "d:/test1/2021.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}
func main() {
	syncChan := make(chan bool, 1)
	go writeDataToFile(syncChan)
	go sort(syncChan)
}
