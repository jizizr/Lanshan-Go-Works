package controller

import (
	"errors"
	"ezgin/dao/database"
	"ezgin/model"
	"ezgin/services"
	"ezgin/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// Register 注册
func Register(c *gin.Context) {
	//获取参数并校验
	ParamUser := new(model.ParamRegisterUser)
	if err := c.ShouldBind(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		log.Println(err)
		return
	}
	if ParamUser.Username == "" || ParamUser.Password == "" || ParamUser.RePassword == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if len(ParamUser.RePassword) > 255 {
		RespFailed(c, CodeTooLongRePassword)
		return
	}
	//根据错误类型返回响应
	if err := services.Register(ParamUser); err != nil {
		if errors.Is(err, database.ErrorUserExist) {
			RespFailed(c, CodeUserExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		log.Printf("%v", err)
		return
	}
	RespSuccess(c, nil)
}

// ResetPwd 重置密码
func ResetPwd(c *gin.Context) {
	//获取参数并校验
	ParamUser := new(model.ParamResetPwdUser)
	if err := c.ShouldBind(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.Username == "" || ParamUser.Password == "" || ParamUser.RePassword == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if len(ParamUser.RePassword) > 255 {
		RespFailed(c, CodeTooLongRePassword)
		return
	}
	uid, err := services.ResetPwd(ParamUser)
	if err != nil {
		if errors.Is(err, database.ErrorUserNotExist) {
			RespFailed(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, database.ErrorInvalidRePwd) {
			RespFailed(c, CodeWrongRePassword)
			return
		}
		RespFailed(c, CodeServiceBusy)
		log.Printf("%v", err)
		return
	}
	token, _ := utils.GenToken(uid)
	RespSuccess(c, &model.UserToken{
		UID:   uid,
		Token: token,
	})
}

// Login 登录
func Login(c *gin.Context) {
	ParamUser := new(model.ParamLoginUser)
	if err := c.ShouldBind(ParamUser); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamUser.Username == "" || ParamUser.Password == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	var (
		token string
		err   error
		uid   int64
	)
	uid, token, err = services.Login(ParamUser)
	//判断错误类型
	if err != nil {
		if errors.Is(err, database.ErrorUserNotExist) {
			RespFailed(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, database.ErrorInvalidPwd) {
			RespFailed(c, CodeWrongPassword)
			return
		}
		log.Println(err)
		RespFailed(c, CodeServiceBusy)
		return
	}
	//返回token
	RespSuccess(c, &model.UserToken{
		UID:   uid,
		Token: token,
	})
}
