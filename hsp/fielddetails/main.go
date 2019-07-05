package main

import (
	"fmt"
)

// 如果结构体的字段类型是：指针，slice，和map的零值都是 nil，即还没有分配空间
// 如果需要使用这样的字段，需要先 make， 才能使用

type Person struct {
	Name string
	Age int
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}

type Monster struct {
	Name string
	Age int
}

func main()  {
	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	// 使用 slice，再次说明，一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100 // ok

	// 使用map，一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"
	fmt.Println(p1)

	// 不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改
	// 不影响另外一个，结构体是值类型
	var monster1 Monster
	monster1.Name = "mon1"
	monster1.Age = 500

	monster2 := monster1
	monster2.Name = "mon2"

	fmt.Println("monster1=",monster1)
	fmt.Println("monster2=", monster2)
}