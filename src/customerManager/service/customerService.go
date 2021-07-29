package service

import (
	"awesomeProject2/src/customerManager/model"
	"fmt"
)

//完成对Customer的操作
//增删改查
type CustomerService struct {
	customers []model.Customer
	//字段 表示当前切片含有有多少个用户
	//该字段后面可以作为新客户的 Id+1
	customerNum int
}

func (this CustomerService) List() []model.Customer {
	return this.customers
}

//编写一个方法 可以返回 *customerService指针
func NewCustomerService() *CustomerService{

	CustomerService := &CustomerService{}
	CustomerService.customerNum = 0
	//customer := model.NewCustomer(1,"张三","男",20,"112","zs@sohu.com")
	//添加用户
	//CustomerService.customers = append(CustomerService.customers,customer)
	return CustomerService
}

//用户添加
func (this *CustomerService) Add(customer model.Customer) bool {
	//确定一个ID的规则  添加顺序
	//指针引用  customerNum才会累加
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//修改客户信息
func (this *CustomerService) Update(id int,name string,gender string,age int,phone string,email string) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("未找到对应ID的客户")
		return false
	}
	//赋值地址，才能真正的改变切片数组
	customer := &this.customers[index]
	customer.Name   = name
	customer.Gender = gender
	customer.Age    = age
	customer.Phone  = phone
	customer.Email  = email
	return true
}

//删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("未找到对应ID的客户")
		return false
	}
	//如何从切片中删除一个元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	fmt.Println("删除成功")
	return true
}

func (this *CustomerService) FindById(id int)  int{
	index := -1
	for i := 0 ; i < len(this.customers) ; i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}