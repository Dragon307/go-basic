package main

import (
	"github.com/EDDYCJY/redis-protocol-example/protocol"
	"log"
	"net"
	"os"
)

const (
	Address = "127.0.0.1:6379"
	Network = "tcp"
)

func Conn(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main()  {
	// 读取入参
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatalln("Os.Args <= 0")
	}

	// 获取请求协议
	reqCommand := protocol.GetRequest(args)

	// 连接redis 服务器
	redisConn, err := Conn(Network, Address)
	if err != nil {
		log.Fatalf("Conn err: %v", err)
	}
	defer redisConn.Close()

	// 写入请求内容
	_, err = redisConn.Write(reqCommand)
	if err != nil {
		log.Fatalf("Conn Write err: %v", err)
	}

	// 读取回复
	command := make([]byte, 1024)
	n, err := redisConn.Read(command)
	if err != nil {
		log.Fatalf("Conn Read err: %v", err)
	}

	// 处理回复
	replay, err := protocol.GetReply(command[:n])
	if err != nil {
		log.Fatalf("protocol.GetReplay err: %v", err)
	}

	// 处理后的回复内容
	log.Printf("Reply: %v", replay)
	// 原始的回复内容
	log.Printf("Command: %v", string(command[:n]))
}

//1、读取命令行参数：获取执行的 Redis 命令
//
//2、获取请求协议参数
//
//3、连接 Redis 服务器，获取连接句柄
//
//4、将请求协议参数写入连接：发送请求的命令行参数
//
//5、从连接中读取返回的数据：读取先前请求的回复数据
//
//6、根据回复“协议”内容，处理回复的数据集
//
//7、输出处理后的回复内容及原始回复内容


// go run main.go SET test01 value01