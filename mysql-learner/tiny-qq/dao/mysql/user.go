package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"tiny-qq/model"
)

const (
	IfExistUserStr = "SELECT EXISTS(SELECT 1 FROM Users WHERE username = ?);"
	AddUserStr     = "INSERT INTO Users(Username,Password) VALUES(?,?);"
	QueryPwdStr    = "SELECT UserID,Password FROM Users WHERE username = ?;"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"tqq",
	)
	//连接数据库
	var err error
	db, err = sql.Open("mysql", dsn)
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
	_, err := db.Exec(AddUserStr, user.Username, user.Password)
	return err
}

// QueryPwd 查询密码
func QueryPwd(username string) (int64, string, error) {
	var uid int64
	var pwd string
	err := db.QueryRow(QueryPwdStr, username).Scan(&uid, &pwd)
	if err != nil {
		return -1, "", err
	}
	return uid, pwd, nil
}
