package utils

import (
	"ezgin/model"
	"github.com/gin-gonic/gin"
)

func GetUid(c *gin.Context) (UserID int64, ok bool) {
	uid, ok := c.Get(model.CtxGetUID)
	if !ok {
		return
	}
	UserID, ok = uid.(int64)
	if !ok {
		return
	}
	return
}
