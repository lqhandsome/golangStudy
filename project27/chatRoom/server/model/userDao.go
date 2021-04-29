package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)
//定义一个UserDao 结构体
//完成对User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个userDao实例
func NewUserDao(pool *redis.Pool)(userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}
//1.根据id返回用户信息
func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error){
	user = &User{}
	msg,err := redis.String(conn.Do("hget","user",id))
	if err !=nil {
		//没有找到用户
		if err != redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
			return
		}
	}

	err = json.Unmarshal([]byte(msg),user)
	if err != nil {
		fmt.Println("json反序列化出错",err)
		return
	}
	return
}

//2.登录校验
func(this *UserDao)  Login(userId int,userPwd string)(user *User,err error) {
	user = &User{}
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn,userId)
	if err != nil {
		return
	}
	//能够取出用户，再判断用户密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}