package main

import (
	"fmt"
	"github.com/gomodule/redigo-master/redis"
)

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数
		IdleTimeout: 100, //表示最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "name", "汤姆猫")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println(r)
	//如果我们从pool取出连接，一定要保证连接池没有关闭
	pool.Close()
	conn2 := pool.Get()
	fmt.Println(conn2)

}
