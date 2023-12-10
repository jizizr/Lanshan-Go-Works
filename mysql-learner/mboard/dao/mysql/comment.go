package mysql

import (
	"database/sql"
	"ezgin/model"
	"time"
)

const (
	AddCommentStr      = "INSERT INTO comments (from_uid, to_uid,message) VALUES (?, ?, ?)"
	AddCommentReplyStr = "INSERT INTO comments (from_uid, to_uid,to_mid,message,reply_to_id) VALUES (?, ?, ?, ?, ?)"
	GetCommentStr      = "SELECT mid,from_uid,message,created_at FROM comments WHERE to_uid = ? AND reply_to_id IS NULL"
	GetCommentReplyStr = "SELECT mid,from_uid,to_uid,to_mid,message,created_at FROM comments WHERE reply_to_id = ?"
)

// PostComment 添加评论
func PostComment(fromUID, toUID int, message string) (int64, error) {
	result, err := db.Exec(AddCommentStr, fromUID, toUID, message)
	if err != nil {
		return -1, err
	}
	mid, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return mid, nil
}

// PostCommentReply 添加评论回复
func PostCommentReply(fromUID, toUID int, toMID int64, message string, replyID int) (int64, error) {
	result, err := db.Exec(AddCommentReplyStr, fromUID, toUID, toMID, message, replyID)
	if err != nil {
		return -1, err
	}
	mid, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return mid, nil
}

// GetComment 获取评论
func GetComment(toUID int) ([]*model.CommentInfo, error) {
	commentStmt, err := db.Prepare(GetCommentStr)
	if err != nil {
		return nil, err
	}
	defer commentStmt.Close()
	replyStmt, err := db.Prepare(GetCommentReplyStr)
	if err != nil {
		return nil, err
	}
	defer replyStmt.Close()
	commentRows, err := commentStmt.Query(toUID)
	if err != nil {
		return nil, err
	}
	defer commentRows.Close()
	var commentInfos []*model.CommentInfo
	for commentRows.Next() {
		var commentInfo model.CommentInfo
		//mid,from_uid,message,created_at
		var t time.Time
		err := commentRows.Scan(&commentInfo.MID, &commentInfo.FromUID, &commentInfo.Message, &t)
		commentInfo.Time = t.Unix()
		if err != nil {
			return nil, err
		}
		replys, err := GetCommentReply(replyStmt, commentInfo.MID)
		if err != nil {
			return nil, err
		}
		commentInfo.Reply = replys
		commentInfos = append(commentInfos, &commentInfo)
	}
	return commentInfos, nil
}

// GetCommentReply 获取评论回复
func GetCommentReply(stmt *sql.Stmt, replyID int64) ([]*model.Reply, error) {
	var replys []*model.Reply
	replyRows, err := stmt.Query(replyID)
	if err != nil {
		return nil, err
	}
	for replyRows.Next() {
		var reply model.Reply
		var t time.Time
		//mid,from_uid,to_uid,to_mid,message,created_at
		err := replyRows.Scan(&reply.MID, &reply.FromUID, &reply.ToUID, &reply.ToMID, &reply.Message, &t)
		if err != nil {
			return nil, err
		}
		reply.Time = t.Unix()
		replys = append(replys, &reply)
	}
	return replys, nil
}
