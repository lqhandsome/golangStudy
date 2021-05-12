package main

import (
	"database/sql"
	"time"

	//gg "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

//创建用户信息表
type UserInfo struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号（member number）唯一且不为空
	num          int     `gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Address      string  `gorm:"index:addr"`      //创建addr的索引
	IgnoreMe     string  `gorm:"-"`               //忽略本字段

}

func main() {
	dsn := "root:123456@(127.0.0.1:33222)/lq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	db.AutoMigrate(UserInfo{})

	//手动创建表
	db.Table("aaa").Raw("select *from aaa")
}

//收到指定表名
func (this UserInfo) TableName() string {

	return "userInfo"
}
