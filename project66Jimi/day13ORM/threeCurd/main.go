package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type User struct {
	ID int64
	Name string
	Age int64
}

func main() {
	//连接数据库
	dsn := "root:123456@(127.0.0.1:33222)/lq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate(&User{})

	//创建一个实例
	u := User{
		1,
		"lq",
		24,
	}
	db.Create(&u)
	u2 := User{
		Name: "lqq",
		Age: 244,
	}
	db.Create(&u2)
}
