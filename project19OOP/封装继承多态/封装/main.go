package main

import "fmt"
import "go_code/project19OOP/untils"
func main() {

	acccount := Account{
		"pf1111",
		"123456",
		100.69,
	}
	//acccount.SelectBalance()
	//acccount.SaveMoney(100)
	//fmt.Println(acccount.Balance)
	p := untils.Person{

	}
	fmt.Println(p.SetName())
	fmt.Println(acccount)

}
type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}


//查询余额
func (account *Account) SelectBalance(){
	var pwd string
	fmt.Scanln(&pwd)
	if account.Pwd != pwd{
		fmt.Println("密码不对")
		return
	}
	fmt.Println(account.Balance)
}

//存款
func (a *Account) SaveMoney(money float64){
	var pwd string
	fmt.Scanln(&pwd)
	if a.Pwd != pwd{
		fmt.Println("密码不对")
		return
	}
	a.Balance += money
	fmt.Println(a.Balance)
}

//取款
func (a *Account) withMoney(money float64,pwd string){
	if a.Balance < money || a.Pwd != pwd {
		fmt.Println("操作受限")
	}
	a.Balance -= money
}