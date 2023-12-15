package utils

import (
	"github.com/gin-gonic/gin"
	"tiny-qq/model"
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
