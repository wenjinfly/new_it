package service

import (
	"fmt"
	"new_it/app/usercentApp/model"
	"new_it/global"
)

type UserService struct{}

func (userService *UserService) Login(u *model.SysUsers) (err error, userInter *model.SysUsers) {
	if nil == global.GLB_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user model.SysUsers
	//err = global.GLB_DB.Where("user_name = ? AND password = ?", u.UserName, u.Password).Preload("authority_id").First(&user).Error
	err = global.GLB_DB.Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user).Error
	return err, &user
}
