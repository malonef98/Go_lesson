package model

import "fmt"

//表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func (this Customer) GetIndo()  string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,this.Age,this.Gender,this.Phone,this.Email)
	return info
}

//使用工厂模式，返回customer实例
func NewCustomer(id int, name string,gender string, age int, phone string, email string) Customer {
	return Customer{
		Id: id,
		Gender: gender,
		Age: age,
		Name: name,
		Phone: phone,
		Email: email,
	}
}

//不带ID
func NewCustomer2(name string,gender string, age int, phone string, email string) Customer {
	return Customer{
		Gender: gender,
		Age: age,
		Name: name,
		Phone: phone,
		Email: email,
	}
}


