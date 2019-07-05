package main

import "fmt"

type Point struct {
	 x int
	 y int
}

func main()  {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	// 如何将 a 赋给一个Point变量
	var b Point
	b = a.(Point)
	fmt.Println(b)
	fmt.Printf("%T %T\n", b, a)

	// 类型断言 带检测
	var x interface{}
	var b2 float32 = 2.1
	x = b2 // 空接口，可以接受任意类型

	if y, ok := x.(float32); ok {
		fmt.Println("convert success", ok)
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("convert error")
	}
}
