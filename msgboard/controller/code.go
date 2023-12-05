package controller

type RespCode int16

const (
	CodeSuccess      RespCode = 0
	CodeInvalidParam RespCode = 1000 + iota
	CodeUserExist
	CodeUserNotExist
	CodeNeedLogin
	CodeInvalidToken
	CodeTooLongRePassword
	CodeServiceBusy
	CodeWrongPassword
	CodeWrongRePassword
	CodeInvalidUser
	CodeCommentNotExist
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeNeedLogin:         "需要登录",
	CodeInvalidToken:      "无效的Token",
	CodeTooLongRePassword: "密保答案过长",
	CodeServiceBusy:       "服务繁忙",
	CodeWrongPassword:     "密码错误",
	CodeWrongRePassword:   "密保不一致",
	CodeInvalidUser:       "用户ID不符",
	CodeCommentNotExist:   "评论不存在",
}

func (code RespCode) Msg() string {
	return codeMsgMap[code]
}

func (code RespCode) Error() string {
	return code.Msg()
}
