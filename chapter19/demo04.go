package main

import (
	"fmt"
	"github.com/gomodule/redigo-master/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("HMSet", "Monster", "name", "牛魔王", "age", 500, "skill", "吃饭")
	if err != nil {
		fmt.Println("HMSet err", err)
		return
	}
	r, err := redis.Strings(conn.Do("keys", "Monster*"))
	if err != nil {
		fmt.Println("HMGet err", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%v]=%v\n", i, v)
	}
}
