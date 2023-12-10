package model

type ParamComment struct {
	ReplyID int    `form:"reply_to_mid" json:"reply_to_mid"`
	FromUID int    `form:"from_uid" json:"from_uid"`
	ToUID   int    `form:"to_uid" json:"to_uid"`
	ToMID   int64  `form:"to_mid" json:"to_mid"`
	Message string `form:"message" json:"message"`
}

type Comment struct {
	ID int64 `json:"mid"`
}

type CommentInfo struct {
	MID     int64    `json:"mid"`
	FromUID int      `json:"from_uid"`
	Time    int64    `json:"time"`
	Message string   `json:"message"`
	Reply   []*Reply `json:"reply"`
}

type Reply struct {
	MID     int64  `json:"mid"`
	FromUID int    `json:"from_uid"`
	ToUID   int    `json:"to_uid"`
	ToMID   int64  `json:"to_mid"`
	Time    int64  `json:"time"`
	Message string `json:"message"`
}
