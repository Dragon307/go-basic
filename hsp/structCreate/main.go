package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main()  {
	// 1
	var p1 Person
	fmt.Println(p1)

	p2 := Person{"Alex", 20}
	fmt.Println(p2)

	// 2， 通过new 创建 &
	var p3 *Person = new(Person)
	// 因为p3是一个指针，因此标准的给字段赋值方式
	// (*p3).Name = "smith" 也可以这样写 p3.Name = "smith"

	// 原因: go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	// 会给 p3 加上 取值运算 (*p3).Name = "smith"
	(*p3).Name="sms"
	fmt.Println(p3)
	p3.Age = 12
	fmt.Println(p3)

	// 方式4
	var person *Person = &Person{}
	//因为person 是一个指针，因此标准的访问字段的方法
	// (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以 person.Name = "scott"
	// 原因和上面一样，底层会对 person.Name = "scott" 进行处理， 会加上 (*person)
	(*person).Name = "scott"
	person.Name = "scott~~"

	(*person).Age = 88
	person.Age = 10
	fmt.Println(*person)

}
