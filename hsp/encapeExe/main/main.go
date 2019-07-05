package main

import (
	"fmt"
	"go-basic/hsp/encapeExe/model"
)

func main() {
	// 创建一个account变量
	account := model.NewAccount("jsa09876", "000000", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
		account.WithDraw(20.0, "000000")
		account.Query("000000")
	} else {
		fmt.Println("创建失败")
	}

}