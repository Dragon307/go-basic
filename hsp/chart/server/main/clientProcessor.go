package main

import (
	"errors"
	"fmt"
	"go-basic/hsp/chart/common"
	Process "go-basic/hsp/chart/server/process"
	"go-basic/hsp/chart/server/utils"
	"net"
)

type ClientProcessor struct {
	conn net.Conn
	buf [8192]byte
}

func (cp *ClientProcessor) ServerProcessMsg(msg *common.Message) (err error) {
	fmt.Println(msg)
	// 根据消息的类型来进行不同的处理
	switch msg.Type {
	case common.LoginMesType: // 如果是用户登录信息，就交给userProcess 处理
		up := &Process.UserProcessor{
			Conn: cp.conn,
		}
		err = up.ServerProcessLogin(msg)
	case common.RegisterMesType:
		up := &Process.UserProcessor{
			Conn: cp.conn,
		}
		err = up.ServerProcessRegister(msg)
	case common.TalkMesType:
		up := &Process.UserProcessor{
			Conn: cp.conn,
		}
		err = up.ServerProcessTalk(msg)

	default:
		err = errors.New("unsupport message")
	}
	return
}

func (cp *ClientProcessor) Process() (err error) {
	for {
		var msg common.Message
		// 因为要返回消息包给客户端，因此创建一个Transfer实例
		tf := &utils.Transfer{
			Conn: cp.conn,
		}
		// 将消息从客户端读取，并封装到msg
		msg, err = tf.ServerReadPackage()
		fmt.Println(msg, "Process")
		if err != nil {
			fmt.Println("server read package err=", err)
			return err
		}

		//将读取并封装好的msg 交给cp的serverProcessMsg
		err = cp.ServerProcessMsg(&msg)
		if err != nil {
			fmt.Println("process msg failed, err=", err)
			//continue
			return err
		} else {
			//如果我们处理没有任何错误，作为测试，就先退出这个goroutine
			//即退出这个for循环，否则会继续执行 readPackage
			//但是，客户端没有发消息，这就服务器会报 server read package err= read header failed
			//如果希望服务器端一直读取客户端信息，下面的return去掉即可
			return err
		}

	}
}