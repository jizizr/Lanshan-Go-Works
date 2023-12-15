package mysql

import (
	"database/sql"
	"tiny-qq/model"
)

const (
	IfGroupExistStr       = "SELECT EXISTS(SELECT 1 FROM FriendGroups WHERE GroupID = ?);"
	IfGroupNameExistStr   = "SELECT EXISTS(SELECT 1 FROM FriendGroups WHERE UserID = ? AND GroupName = ?);"
	CreateGroupStr        = "INSERT INTO FriendGroups(GroupName,UserID) VALUES(?,?);"
	QueryGroupCreatorStr  = "SELECT UserID FROM FriendGroups WHERE GroupID = ?;"
	DeleteGroupStr        = "DELETE FROM FriendGroups WHERE GroupID = ?;"
	IfGroupUserExistStr   = "SELECT EXISTS(SELECT 1 FROM UserGroupRelations WHERE UserID = ? AND GroupID = ?);"
	AddGroupUserStr       = "INSERT INTO UserGroupRelations(UserID,GroupID) VALUES(?,?);"
	DeleteGroupUserStr    = "DELETE FROM UserGroupRelations WHERE UserID = ? AND GroupID = ?;"
	DeleteGroupAllUserStr = "DELETE FROM UserGroupRelations WHERE GroupID = ?;"
	QueryGroupUserStr     = `
		SELECT 
			Users.UserID,
			Users.Username
		FROM 
			Users
		JOIN 
			UserGroupRelations ON Users.UserID = UserGroupRelations.UserID
		JOIN 
			FriendGroups ON UserGroupRelations.GroupID = FriendGroups.GroupID
		WHERE 
			FriendGroups.GroupID = ?;
		`
)

// CheckGroup 检查群组是否已存在
// 如果存在返回ErrorGroupExist
// 如果不存在返回nil
func CheckGroup(gid int64) error {
	var exists bool
	err := db.QueryRow(IfGroupExistStr, gid).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrorGroupExist
	}
	return nil
}

// CheckGroupName 检查群组名称是否已存在
// 如果存在返回ErrorGroupNameExist
// 如果不存在返回nil
func CheckGroupName(param *model.ParamGroup) error {
	var exists bool
	err := db.QueryRow(IfGroupNameExistStr, param.UserID, param.GroupName).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrorGroupNameExist
	}
	return nil
}

// QueryGroupCreator 查询分组创建者
func QueryGroupCreator(gid int64) (int64, error) {
	var creator int64
	err := db.QueryRow(QueryGroupCreatorStr, gid).Scan(&creator)
	if err != nil {
		return 0, err
	}
	return creator, nil
}

// CreateGroup 创建群组
func CreateGroup(param *model.ParamGroup) (int64, error) {
	result, err := db.Exec(CreateGroupStr, param.GroupName, param.UserID)
	if err != nil {
		return 0, err
	}
	return GetIncrementID(result)
}

// DeleteGroup 删除群组
func DeleteGroup(param *model.ParamGroupID) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// 先删除分组中的所有用户关系
	if _, err := tx.Exec(DeleteGroupAllUserStr, param.GroupID); err != nil {
		tx.Rollback()
		return err
	}
	// 然后删除分组本身
	if _, err := tx.Exec(DeleteGroupStr, param.GroupID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// CheckGroupUser 检查群组用户是否已存在
// 如果存在返回ErrorGroupUserExist
// 如果不存在返回nil
func CheckGroupUser(param *model.ParamModifyGroupUser) error {
	var exists bool
	err := db.QueryRow(IfGroupUserExistStr, param.UserID, param.GroupID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrorGroupUserExist
	}
	return nil
}

// GetIncrementID 获取自增ID
func GetIncrementID(result sql.Result) (int64, error) {
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// AddGroupUser 添加群组用户
func AddGroupUser(param *model.ParamModifyGroupUser) (int64, error) {
	result, err := db.Exec(AddGroupUserStr, param.UserID, param.GroupID)
	if err != nil {
		return 0, err
	}
	return GetIncrementID(result)
}

// DeleteGroupUser 删除群组用户
func DeleteGroupUser(param *model.ParamModifyGroupUser) error {
	_, err := db.Exec(DeleteGroupUserStr, param.UserID, param.GroupID)
	return err
}

// QueryGroupUser 查询群组用户
func QueryGroupUser(param *model.ParamGroupID) ([]*model.UserFriend, error) {
	rows, err := db.Query(QueryGroupUserStr, param.GroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.UserFriend
	for rows.Next() {
		user := new(model.UserFriend)
		err := rows.Scan(&user.UserID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
