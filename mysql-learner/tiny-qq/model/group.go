package model

type ParamGroup struct {
	UserID    int    `form:"user_id" json:"user_id"`
	GroupName string `form:"name" json:"name"`
}

type ParamGroupID struct {
	GroupID int `form:"group_id" json:"group_id"`
}

type ParamModifyGroupUser struct {
	GroupID int `form:"group_id" json:"group_id"`
	UserID  int `form:"user_id" json:"user_id"`
}

type IncrementID struct {
	ID int64 `json:"id"`
}
