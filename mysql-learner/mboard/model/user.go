package model

const CtxGetUID = "UID"

type ParamRegisterUser struct {
	Username   string `form:"username" json:"username"`
	Password   string `form:"password" json:"password"`
	RePassword string `form:"re_password" json:"re_password"`
}

type ParamLoginUser struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type ParamResetPwdUser struct {
	Username   string `form:"username" json:"username"`
	Password   string `form:"password" json:"password"`
	RePassword string `form:"re_password" json:"re_password"`
}

type User struct {
	Username string
	Password string
	//密保答案
	RePassword string
}

type UserToken struct {
	UID   int    `json:"uid"`
	Token string `json:"token"`
}
