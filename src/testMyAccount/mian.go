package main

import "fmt"

func main() {
	key := ""
	loop := true
	//账户变量
	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t账户金额\t收支金额\t说  明"
	details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)

	flag := false


	for  {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("---------- 1  收支明细   ----------")
		fmt.Println("---------- 2  登记收入   ----------")
		fmt.Println("---------- 3  登记支出   ----------")
		fmt.Println("---------- 4  退出软件   ----------")
		fmt.Println("---------  请选择1 - 4   ----------")

		fmt.Scanln(&key)

		switch key {
		case "1":
		   	fmt.Println("-------- 当前收支状况 --------")
			if flag {
				fmt.Println(details)
			}else {
				fmt.Println("当前没有收支情况，请新添加")
			}
		case "2":
			fmt.Println("输入本次收入金额：")
			fmt.Scanln(&money)
			fmt.Println("输入本次收入说明：")
			fmt.Scanln(&note)
			balance += money
			//将收入情况记录在details中
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
			flag = true
		case "3":
			fmt.Println("输入本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("支出的金额大于余额！！！")
				break
			}
			fmt.Println("输入本次支出说明：")
			fmt.Scanln(&note)
			balance -= money
			//将收入情况记录在details中
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
			flag = true
		case "4":
			fmt.Println("退出软件")
			fmt.Println("你确定要退出软件吗？y/n")
			choice := ""
			for  {
				fmt.Scanln(&choice)
				if choice == "y" {
					fmt.Println("退出成功")
					loop = false
					break
				}else if choice == "n"  {
					fmt.Println("退出失败")
					break
				}else {
					fmt.Println("输入错误，请重新输入 y/n")
				}
			}
		default:
		}

		if loop != true {
			break
		}
	}
}
