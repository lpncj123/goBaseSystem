package main

import "fmt"

func main() {
	key := ""
	loop := true
	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t账户金额\t收支金额\t说   明"
	flag := false
	for {
		fmt.Println("------------家庭收支记账软件-------------------")
		fmt.Println("------------1、收支明细-------------------")
		fmt.Println("------------2、登记收入-------------------")
		fmt.Println("------------3、登记支出-------------------")
		fmt.Println("------------4、退出软件-------------------")
		fmt.Println("请选择1到4")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if !flag {
				fmt.Println("当前还没有明细，添加一笔")
				break
			}
			fmt.Println("----当前收支明细记录----")
			fmt.Println(details)
		case "2":
			fmt.Println("输入收入金额")
			fmt.Scanln(&money)
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			balance += money
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("输入支出金额")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			balance -= money
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
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
				loop = false
			}

		default:
			fmt.Println("请输入正确的编号")
		}
		if !loop {
			fmt.Println("退出系统")
			break
		}
	}
}
