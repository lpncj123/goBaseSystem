package service

import (
	"fmt"
	"gobase/customerManager/modal"
)

// CustomerService 该service，完成对Customer的操作，包括
// 增删改查
type CustomerService struct {
	customers []modal.Customer
	//	声明一个字段，表示当前切片含有多少个客户，还可以作为客户id
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := modal.NewCustomer(1, "张三", "男", 18, "15048293147", "257@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}
func (customerService *CustomerService) List() []modal.Customer {
	return customerService.customers
}

func (customerService *CustomerService) Add(customer modal.Customer) bool {
	customerService.customerNum++
	customer.Id = customerService.customerNum
	customerService.customers = append(customerService.customers, customer)
	return true
}

func (customerService *CustomerService) DeleteById(Id int) bool {
	index := customerService.FindCustomerById(Id)
	if index == -1 {
		return false
	} else {
		customerService.customers = append(customerService.customers[:index], customerService.customers[index+1:]...)
	}
	return true
}

func (customerService *CustomerService) FindCustomerById(Id int) int {
	index := -1
	for i, customer := range customerService.customers {
		if customer.Id == Id {
			index = i
			break
		}
	}
	return index
}

func (customerService *CustomerService) Update(customer modal.Customer) bool {
	index := customerService.FindCustomerById(customer.Id)
	if index == -1 {
		fmt.Println("记录不存在")
		return false
	}
	customerService.customers[index] = customer
	return true
}
