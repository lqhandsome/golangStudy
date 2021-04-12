package service
import (
	"go_code/goProject/customerManage/model"

)
type CustomerService struct {
	customers []model.Customer
	CustomerNum int
}
//根据id查找
func (c CustomerService)FindById(id int) int {
	for i,v := range c.customers{
		if v.Id == id {
			return i
		}
	}
	return -1
}
//根据id删除用户
func (c *CustomerService)DeleteCustomerById(id int) bool  {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.customers = append(c.customers[:index],c.customers[index+1:]...)
	return true
}
//返回 *CustomerService
func NewCustomerService() *CustomerService	{

	CustomerService := &CustomerService{}
	CustomerService.CustomerNum =1
	//初始化一个用户
	customer := model.NewCustomer(
		1,"lq","man",23,"17538386243","lqhandsome@163.com")
	CustomerService.customers = append(CustomerService.customers,customer)
	return CustomerService

}
//返回一个客户信息
func ReturnNewCustomerInfo(id int,name string,gender string,age int,phone string,email string) model.Customer{

	return model.Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
//返回所有客户信息
func (c *CustomerService) List() []model.Customer{
	return c.customers
}

//添加客户
func (c *CustomerService) AddCustomer(customer model.Customer) bool{
	c.CustomerNum++
	c.customers = append(c.customers,customer)
	return true
}