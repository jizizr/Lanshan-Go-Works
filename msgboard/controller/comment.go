package controller

import (
	"errors"
	"ezgin/dao/database"
	"ezgin/model"
	"ezgin/services"
	"ezgin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func PostComment(c *gin.Context) {
	uid, ok := utils.GetUid(c)
	if !ok {
		RespFailed(c, CodeNeedLogin)
		return
	}
	//获取参数并校验
	ParamComment := new(model.ParamComment)
	if err := c.ShouldBind(ParamComment); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if ParamComment.ToUID == 0 || ParamComment.Message == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	//根据错误类型返回响应
	err := services.PostComment(uid, ParamComment)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		log.Printf("%v", err)
		return
	}
	RespSuccess(c, nil)
}

func GetComment(c *gin.Context) {
	id := c.Param("uid")
	//string to int
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	commentInfo, err := services.GetComment(uid)
	if err != nil {
		if errors.Is(err, database.ErrorCommentNotExist) {
			RespFailed(c, CodeCommentNotExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		log.Printf("%v", err)
		return
	}
	RespSuccess(c, commentInfo)
}
