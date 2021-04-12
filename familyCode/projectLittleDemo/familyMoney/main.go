package main

import (
	"fmt"
)

func main()  {
	key := 0
	loop := true
	balance := 100.0
	money := 0.0
	note := ""
	details := "收支\t账户金额\t收支金额\t说	明"

	//现实主菜单
	for {
		fmt.Println(".........家庭收支账单......")
		fmt.Println("		1 收支明细")
		fmt.Println("		2 登记收入")
		fmt.Println("		3 登记支出")
		fmt.Println("		4 退出系统")
		fmt.Println("请选择<1-4>:")
		fmt.Scanln(&key)

		switch key {
			case 1:
				fmt.Println("当前收支明细记录")
				fmt.Println(details)
			case 2:
				fmt.Println("本次收入金额")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("本次收入说明")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",balance,money,note)
		case 3:
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)

			if balance < money {
				fmt.Println("支出失败，金额不足~")
			} else {
				balance -= money
				fmt.Println("本次支出说明")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n支出\t%.2f\t\t%v\t\t%v",balance,money,note)
			}
		case 4:
			loop = false
		}

		if !loop {
			break
		}
	}
}
