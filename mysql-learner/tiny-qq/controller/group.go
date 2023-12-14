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

// CreateGroup 创建分组
func CreateGroup(c *gin.Context) {
	param := new(model.ParamGroup)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.UserID == 0 || param.GroupName == "" {
		RespFailed(c, CodeInvalidParam)
		return
	}
	var (
		err error
		id  int64
	)
	uid, _ := utils.GetUid(c)
	if uid != param.UserID {
		RespFailed(c, CodeInvalidUser)
	}
	if id, err = services.CreateGroup(param); err != nil {
		if errors.Is(err, mysql.ErrorGroupNameExist) {
			RespFailed(c, CodeGroupNameExist)
			return
		} else if errors.Is(err, mysql.ErrorGroupExist) {
			RespFailed(c, CodeGroupExist)
		} else {
			RespFailed(c, CodeServiceBusy)
			log.Println(err)
			return
		}
	}
	data := model.IncrementID{
		ID: id,
	}
	RespSuccess(c, data)
}

// DeleteGroup 删除分组
func DeleteGroup(c *gin.Context) {
	param := new(model.ParamGroupID)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.GroupID == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if err := services.DeleteGroup(param); err != nil {
		if errors.Is(err, mysql.ErrorGroupNotExist) {
			RespFailed(c, CodeGroupNotExist)
			return
		} else {
			RespFailed(c, CodeServiceBusy)
			log.Println(err)
			return
		}
	}
	RespSuccess(c, nil)
}

// AddGroupUser 向分组添加用户
func AddGroupUser(c *gin.Context) {
	param := new(model.ParamModifyGroupUser)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.GroupID == 0 || param.UserID == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	var (
		err error
		id  int64
	)
	if id, err = services.AddGroupUser(param); err != nil {
		if errors.Is(err, mysql.ErrorGroupNotExist) {
			RespFailed(c, CodeGroupNotExist)
			return
		} else if errors.Is(err, mysql.ErrorGroupUserExist) {
			RespFailed(c, CodeGroupUserExist)
			return
		} else {
			RespFailed(c, CodeServiceBusy)
			log.Println(err)
			return
		}
	}
	data := model.IncrementID{
		ID: id,
	}
	RespSuccess(c, data)
}

// DeleteGroupUser 从分组删除用户
func DeleteGroupUser(c *gin.Context) {
	param := new(model.ParamModifyGroupUser)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.GroupID == 0 || param.UserID == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}

	if err := services.DeleteGroupUser(param); err != nil {
		if errors.Is(err, mysql.ErrorGroupNotExist) {
			RespFailed(c, CodeGroupNotExist)
			return
		} else if errors.Is(err, mysql.ErrorGroupUserNotExist) {
			RespFailed(c, CodeGroupUserNotExist)
			return
		} else {
			RespFailed(c, CodeServiceBusy)
			log.Println(err)
			return
		}
	}
	RespSuccess(c, nil)
}

// QueryGroupsList 查询分组用户列表
func QueryGroupsList(c *gin.Context) {
	param := new(model.ParamGroupID)
	if err := c.ShouldBind(param); err != nil {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if param.GroupID == 0 {
		RespFailed(c, CodeInvalidParam)
		return
	}
	if userList, err := services.QueryGroupsList(param); err != nil {
		if errors.Is(err, mysql.ErrorGroupNotExist) {
			RespFailed(c, CodeGroupNotExist)
			return
		} else {
			RespFailed(c, CodeServiceBusy)
			log.Println(err)
			return
		}
	} else {
		RespSuccess(c, userList)
		return
	}
}
