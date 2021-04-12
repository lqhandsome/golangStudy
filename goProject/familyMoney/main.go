package main

import "fmt"

func main() {
	//保存用户输入的选项
	key := ""
	//控制是否退出for
	loop := true
	//定义账户余额
	//balance := 100.0
	//每次收支的金额
	//money := 0.0
	//每次收支的说明
	//note := ""
	//收支的记录
	noteDetail := "收支\t账户金额\t收支金额\t说	明"
	for {
		fmt.Println("............家庭记账软件................")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出")

		fmt.Print("请选择<1-4>")
		fmt.Scanln(&key)
		switch key {
			case "1":
				fmt.Println("................当前收支明细记录..............")
				fmt.Println(noteDetail)
			case "2":
			case "3":
				fmt.Println("登记支出")
			case "4":
				loop = false

		}
		if !loop {

			break
		}
	}
	fmt.Println("您已经退出记账软件...")
}
