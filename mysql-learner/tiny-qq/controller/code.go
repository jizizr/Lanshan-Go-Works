package controller

type RespCode int16

const (
	CodeSuccess      RespCode = 0
	CodeInvalidParam RespCode = 1000 + iota
	CodeUserExist
	CodeUserNotExist
	CodeNeedLogin
	CodeInvalidToken
	CodeServiceBusy
	CodeWrongPassword
	CodeInvalidUser
	CodeFriendExist
	CodeFriendNotExist
	CodeGroupExist
	CodeGroupNotExist
	CodeGroupNameExist
	CodeGroupUserExist
	CodeGroupUserNotExist
)

var codeMsgMap = map[RespCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeNeedLogin:         "需要登录",
	CodeInvalidToken:      "无效的Token",
	CodeServiceBusy:       "服务繁忙",
	CodeWrongPassword:     "密码错误",
	CodeInvalidUser:       "用户ID不符",
	CodeFriendExist:       "好友已存在",
	CodeFriendNotExist:    "好友账号不存在",
	CodeGroupExist:        "分组已存在",
	CodeGroupNotExist:     "分组不存在",
	CodeGroupNameExist:    "分组名已存在",
	CodeGroupUserExist:    "分组中用户已存在",
	CodeGroupUserNotExist: "分组中用户不存在",
}

func (code RespCode) Msg() string {
	return codeMsgMap[code]
}

func (code RespCode) Error() string {
	return code.Msg()
}
