package main

import (
	"fmt"
	//gg "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)
//创建用户信息表
type UserInfo struct {
	ID int
	Name string
	Gender string
	Hobby string
}
func main() {
	dsn := "root:123456@(127.0.0.1:33222)/lq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	db.AutoMigrate(&UserInfo{})

	//创建一行数据
	//u1 := &UserInfo{
	//	1,
	//	"lq",
	//	"man",
	//	"ITT",
	//}
	//db.Create(u1)

	//查询
	var u UserInfo
	db.First(&u)
	fmt.Println(u)

	//修改
	db.Model(&u).Update("hobby","ITT")

	//删除
	db.Delete(&u)
}
