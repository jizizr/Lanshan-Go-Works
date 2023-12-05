package services

import (
	"ezgin/dao/database"
	"ezgin/model"
)

func PostComment(fromUID int64, comment *model.ParamComment) error {
	return database.PostComment(fromUID, comment)
}

func GetComment(uid int64) ([]*model.CommentDB, error) {
	return database.GetComment(uid)
}
