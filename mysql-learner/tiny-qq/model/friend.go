package model

type ParamModifyFriend struct {
	UserID   int `form:"user_id" json:"user_id"`
	FriendID int `form:"friend_id" json:"friend_id"`
}

type UserFriend struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}
