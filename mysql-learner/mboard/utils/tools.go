package utils

import (
	"ezgin/model"
	"github.com/gin-gonic/gin"
)

func GetUid(c *gin.Context) (UserID int, ok bool) {
	uid, ok := c.Get(model.CtxGetUID)
	if !ok {
		return
	}
	UserID, ok = uid.(int)
	if !ok {
		return
	}
	return
}
