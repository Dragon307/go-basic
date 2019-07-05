package process

import "fmt"

// 用户的管理操作
// 比如对在线用户的管理

type ClientMgr struct {
	onlineUsers map[int]*UserProcessor
}

var (
	clientMgr *ClientMgr
)
// 完成初始化， 切片不初始化，不能使用
func init() {
	clientMgr = &ClientMgr{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

// 一个 ClientProcessor 实例，就对应一个登录的用户
func (p *ClientMgr) AddClient(userId int, client *UserProcessor)  {
	p.onlineUsers[userId] = client
}

func (p *ClientMgr) GetClient(userId int) (client *UserProcessor, err error) {
	client, ok := p.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("id= %d 的用户,不存在..", userId)
		return
	}

	return
}

func (p *ClientMgr) GetAllUsers() map[int]*UserProcessor {
	return p.onlineUsers
}

// 当有一个用户离线后，就从 onlineUsers 切片中删除掉
func (p *ClientMgr) DelClient(userId int)  {
	delete(p.onlineUsers, userId)
}