package database

import (
	"ezgin/model"
)

var CommentDb *CommentDB

func init() {
	CommentDb = NewCommentDB()
	CommentDb.LoadFromFile("COMMENT")
}

func PostComment(fromUID int64, c *model.ParamComment) error {
	CDB := &model.CommentDB{
		ToUID:   c.ToUID,
		FromUID: fromUID,
		Message: c.Message,
	}
	return CommentDb.AddComment(CDB)
}

// GetComment 获取评论
func GetComment(toUID int64) ([]*model.CommentDB, error) {
	return CommentDb.QueryComment(toUID)
}
