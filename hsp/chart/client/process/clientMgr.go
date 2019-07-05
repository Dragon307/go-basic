package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-basic/hsp/chart/common"
	"go-basic/hsp/chart/client/utils"
	"net"
)

// 客户端这边用于维护客户信息的（比如列表）
var onlineUserMap map[int]*common.User = make(map[int]*common.User, 10)
var UserId int

// 输出当前在线用户信息列表
func outputUserOnLine()  {
	fmt.Println("在线用户列表如下：")
	for id, _ := range onlineUserMap {
		if id == UserId {
			continue
		}
		fmt.Println("在线用户id:\t", id)
	}
}

func enterTalk(_ net.Conn, str string) (err error) {

	//链接到Redis
	conn, err := utils.ConnNet()
	if err != nil {
		fmt.Println("error in connect", err)
	}


	var msg common.Message
	msg.Type = common.TalkMesType

	var talkMsg common.TalkMes
	talkMsg.Id = UserId
	talkMsg.Msg = str

	// 将 msg 序列化，便于网络传输
	data, err := json.Marshal(talkMsg)
	if err != nil {
		return err
	}

	// 将消息的内容（data）放入到msg的Data字段
	msg.Data = string(data)

	// 然后对msg进行序列化
	data, err = json.Marshal(msg)
	if err != nil {
		return err
	}

	var buf [4]byte
	// 消息长度
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], packLen)

	// 先发送消息的长度，便于服务器接收
	n, err := conn.Write(buf[:])

	if err != nil || n != 4 {
		fmt.Println("enterTalk write data  failed", err)
		return
	}
	//发送消息给服务器
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	// 读取服务器返回的消息包
	//tf := &utils.Transfer{
	//	Conn: conn,
	//}
	//msg, err = tf.ClientReadPackage()
	//
	//if err != nil {
	//	fmt.Println("readPackage failed, err=", err)
	//}
	//fmt.Println(msg, "msg")
	return
}

// 更新客户端 当前在线用户信息列表
//注意： 因为服务器端会返回一个新的登录用户信息包，因此更新返回这个用户的状态即可
//(1) 一种情况是新登录的用户 [加入到 onlineUserMap]
//(2) 一种情况是登录的用户下线了 [更新 ]

func updateUserStatus(userStatusNotifyMes common.UserStatusNotifyMes)  {
	//看看是否已经在onlineUserMap了
	user, ok := onlineUserMap[userStatusNotifyMes.UserId]

	if !ok {
		//如果不在onlineUserMap,就先创建一个
		user = &common.User{}
		user.Id = userStatusNotifyMes.UserId
	}
	//更新状态[1.在线 2.离线]
	user.Status = userStatusNotifyMes.Status
	//更新onlineUserMap
	onlineUserMap[user.Id] = user
	//输出最新的在线用户列表
	outputUserOnLine()

}

