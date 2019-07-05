package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go-basic/hsp/chart/common"
	"net"
	"time"
)

// 这个结构体，完成对服务器发送和接收消息包的读取
type Transfer struct {
	Conn net.Conn
	Buf [8192]byte
}

func (tf *Transfer) ClientReadPackage() (msg common.Message, err error) {
	n, err := tf.Conn.Read(tf.Buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		fmt.Println("如果读取错误, 则休息30秒，再读数据...")
		time.Sleep(time.Second * 30)
		return
	}

	var packLen uint32
	packLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	n, err = tf.Conn.Read(tf.Buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	err = json.Unmarshal(tf.Buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return

}

func ConnNet() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	return
}