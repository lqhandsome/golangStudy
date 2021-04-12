package utils

import "fmt"

type myFamilyAccount struct {
	key int
	loop bool
	balance float64
	money float64
	note string
	details string
}

func (this *myFamilyAccount)showDetail(){
	fmt.Println("当前收支明细记录")
	fmt.Println(this.details)
}

func (this *myFamilyAccount)income(){
	fmt.Println("本次收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",this.balance,this.money,this.note)
}
func (this *myFamilyAccount)pay(){
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)

	if this.balance < this.money {
		fmt.Println("支出失败，金额不足~")
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出\t%.2f\t\t%v\t\t%v",this.balance,this.money,this.note)
	}
}


func (this *myFamilyAccount) MainMenu(){

	for {
		fmt.Println(".........家庭收支账单......")
		fmt.Println("		1 收支明细")
		fmt.Println("		2 登记收入")
		fmt.Println("		3 登记支出")
		fmt.Println("		4 退出系统")
		fmt.Println("请选择<1-4>:")
		fmt.Scanln(&this.key)

		switch this.key {
		case 1:
			this.showDetail()
		case 2:
			this.income()
		case 3:
			this.pay()
		case 4:
			this.loop = true
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if this.loop {
			break
		}
	}
}

func TestmyFamilyAccount()*myFamilyAccount{
	return &myFamilyAccount{
		key :0,
		loop : false,
		balance : 100.0,
		money : 0.0,
		note : "",
		details : "收支\t账户金额\t收支金额\t说\t明",
	}
}
