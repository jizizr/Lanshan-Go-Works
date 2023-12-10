package services

import (
	"ezgin/dao/mysql"
	"ezgin/model"
)

func PostComment(comment *model.ParamComment) (int64, error) {
	if comment.ReplyID != 0 {
		return mysql.PostCommentReply(comment.FromUID, comment.ToUID, comment.ToMID, comment.Message, comment.ReplyID)
	}
	return mysql.PostComment(comment.FromUID, comment.ToUID, comment.Message)
}

func GetComment(uid int) ([]*model.CommentInfo, error) {
	return mysql.GetComment(uid)
}
