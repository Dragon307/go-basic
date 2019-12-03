package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on "+ *host + ":" + *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		//data := make([]byte, 1024)
		//conn.Read(data)
		//fmt.Println(string(data))

		// logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn)  {
	defer conn.Close()

	for {
		io.Copy(conn, conn)
	}
}