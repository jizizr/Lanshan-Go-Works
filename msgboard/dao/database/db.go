package database

import (
	"encoding/gob"
	"ezgin/model"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type UserDB struct {
	data         sync.Map
	autoIncrease atomic.Int64
}

func NewUserDB() *UserDB {
	return &UserDB{
		data:         sync.Map{},
		autoIncrease: atomic.Int64{},
	}
}

func (db *UserDB) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	db.data.Range(func(key, value interface{}) bool {
		err = encoder.Encode(key)
		if err != nil {
			return false
		}
		err = encoder.Encode(value)
		return err == nil
	})

	return err
}

func (db *UserDB) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var key string
	var value model.User
	for {
		err = decoder.Decode(&key)
		if err != nil {
			break
		}
		err = decoder.Decode(&value)
		if err != nil {
			break
		}
		db.data.Store(key, &value)
	}

	if err != io.EOF {
		return err
	}

	return nil
}

func (db *UserDB) AddUser(user *model.User) error {
	uid := db.autoIncrease.Add(1)
	user.UID = uid
	db.data.Store(user.Username, user)
	return nil
}

func (db *UserDB) QueryPwd(username string) (int64, string, error) {
	value, ok := db.data.Load(username)
	if !ok {
		return -1, "", ErrorUserNotExist
	}
	user := value.(*model.User)
	return user.UID, user.Password, nil
}

func (db *UserDB) QueryRePwd(username string) (int64, string, error) {
	value, ok := db.data.Load(username)
	if !ok {
		return -1, "", ErrorUserNotExist
	}
	user := value.(*model.User)
	return user.UID, user.RePassword, nil
}

func (db *UserDB) UpdatePwd(user *model.ParamResetPwdUser) error {
	value, ok := db.data.Load(user.Username)
	if !ok {
		return ErrorUserNotExist
	}
	oldUser := value.(*model.User)
	oldUser.Password = user.Password
	db.data.Store(user.Username, oldUser)
	return nil
}

func (db *UserDB) CheckUser(username string) error {
	_, ok := db.data.Load(username)
	if ok {
		return ErrorUserExist
	}
	return nil
}

type CommentDB struct {
	data sync.Map
}

func NewCommentDB() *CommentDB {
	return &CommentDB{
		data: sync.Map{},
	}
}

func (db *CommentDB) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	db.data.Range(func(key, value interface{}) bool {
		err = encoder.Encode(key)
		if err != nil {
			return false
		}
		err = encoder.Encode(value)
		return err == nil
	})

	return err
}

func (db *CommentDB) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var key int64
	var value []*model.CommentDB
	for {
		err = decoder.Decode(&key)
		if err != nil {
			break
		}
		err = decoder.Decode(&value)
		if err != nil {
			break
		}
		db.data.Store(key, value)
	}

	if err != io.EOF {
		return err
	}

	return nil
}

func (db *CommentDB) AddComment(comment *model.CommentDB) error {
	comment.Time = time.Now().Unix()

	db.data.LoadOrStore(comment.ToUID, []*model.CommentDB{comment})

	actual, _ := db.data.Load(comment.ToUID)
	comments := actual.([]*model.CommentDB)
	comments = append(comments, comment)

	db.data.Store(comment.ToUID, comments)

	return nil
}

func (db *CommentDB) QueryComment(toUID int64) ([]*model.CommentDB, error) {
	value, ok := db.data.Load(toUID)
	if !ok {
		return nil, ErrorCommentNotExist
	}
	return value.([]*model.CommentDB), nil
}
