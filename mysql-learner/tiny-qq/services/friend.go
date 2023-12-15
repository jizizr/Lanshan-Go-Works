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

func QueryFriendsList(uid int64) ([]model.UserFriend, error) {
	return mysql.QueryFriendsList(uid)
}

func SearchFriend(uid int64, username string) ([]*model.UserFriend, error) {
	return mysql.SearchFriend(uid, username)
}
