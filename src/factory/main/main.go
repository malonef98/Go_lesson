package main

import (
	"awesomeProject2/src/factory/module"
)

func main() {
	var acc1 module.Account
	acc1.SetAccount("ss111111111111111")
	acc1.SetPwd("111111")
	acc1.SetBalance(5000)
}
