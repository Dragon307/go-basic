package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name string  `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main()  {
	var a A
	var b B
	b.Num = 12
	a = A(b) // 可以转换，但是有要求，就是结构体的字段要完全一致
	b.Num = 13
	fmt.Println(a,b)

	// 1 创建 Monster 变量
	monster := Monster{"alex", 190, "hahha"}

	// 2 将 monster 变量序列化为 json 格式字串
	// json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}