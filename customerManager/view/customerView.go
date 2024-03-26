package main

import (
	"fmt"
	"gobase/customerManager/modal"
	"gobase/customerManager/service"
)

type CustomerView struct {
	//接收用户输入
	key string
	//表示是否循环
	loop            bool
	customerService *service.CustomerService
}

// 显示所有客户端信息
func (customerView *CustomerView) list() {
	customers := customerView.customerService.List()
	fmt.Println("--------------------客户列表-------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	//for i, customer := range customers {
	for _, customer := range customers {
		fmt.Println(customer.GetInfo())
	}
	fmt.Println("--------------------客户列表完成-------------------")
}
func (customerView *CustomerView) add() {
	fmt.Println("------------------添加客户--------------------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := modal.NewCustomer2(name, gender, age, phone, email)
	if customerView.customerService.Add(customer) {
		fmt.Println("------添加成功---------")
	} else {
		fmt.Println("------添加失败---------")
	}

}
func (customerView *CustomerView) update() {
	fmt.Println("------------------修改客户--------------------")
	fmt.Println("Id")
	Id := 0
	fmt.Scanln(&Id)
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := modal.NewCustomer(Id, name, gender, age, phone, email)
	if customerView.customerService.Update(customer) {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")
	}

}
func (customerView *CustomerView) deleteById() {
	fmt.Println("------------------删除客户--------------------")
	fmt.Println("要删除ID，-1退出删除")
	Id := -1
	fmt.Scanln(&Id)
	if Id == -1 {
		return
	}
	fmt.Println("是否确认删除y/n")
	flag := ""
	fmt.Scanln(&flag)
	if flag == "n" {
		return
	}
	if customerView.customerService.DeleteById(Id) {
		fmt.Println("------删除成功---------")
	} else {
		fmt.Println("------删除失败---------")
	}
}

// 显示主菜单
func (customerView *CustomerView) mainMenu() {
	for {
		fmt.Println("--------------客户信息管理软件-----------------")
		fmt.Println("--------------1、添加用户-----------------")
		fmt.Println("--------------2、修改用户-----------------")
		fmt.Println("--------------3、删除用户-----------------")
		fmt.Println("--------------4、查询用户-----------------")
		fmt.Println("--------------5、退出-----------------")
		fmt.Println("--------------请选择1-5-----------------")
		fmt.Scanln(&customerView.key)
		switch customerView.key {
		case "1":
			customerView.add()
		case "2":
			customerView.update()
		case "3":
			customerView.deleteById()
		case "4":
			customerView.list()
		case "5":
			customerView.loop = false
		default:
			fmt.Println("输入有误，请重新输入")

		}
		if !customerView.loop {
			break
		}
	}
	fmt.Println("你退出了系统")
}
func main() {
	customerView := CustomerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
