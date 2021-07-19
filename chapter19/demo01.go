package main

import (
	"fmt"
	"github.com/gomodule/redigo-master/redis"
)

func main() {
	//通过go 向redis写入数据和读取数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	//fmt.Println("conn suc",conn)
	_, err = conn.Do("set", "name", "jack张磊")
	if err != nil {
		fmt.Println("set=", err)
		return
	}
	r, err := redis.String(conn.Do("get", "name007"))
	if err != nil {
		fmt.Println("get=", err)
		return
	}
	fmt.Println("ok", r)

}
