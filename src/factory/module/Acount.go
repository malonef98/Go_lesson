package module

import "fmt"

//结构体内的字段都是小写，封装数据，外部包无法直接获取
type Account struct {
	account string
	balance float64
	pwd string
}

//编写SetXX函数，让外部使用者能够对封装的字段进行操作
func (account *Account) SetAccount (a string)  {
	if len(a) < 16 {
		fmt.Println("账户非法，长度不足16位")
		return
	}
	account.account = a
	fmt.Println("账户设置成功")
	return
}

func (account *Account) SetPwd (pwd string)  {
	if len(pwd) != 6 {
		fmt.Println("密码非法，长度必须为6位")
		return
	}
	account.pwd = pwd
	fmt.Println("账户密码设置成功")
	return
}

func (account *Account) SetBalance (b float64)  {
	if b < 20 {
		fmt.Println("余额必须大于20")
		return
	}
	account.balance = b
	fmt.Println("账户余额成功")
	return
}