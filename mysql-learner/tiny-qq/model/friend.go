package model

type ParamModifyFriend struct {
	UserID   int64 `form:"user_id" json:"user_id"`
	FriendID int64 `form:"friend_id" json:"friend_id"`
}

type UserFriend struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}
