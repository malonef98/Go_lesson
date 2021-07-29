package utils

import "fmt"

//使用面向对象的方法完成

type FamilyAccount struct {
	//保留用户输入字段
	key string

	//循环标志
	loop bool

	banlance float64

	money float64

	note string

	flag bool

	details string

}


func (this *FamilyAccount) showDetails()  {
	fmt.Println("-------- 当前收支状况 --------")
	if this.flag {
		fmt.Println(this.details)
	}else {
		fmt.Println("当前没有收支情况，请新添加")
	}
}

func (this *FamilyAccount) pay()  {
	fmt.Println("输入本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.banlance {
		fmt.Println("支出的金额大于余额！！！")
		return
	}
	fmt.Println("输入本次支出说明：")
	fmt.Scanln(&this.note)
	this.banlance -= this.money
	//将收入情况记录在details中
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v",this.banlance,this.money,this.note)
	this.flag = true
}

func (this *FamilyAccount) income()  {
	fmt.Println("输入本次收入金额：")
	fmt.Scanln(&this.money)
	fmt.Println("输入本次收入说明：")
	fmt.Scanln(&this.note)
	this.banlance += this.money
	//将收入情况记录在details中
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v",this.banlance,this.money,this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
	fmt.Println("退出软件")
	fmt.Println("你确定要退出软件吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" {
			fmt.Println("退出成功")
			this.loop = false
			break
		} else if choice == "n" {
			fmt.Println("退出失败")
			break
		} else {
			fmt.Println("输入错误，请重新输入 y/n")
		}
	}
}


func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key: "",
		loop: true,
		banlance: 0,
		money: 0,
		note: "",
		details: "\n收支\tt账户金额\t收支金额\t说  明",
		flag: false,
	}
}


func (this *FamilyAccount) MainMenu()  {
	for  {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("---------- 1  收支明细   ----------")
		fmt.Println("---------- 2  登记收入   ----------")
		fmt.Println("---------- 3  登记支出   ----------")
		fmt.Println("---------- 4  退出软件   ----------")
		fmt.Println("---------  请选择1 - 4   ----------")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
		}
		if this.loop != true {
			break
		}
	}
}
