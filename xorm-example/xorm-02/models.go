package main

import (
	"errors"
	_ "github.com/mysql"
	"github.com/xorm"
	"log"
	"os"
	_ "reflect"
)

// 银行账户
type Account struct {
	Id int64
	Name string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // 乐观锁
}

func (a *Account) BeforeInsert()  {
	log.Printf("Before insert: %s\n", a.Name)
}

func (a *Account) AfterInsert()  {
	log.Printf("After insert: %s\n", a.Name)
}

// ORM 引擎
var x *xorm.Engine

func init()  {
	// 创建 ORM 引擎与数据库
	var err error
	x, err = xorm.NewEngine("mysql", "root:Liuji0108@tcp(localhost:3306)/bank?charset=utf8")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

	// 记录日志
	f, err := os.Create("sql.log")
	if err != nil {
		log.Fatalf("Fail to create log file: %v\n", err)
		return
	}
	x.SetLogger(xorm.NewSimpleLogger(f))

	// 设置默认 LRU 缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	x.SetDefaultCacher(cacher)

}

// 创建新的账户
func NewAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 获取账户数量
func getAccountCount() (int64, error) {
	return x.Count(new(Account))
}

// 获取账户信息
func getAccount(id int64) (*Account, error) {
	a := &Account{}
	has, err := x.Id(id).Get(a)

	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account does not exist!")
	}
	return a, nil
}