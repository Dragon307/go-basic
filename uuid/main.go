package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)
// uuid是Universally Unique Identifier的缩写，即通用唯一识别码。
func main() {
	// 创建
	u1, _:= uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	// 解析
	u2, err := uuid.FromString("10a4005c-5ab9-452f-ab64-5b3aa8e7bb7b");

	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
