package services

import (
	"tiny-qq/dao/mysql"
	"tiny-qq/model"
)

func AddFriend(param *model.ParamModifyFriend) error {
	if err := mysql.CheckFriend(param); err != nil {
		return err
	}
	return mysql.AddFriend(param)
}

func DeleteFriend(param *model.ParamModifyFriend) error {
	return mysql.DeleteFriend(param)
}

// QueryFriendsList TODO: 验证用户查询权限
func QueryFriendsList(uid int) ([]model.UserFriend, error) {
	return mysql.QueryFriendsList(uid)
}
