package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"tiny-qq/dao/mysql"
	"tiny-qq/model"
	"tiny-qq/services"
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
	if ParamUser.Username == "" || ParamUser.Password == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//根据错误类型返回响应
	if err := services.Register(ParamUser); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			RespFailed(c, CodeUserExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		log.Printf("%v", err)
		return
	}
	RespSuccess(c, nil)
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
		uid   int
	)
	uid, token, err = services.Login(ParamUser)
	//判断错误类型
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			RespFailed(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorInvalidPwd) {
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
