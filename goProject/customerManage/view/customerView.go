package main

import (
	"fmt"
	"go_code/goProject/customerManage/service"
)

type customerView struct {
	key int //接收客户输入
	loop bool //表示是否循环显示主菜单
	customerService *service.CustomerService
}
//显示主菜单
func (this *customerView) mainMenu() {
	for  {
		fmt.Println("................客户信息管理系统...................")
		fmt.Println("                1 添加客户")
		fmt.Println("                2 修改客户")
		fmt.Println("                3 删除客户")
		fmt.Println("                4 客户列表")
		fmt.Println("                5 退    出")
		fmt.Println("请选择(1-5): ")
		fmt.Scanln(&this.key)

		switch this.key {
		case 1:
			this.AddCustomer()
		case 2:
			this.AddCustomer()
		case 3:
			this.deleteCustomer()
		case 4:
			this.list()
		case 5:
			this.loop = false
		default:
			fmt.Println("选择错误请重新选择")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您已退出...")
}

//显示客户信息
func (this *customerView) list() {
	//首先获取当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("................客户列表................")
	fmt.Println("id\t姓名\t性别\t年龄\t手机号\t\t\t邮箱\t")
	for _,v := range customers{
		fmt.Println(v.ReturnUser())
	}
	fmt.Println("................结束................")
}

//添加客户
func (this *customerView) AddCustomer(){
	fmt.Println("............添加客户............")
	fmt.Println("姓名:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("性别:")
	var gander string
	fmt.Scanln(&gander)
	fmt.Println("年龄:")
	var age int
	fmt.Scanln(&age)
	fmt.Println("电话:")
	var phone string
	fmt.Scanln(&phone)
	fmt.Println("邮件:")
	var email string
	fmt.Scanln(&email)
	num := this.customerService.CustomerNum
	num++
	customer := service.ReturnNewCustomerInfo(num,name,gander,age,phone,email)
	res := this.customerService.AddCustomer(customer)
	if res {
		fmt.Println("添加成功~")
	}
}

//删除用户
func (this *customerView) deleteCustomer()  {
	fmt.Println(".........删除客户.......")
	fmt.Println("请选择待删除客户编号(输入-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	res := this.customerService.FindById(id)
	if res == -1 {
		fmt.Println("没有找到该id,请重新输入!")
		return
	}
	this.customerService.DeleteCustomerById(id)
	fmt.Println("刪除成功")
}
func main() {
	customer := customerView{
		0,
		true,
		service.NewCustomerService(),
	}
	customer.mainMenu()
}