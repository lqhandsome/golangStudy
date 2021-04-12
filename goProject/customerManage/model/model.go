package model

import "fmt"

//声明一个customer结构体
type  Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//返回格式话的客户信息
func (c *Customer) ReturnUser() (info string){

	info = fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",c.Id,c.Name,c.Gender,c.Age,c.Phone,c.Email)
	return

}
func NewCustomer (id int,name string,gender string,age int,
	phone string,email string)Customer {
	return Customer{
		id,
		name,
		gender,
		age,
		phone,
		email,
	}
}