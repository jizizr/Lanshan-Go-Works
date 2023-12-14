package mysql

import "errors"

var (
	ErrorUserExist         = errors.New("用户已存在")
	ErrorUserNotExist      = errors.New("用户不存在")
	ErrorInvalidPwd        = errors.New("密码错误")
	ErrorFriendExist       = errors.New("好友已存在")
	ErrorFriendNotExist    = errors.New("好友账号不存在")
	ErrorGroupExist        = errors.New("分组已存在")
	ErrorGroupNotExist     = errors.New("分组不存在")
	ErrorGroupNameExist    = errors.New("分组名称已存在")
	ErrorGroupUserExist    = errors.New("分组中用户已存在")
	ErrorGroupUserNotExist = errors.New("分组用户不存在")
)
