package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//方法
// 1 存款
func (account *Account) Deposite(money float64, pwd string){
	//判断密码
	if pwd != account.Pwd{
		println("密码错误")
		return
	}

	//判断输入金额
	if money < 0  {
		println("输入金额错误")
		return
	}

	account.Balance += money
	fmt.Println("存款成功")
}

//  2 取款
func (account *Account) WithDraw (money float64, pwd string) {
	//判断密码
	if pwd != account.Pwd{
		println("密码错误")
		return
	}

	//判断输入金额
	if money < 0  {
		println("输入金额错误")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功")
}

// 3 查询余额
func (account *Account) Query (pwd string) {
	//判断密码
	if pwd != account.Pwd {
		println("密码错误")
		return
	}

	fmt.Println("余额为：",account.Balance)

}

func main() {

	account := Account{
		AccountNo: "GS111",
		Pwd: "666",
		Balance: 100000,
	}

	account.Deposite(1500,"666")
	account.Query("666")
	account.WithDraw(5500,"666")
	account.Query("666")

}

