package main

import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)

	return infoStr
}

type Box struct {
	len float64
	width float64
	height float64
}

// 声明一个方法获取立方体的体积
func (box *Box) getVolume() float64 {
	return box.len * box.width * box.height
}

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) showPrice()  {
	if visitor.Age >= 90 || visitor.Age <=8 {
		fmt.Println("考虑到安全，就不要玩了")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元 \n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费 \n", visitor.Name, visitor.Age)
	}
}

func main()  {
	var stu = Student{
		name : "tom",
		gender : "male",
		age : 18,
		id : 1000,
		score : 99.98,
	}
	fmt.Println(stu.say())

	//测试代码
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0
	volumn := box.getVolume()
	fmt.Printf("体积为=%.2f", volumn)

	// 测试
	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序....")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}