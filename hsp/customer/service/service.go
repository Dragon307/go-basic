package service

import (
	"go-basic/hsp/customer/model"
)

type CustomerService struct {
	// 定义一个切片，可以存放客户信息
	customers []*model.Customer
	// 定义客户的实际个数
	customerNum int
}

// 先创建一个Customer对象，放到 service 的 Customers 切片中作为测试数据
func NewCustomerService() *CustomerService {

	customerService := &CustomerService{}

	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男",
		10, "999", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

// 返回客户信息数组
func (cs *CustomerService) List() []*model.Customer {
	return cs.customers
}

// 完成添加客户的功能
func (cs *CustomerService) Add(customer *model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// 删除一个客户
func (cs *CustomerService) Delete(id int) bool {

	// 先根据id去得到该id的客户对应元素的下标
	index := cs.FindById(id)

	if index == -1 {
		return false
	}

	// 删除切片中的一个元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// 根据id查找客户, 返回对应的下标

func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < cs.customerNum; i++ {
		if cs.customers[i].Id == id {
			index = i
			break
		}
	}
	return index
}