package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initRedis(addr string, idleConn, maxConn int, idleTimeout time.Duration)  {

	pool = &redis.Pool{
		MaxIdle: idleConn,
		MaxActive: maxConn,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	return
}

// 从 Redis 的链接池中，得到一个连接

func GetConn() redis.Conn {
	return pool.Get()
}

//关闭一个从Redis的链接池中获取到的链接
func PutConn(conn redis.Conn) {
	conn.Close()
}