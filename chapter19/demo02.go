package main

import (
	"fmt"
	"github.com/gomodule/redigo-master/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	_, err = conn.Do("HS	et", "user01", "age", 20)
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	fmt.Println(r1, r2)

}
