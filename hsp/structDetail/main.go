package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func (rect Rect) updateLeftUpX(n int) {
	rect.leftUp.x = n
}
func main()  {
	r1 := Rect{Point{1,2}, Point{3, 4}}
	r1.updateLeftUpX(11);
	fmt.Println(r1)
	// r1 有四个int，在内存中是连续分布
	// 打印地址
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	// r2 有两个 *Point 类型，这个两个 *Point 类型的本身地址也是连续的
	// 但是他们指向的地址不一定是连续
	r2 := Rect2{&Point{10,20}, &Point{30,40}}

	//打印地址
	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p \n",
		&r2.leftUp, &r2.rightDown)

	//他们指向的地址不一定是连续...， 这个要看系统在运行时是如何分配
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p \n",
		r2.leftUp, r2.rightDown)
}