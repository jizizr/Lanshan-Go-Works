package mysql

import "errors"

var (
	ErrorUserExist    = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidPwd   = errors.New("密码错误")
	ErrorInvalidRePwd = errors.New("密保错误")
)
