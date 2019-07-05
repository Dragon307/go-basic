package main

import "fmt"

type Stu struct {
	Name string
	Age int
}

func main()  {

	var stu1 = Stu{"s1", 19}
	stu2 := Stu{"s2", 19}

	stu3 := Stu{
		Name: "jack",
		Age: 30,
	}

	// 返回结构体的指针类型
	var stu4 *Stu = &Stu{"s4", 22}

	fmt.Println(stu1,"\n", stu2,"\n", stu3,"\n", stu4)
}
