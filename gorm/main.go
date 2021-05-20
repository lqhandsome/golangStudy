package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
type Course struct {
	Cno string `gorm:"primaryKey"`
	Cname string
	Tno string
}

type User struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	Name string
	Sex string
}
func main() {
	dsn := "lq:lqhandsome@(47.107.157.94:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//禁止添加给表名添加s
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	var course *Course = &Course{}
	//user := &User{}
	db.First(course)
	//db.AutoMigrate(&User{})
	user := []User{{Name:"lq",Sex: "man"},{Name:"lq",Sex: "man"}}
	db.Create(&user)

	aa := &User{
		Name: "aa",
		Sex: "woman",
	}
	db.Create(aa)
	db.CreateInBatches(user,10)
	for i,user := range user {
		fmt.Println(user,i)
	}
	//user := make(map[string]interface{})
	//db.First(user)
	fmt.Println(*course)
}
