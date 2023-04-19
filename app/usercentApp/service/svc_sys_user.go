package service

import (
	"errors"
	"fmt"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
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

func (userService *UserService) Register(u model.SysUsers) (err error, userInter model.SysUsers) {
	var user model.SysUsers
	if !errors.Is(global.GLB_DB.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}

	//u.UUID = uuid.NewV4()
	err = global.GLB_DB.Create(&u).Error
	return err, u
}
