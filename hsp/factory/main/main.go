package main

import (
	"fmt"
	"go-basic/hsp/factory/model"
)

func main()  {

	// student结构体字母是小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom", 98.9)
	fmt.Println(*stu)
	var stu2 = model.NewStudent("jon", 88)
	fmt.Println(stu2)
	fmt.Println("name=", stu.Name, " score=", stu.GetScore(),
		stu2.Name, stu2.GetScore())

}