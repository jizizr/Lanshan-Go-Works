package mysql

import (
	"github.com/go-sql-driver/mysql"
	"tiny-qq/model"
)

const (
	IfExistFriendStr    = "SELECT EXISTS(SELECT 1 FROM Friends WHERE UserID = ? AND FriendID = ?);"
	AddFriendStr        = "INSERT INTO Friends(UserID,FriendID) VALUES(?,?);"
	DeleteFriendStr     = "DELETE FROM Friends WHERE UserID = ? AND FriendID = ?;"
	QueryFriendsListStr = `
        SELECT u.UserID, u.Username
        FROM Users u
        JOIN (SELECT FriendID FROM Friends WHERE UserID = ?) AS f ON u.UserID = f.FriendID
    `
	SearchFriendStr = `
		SELECT u.UserID, u.Username
		FROM Users u 
		JOIN Friends ON u.UserID = Friends.FriendID 
		WHERE Friends.UserID = ? AND u.Username LIKE CONCAT('%', ?, '%');
`
)

// CheckFriend 检查好友是否已存在
func CheckFriend(param *model.ParamModifyFriend) error {
	var exists bool
	err := db.QueryRow(IfExistFriendStr, param.UserID, param.FriendID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrorFriendExist
	}
	return nil
}

// AddFriend 添加好友
func AddFriend(param *model.ParamModifyFriend) error {
	_, err := db.Exec(AddFriendStr, param.UserID, param.FriendID)
	if err != nil && err.(*mysql.MySQLError).Number == 1452 {
		return ErrorFriendNotExist
	}
	return err
}

// DeleteFriend 删除好友
func DeleteFriend(param *model.ParamModifyFriend) error {
	_, err := db.Exec(DeleteFriendStr, param.UserID, param.FriendID)
	return err
}

// QueryFriendsList 查询好友列表
func QueryFriendsList(uid int64) ([]model.UserFriend, error) {
	rows, err := db.Query(QueryFriendsListStr, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.UserFriend
	for rows.Next() {
		var user model.UserFriend
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// SearchFriend 搜索好友
func SearchFriend(uid int64, username string) ([]*model.UserFriend, error) {
	rows, err := db.Query(SearchFriendStr, uid, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.UserFriend
	for rows.Next() {
		user := new(model.UserFriend)
		if err := rows.Scan(&user.UserID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
