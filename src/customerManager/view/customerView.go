package main

import (
	"awesomeProject2/src/customerManager/model"
	"awesomeProject2/src/customerManager/service"
	"fmt"
)

type customerView struct {
	loop bool
	key string
	customerservice *service.CustomerService
}

func (this *customerView) list()  {
	//获取当前所有客户信息  在切片中
	customers := this.customerservice.List()

	fmt.Println("--------------- 客户列表 ---------------")
	fmt.Println("ID\t名字\t年龄\t性别\t电话\t邮箱")
	for i := 0 ; i < len(customers); i++ {
		fmt.Println(customers[i].GetIndo())
	}
}

func (this *customerView) add()  {
	fmt.Println("---------- 添加客户 ----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电子邮件：")
	email := ""
	fmt.Scanln(&email)

	//id系统分配
	customer := model.NewCustomer2(name,gender,age,phone,email)

	this.customerservice.Add(customer)

}

func (this *customerView) delete()  {
	fmt.Println("---------- 删除客户 ----------")
	fmt.Println("输入待删除的客户id (-1退出) ")
	id := -1
	fmt.Scanln(&id)
	if id == -2 {
		return
	}
	fmt.Println("确认是否删除 y/n?")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		this.customerservice.Delete(id)
	} else if choice == "n" || choice == "N"  {
		fmt.Println("放弃删除")
		return
	}
}

func (this *customerView) update()  {
	fmt.Println("---------- 修改客户 ----------")
	fmt.Println("需要修改的用户ID：")
	id := -1
	fmt.Scanln(&id)
	fmt.Println("修改的姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("修改的性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("修改的年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("修改的电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("修改的电子邮件：")
	email := ""
	fmt.Scanln(&email)

	this.customerservice.Update(id,name,gender,age,phone,email)
}

func (this *customerView) mainMenu()  {
	for  {
		fmt.Println("\n---------------- 客户信息管理软件 ----------------")
		fmt.Println("----------------   1 添加客户    ----------------")
		fmt.Println("----------------   2 修改客户    ----------------")
		fmt.Println("----------------   3 删除客户    ----------------")
		fmt.Println("----------------   4 客户列表    ----------------")
		fmt.Println("----------------   5 退出软件    ----------------")
		fmt.Println("请选 1 - 5")

		fmt.Scanln(&this.key)

		switch this.key{
		case "1":
			fmt.Println("添加客户")
			this.add()
		case "2":
			fmt.Println("修改客户")
			this.update()
		case "3":
			fmt.Println("删除客户")
			this.delete()
		case "4":
			fmt.Println("客户列表")
			this.list()
		case "5":
			fmt.Println("退出软件")
			this.loop = false
		default:

		}
		if this.loop != true {
			break
		}
	}
}


func main() {
	v := customerView{
		loop: true,
	}
	v.customerservice = service.NewCustomerService()
	v.mainMenu()
}