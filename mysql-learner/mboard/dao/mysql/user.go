package mysql

import (
	"database/sql"
	"ezgin/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	IfExistUserStr = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?);"
	AddUserStr     = "INSERT INTO users(username,password,re_password) VALUES(?,?,?);"
	QueryPwdStr    = "SELECT uid,password FROM users WHERE username = ?;"
	QueryRePwdStr  = "SELECT uid,re_password FROM users WHERE username = ?;"
	UpdatePwdStr   = "UPDATE users SET password = ? WHERE username = ?;"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"ezgin",
	)
	//连接数据库
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}

// CheckUser 检查用户名是否已存在
func CheckUser(username string) error {
	var exists bool
	err := db.QueryRow(IfExistUserStr, username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrorUserExist
	}
	return nil
}

// AddUser 添加用户
func AddUser(user *model.User) error {
	_, err := db.Exec(AddUserStr, user.Username, user.Password, user.RePassword)
	return err
}

// QueryPwd 查询密码
func QueryPwd(username string) (int, string, error) {
	var uid int
	var pwd string
	err := db.QueryRow(QueryPwdStr, username).Scan(&uid, &pwd)
	if err != nil {
		return -1, "", err
	}
	return uid, pwd, nil
}

func QueryRePwd(username string) (int, string, error) {
	var uid int
	var rePwd string
	err := db.QueryRow(QueryRePwdStr, username).Scan(&uid, &rePwd)
	if err != nil {
		return -1, "", err
	}
	return uid, rePwd, nil
}

// UpdatePwd 更新密码
func UpdatePwd(user *model.ParamResetPwdUser) error {
	_, err := db.Exec(UpdatePwdStr, user.Password, user.Username)
	return err
}
