package model

type ParamComment struct {
	ToUID   int64  `form:"to_uid" json:"to_uid"`
	Message string `form:"message" json:"message"`
}

type Comment struct {
	ID int64 `json:"mid"`
}

type CommentDB struct {
	ToUID   int64  `json:"to_uid"`
	FromUID int64  `json:"from_uid"`
	Message string `json:"message"`
	Time    int64  `json:"created_at"`
}
