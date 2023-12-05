package database

import (
	"ezgin/model"
	_ "github.com/go-sql-driver/mysql"
)

var UserDb *UserDB

func init() {
	UserDb = NewUserDB()
	UserDb.LoadFromFile("USER")
}

// CheckUser 检查用户名是否已存在
func CheckUser(username string) error {
	return UserDb.CheckUser(username)
}

// AddUser 添加用户
func AddUser(user *model.User) error {
	return UserDb.AddUser(user)
}

// QueryPwd 查询密码
func QueryPwd(username string) (int64, string, error) {
	return UserDb.QueryPwd(username)
}

func QueryRePwd(username string) (int64, string, error) {
	return UserDb.QueryRePwd(username)
}

// UpdatePwd 更新密码
func UpdatePwd(user *model.ParamResetPwdUser) error {
	return UserDb.UpdatePwd(user)
}
