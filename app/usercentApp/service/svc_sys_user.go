package service

import (
	"errors"
	"fmt"
	"new_it/app/usercentApp/model"
	"new_it/common"
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

func (userService *UserService) Register(u *model.SysUsers) (err error, userInter *model.SysUsers) {
	var user model.SysUsers
	if !errors.Is(global.GLB_DB.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}

	//u.UUID = uuid.NewV4()
	err = global.GLB_DB.Create(&u).Error
	return err, u
}

func (userService *UserService) ChangePassword(u *model.SysUsers, newPassword string) (err error, userInter *model.SysUsers) {
	var user model.SysUsers
	err = global.GLB_DB.Where("user_name = ? AND password = ?", u.UserName, u.Password).First(&user).Update("password", newPassword).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(user_id uint64) (err error) {

	var user model.SysUsers

	if errors.Is(global.GLB_DB.Where("user_id = ?", user_id).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户不存在")
	}

	pass := "123456" + "new_it"
	err = global.GLB_DB.Model(&model.SysUsers{}).Where("user_id = ?", user_id).Update("password", common.MD5([]byte(pass))).Error
	return err
}
