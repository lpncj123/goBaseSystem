package utils

import "fmt"

type MyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
	flag    bool
}

func NewMyAccount() *MyAccount {
	return &MyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说   明",
		flag:    false,
	}
}

func (account *MyAccount) showDetails() {
	if !account.flag {
		fmt.Println("当前还没有明细，添加一笔")
	} else {
		fmt.Println("----当前收支明细记录----")
		fmt.Println(account.details)
	}
}

func (account *MyAccount) shouRu() {
	fmt.Println("输入收入金额")
	fmt.Scanln(&account.money)
	fmt.Println("本次收入说明")
	fmt.Scanln(&account.note)
	account.balance += account.money
	account.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", account.balance, account.money, account.note)
	account.flag = true
}

func (account *MyAccount) zhiChu() {
	fmt.Println("输入支出金额")
	fmt.Scanln(&account.money)
	if account.money > account.balance {
		fmt.Println("余额不足")
		//break
	}
	fmt.Println("本次支出说明")
	fmt.Scanln(&account.note)
	account.balance -= account.money
	account.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", account.balance, account.money, account.note)
	account.flag = true
}

func (account *MyAccount) tuiChu() {
	fmt.Println("确认退出吗，请输入y/n")
	flag := ""
	for {
		fmt.Scanln(&flag)
		if flag == "y" || flag == "n" {
			break
		}
		fmt.Println("输入有误，请重新输入")
	}
	if flag == "y" {
		account.loop = false
	}
}

func (account *MyAccount) MainMenu() {
	for {
		fmt.Println("------------家庭收支记账软件-------------------")
		fmt.Println("------------1、收支明细-------------------")
		fmt.Println("------------2、登记收入-------------------")
		fmt.Println("------------3、登记支出-------------------")
		fmt.Println("------------4、退出软件-------------------")
		fmt.Println("请选择1到4")
		fmt.Scanln(&account.key)
		switch account.key {
		case "1":
			account.showDetails()
		case "2":
			account.shouRu()
		case "3":
			account.zhiChu()
		case "4":
			account.tuiChu()
		default:
			fmt.Println("请输入正确的编号")
		}
		if !account.loop {
			fmt.Println("退出系统")
			break
		}
	}

}
