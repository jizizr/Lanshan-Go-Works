package services

import (
	"tiny-qq/dao/mysql"
	"tiny-qq/model"
)

func CreateGroup(param *model.ParamGroup) (int64, error) {
	if err := mysql.CheckGroupName(param); err != nil {
		return 0, err
	}
	return mysql.CreateGroup(param)
}

func QueryGroupCreator(gid int64) (int64, error) {
	if err := mysql.CheckGroup(gid); err == nil {
		return 0, mysql.ErrorGroupNotExist
	}
	return mysql.QueryGroupCreator(gid)
}

func DeleteGroup(param *model.ParamGroupID) error {
	if err := mysql.CheckGroup(param.GroupID); err == nil {
		return mysql.ErrorGroupNotExist
	}
	return mysql.DeleteGroup(param)
}

func AddGroupUser(param *model.ParamModifyGroupUser) (int64, error) {
	if err := mysql.CheckGroup(param.GroupID); err == nil {
		return 0, mysql.ErrorGroupNotExist
	} else if err := mysql.CheckGroupUser(param); err != nil {
		return 0, err
	}
	return mysql.AddGroupUser(param)
}

func DeleteGroupUser(param *model.ParamModifyGroupUser) error {
	if err := mysql.CheckGroup(param.GroupID); err == nil {
		return mysql.ErrorGroupNotExist
	} else if err := mysql.CheckGroupUser(param); err == nil {
		return mysql.ErrorGroupUserNotExist
	}
	return mysql.DeleteGroupUser(param)
}

func QueryGroupsList(param *model.ParamGroupID) ([]*model.UserFriend, error) {
	if err := mysql.CheckGroup(param.GroupID); err == nil {
		return nil, mysql.ErrorGroupNotExist
	}
	return mysql.QueryGroupUser(param)
}
