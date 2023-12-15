package model

const CtxGetUID = "UID"

type ParamRegisterUser struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type ParamLoginUser struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type User struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserToken struct {
	UID   int64  `json:"uid"`
	Token string `json:"token"`
}
