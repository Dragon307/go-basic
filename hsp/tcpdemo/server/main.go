package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	// 循环接收客户端发送的数据
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		// 1 等待客户端通过 conn 发送消息
		// 2 如果客户端没有 write ， 那么协程就阻塞再这里
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return
		}
		//  显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听....")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1. tcp 表示使用网络协议是tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//循环等待客户端来链接我
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)

		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备其一个协程，为客户端服务
		go process(conn)
	}

	//fmt.Printf("listen suc=%v\n", listen)
}