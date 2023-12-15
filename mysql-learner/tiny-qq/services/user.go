package services

import (
	"errors"
	"tiny-qq/dao/mysql"
	"tiny-qq/model"
	"tiny-qq/utils"
)

func Register(ParamUser *model.ParamRegisterUser) error {
	//检查用户名是否已存在
	if err := mysql.CheckUser(ParamUser.Username); err != nil {
		return err
	}
	user := &model.User{
		Username: ParamUser.Username,
		Password: utils.Md5(ParamUser.Password),
	}
	return mysql.AddUser(user)
}

func Login(ParamUser *model.ParamLoginUser) (int64, string, error) {
	if err := mysql.CheckUser(ParamUser.Username); !errors.Is(err, mysql.ErrorUserExist) {
		if err != nil {
			return -1, "", err
		}
		return -1, "", mysql.ErrorUserNotExist
	}
	uid, pwd, err := mysql.QueryPwd(ParamUser.Username)
	if err != nil {
		return -1, "", err
	}
	if utils.Md5(ParamUser.Password) != pwd {
		return -1, "", mysql.ErrorInvalidPwd
	}
	token, _ := utils.GenToken(uid)
	return uid, token, nil
}
