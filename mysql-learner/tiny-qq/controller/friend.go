package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"tiny-qq/dao/mysql"
	"tiny-qq/model"
	"tiny-qq/services"
	"tiny-qq/utils"
)

// AddFriend 添加好友
func AddFriend(c *gin.Context) {
	param := new(model.ParamModifyFriend)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.UserID == 0 || param.FriendID == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, _ := utils.GetUid(c)
	if uid != param.UserID {
		RespFailed(c, CodeInvalidUser)
	}
	if err := services.AddFriend(param); err != nil {
		if errors.Is(err, mysql.ErrorFriendNotExist) {
			RespFailed(c, CodeFriendNotExist)
			return
		} else if errors.Is(err, mysql.ErrorFriendExist) {
			RespFailed(c, CodeFriendExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		log.Println(err)
		return
	}
	RespSuccess(c, nil)
}

// DeleteFriend 删除好友
func DeleteFriend(c *gin.Context) {
	param := new(model.ParamModifyFriend)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.UserID == 0 || param.FriendID == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	uid, _ := utils.GetUid(c)
	if uid != param.UserID {
		RespFailed(c, CodeInvalidUser)
	}
	if err := services.DeleteFriend(param); err != nil {
		if errors.Is(err, mysql.ErrorFriendNotExist) {
			RespFailed(c, CodeFriendNotExist)
			return
		}
		RespFailed(c, CodeServiceBusy)
		log.Println(err)
		return
	}
	RespSuccess(c, nil)
}

// QueryFriendsList 查询好友列表
func QueryFriendsList(c *gin.Context) {
	uid, _ := utils.GetUid(c)
	list, err := services.QueryFriendsList(uid)
	if err != nil {
		RespFailed(c, CodeServiceBusy)
		log.Println(err)
		return
	}
	RespSuccess(c, list)
}
