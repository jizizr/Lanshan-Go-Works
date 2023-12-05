package services

import (
	"errors"
	"ezgin/dao/database"
	"ezgin/model"
	"ezgin/utils"
)

func Register(ParamUser *model.ParamRegisterUser) error {
	//检查用户名是否已存在
	if err := database.CheckUser(ParamUser.Username); err != nil {
		return err
	}
	user := &model.User{
		Username:   ParamUser.Username,
		Password:   utils.Md5(ParamUser.Password),
		RePassword: ParamUser.RePassword,
	}
	return database.AddUser(user)
}

func ResetPwd(ParamUser *model.ParamResetPwdUser) (int64, error) {
	//检查用户名是否已存在
	if err := database.CheckUser(ParamUser.Username); !errors.Is(err, database.ErrorUserExist) {
		if err != nil {
			return -1, err
		}
		return -1, database.ErrorUserNotExist
	}
	uid, rePwd, err := database.QueryRePwd(ParamUser.Username)
	if err != nil {
		return -1, err
	}
	if ParamUser.RePassword != rePwd {
		return -1, database.ErrorInvalidRePwd
	}
	ParamUser.Password = utils.Md5(ParamUser.Password)
	return uid, database.UpdatePwd(ParamUser)
}

func Login(ParamUser *model.ParamLoginUser) (int64, string, error) {
	if err := database.CheckUser(ParamUser.Username); !errors.Is(err, database.ErrorUserExist) {
		if err != nil {
			return -1, "", err
		}
		return -1, "", database.ErrorUserNotExist
	}
	uid, pwd, err := database.QueryPwd(ParamUser.Username)
	if err != nil {
		return -1, "", err
	}
	if utils.Md5(ParamUser.Password) != pwd {
		return -1, "", database.ErrorInvalidPwd
	}
	token, _ := utils.GenToken(uid)
	return uid, token, nil
}
