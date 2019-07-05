package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 定义一个全局的 pool
var pool *redis.Pool

// 当启动程序时， 就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle: 8, // 最大空闲链接数
		MaxActive: 0,  // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (conn redis.Conn, e error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "汤姆猫~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	// 取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

	// 如果我们需要从 pool 取出链接，一定保证链接池是没有关闭的
	// pool.Close()

	conn2 := pool.Get()

	_, err = conn2.Do("Set", "name2", "汤姆猫~~2")
	if err != nil {
		fmt.Println("conn.Do err~~~~=", err)
		return
	}

	//取出
	r2, err := redis.String(conn2.Do("Get", "name2"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("r=", r2)
}