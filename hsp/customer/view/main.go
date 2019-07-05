package main

import (
	"fmt"
	"go-basic/hsp/customer/model"
	"go-basic/hsp/customer/service"
)

type CustomerView struct {
	// 接受用户输入
	key string
	//定义一个CustomerService 字段，主要用于完成对客户信息的各种操作。
	customerService *service.CustomerService
}

func (cv *CustomerView) mainMenu()  {
forloop:
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println();
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退          出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			cv.add()
		case "2":

		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			break forloop
		default:
			fmt.Println("输入错误");
		}
	}
}

// 显示客户信息
func (cv *CustomerView) list()  {
	customerList := cv.customerService.List()

	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	//遍历
	for i := 0; i < len(customerList); i++ {
		fmt.Println(customerList[i].GetInfo())
	}
	fmt.Println("---------------------------客户列表完成---------------------------")
}

// 添加用户信息
func (cv *CustomerView) add()  {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	age := 0
	fmt.Println("年龄：")
	fmt.Scanln(&age)
	fmt.Println("电话：");
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：");
	email := ""
	fmt.Scanln(&email)

	// 创建一个 Customer 对象
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if cv.customerService.Add(customer) {
		fmt.Println("---------------------添加客户成功---------------------")
	} else {
		fmt.Println("---------------------添加客户失败---------------------")
	}
}

func (cv *CustomerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)")
	id := 0
	fmt.Scanln(&id)
	//如果用户输入-1
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N)：")

	choice := ""
	fmt.Scanln(&choice) // 可以

	if choice == "Y" || choice == "y"  {

		if cv.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败,id不存在---------------------")
		}
	}
}

func main()  {

	customerView := CustomerView{}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}