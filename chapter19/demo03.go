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
	_, err = conn.Do("HMSet", "user02", "name", "smith张磊", "age", 25)
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err=", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%v]=%v\n", i, v)
	}
}
