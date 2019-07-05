package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

// 演示将json字符串，反序列化成 struct
func unmarshalStruct()  {
	str := "{\"Name\":\"牛魔王~~~\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	// 定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}

	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)

}

// 将map进行序列化
func testMap() string {
	// 定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿~~~~~~"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化的结果
	//fmt.Printf("a map 序列化后=%v\n", string(data))
	return string(data)
}

// 演示将json字符串，反序列化成map
func unmarshalMap()  {
	str := testMap()
	var a map[string]interface{}

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

// 演示将json字符串，反序列化成切片
func unmarshalSlice()  {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	// 定义一个slice
	var slice []map[string]interface{}
	// 反序列化，不需要make，因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}


func main()  {
	unmarshalStruct()
	testMap()
	unmarshalMap()
	unmarshalSlice()
}