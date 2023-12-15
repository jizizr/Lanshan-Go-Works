package model

type ParamGroup struct {
	UserID    int64  `form:"user_id" json:"user_id"`
	GroupName string `form:"name" json:"name"`
}

type ParamGroupID struct {
	GroupID int64 `form:"group_id" json:"group_id"`
}

type ParamModifyGroupUser struct {
	GroupID int64 `form:"group_id" json:"group_id"`
	UserID  int64 `form:"user_id" json:"user_id"`
}

type IncrementID struct {
	ID int64 `json:"id"`
}
