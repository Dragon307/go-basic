package main

import "fmt"

type Person struct {
	Name string
}

// 给Person 结构体添加 speak 方法
func (p Person) speak()  {
	fmt.Println(p.Name, "是一个goodman~")
}
// 给 Person 结构体添加 jisuan 方法，可以计算 1+.. + 1000 的结果
// 说明方法体内可以函数一样，进行各种运算

func (p Person) jisuan()  {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=",res)
}

// 给 Person 结构体 jisuan2 方法，该方法可以接收一个参数n，计算从1+...+n 的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

// 给Person结构体添加getSum方法
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 给Person类型绑定一方法
func (person Person) test() {
	person.Name = "jack"
	fmt.Println("test() name=", person.Name)
}

func main()  {
	var p Person
	p.Name = "tom"
	p.test()
	fmt.Println("main() p.Name=", p.Name) //输出 tom

	p.speak()
	p.jisuan()
	p.jisuan2(20)

	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println("res=", res)
}