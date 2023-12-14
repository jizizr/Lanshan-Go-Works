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
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserToken struct {
	UID   int    `json:"uid"`
	Token string `json:"token"`
}
