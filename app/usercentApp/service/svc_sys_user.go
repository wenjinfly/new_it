package service

import (
	"errors"
	"fmt"
	"new_it/app/usercentApp/model"
	"new_it/common"
	"new_it/global"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Login(u *model.SysUsers) (err error, userInter *model.SysUsers) {
	if nil == global.GLB_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user model.SysUsers
	err = global.GLB_DB.Where("user_name = ? AND password = ?", u.UserName, u.Password).Preload("Authorities").First(&user).Error
	if err != nil {
		return err, &user
	}

	var sa model.SysAuthorities
	err = global.GLB_DB.Where("authority_id = ?", user.AuthorityId).First(&sa).Error

	user.Authority = sa

	return err, &user
}

func (userService *UserService) Register(u *model.SysUsers) (err error, userInter *model.SysUsers) {
	var user model.SysUsers
	if !errors.Is(global.GLB_DB.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}

	uid, _ := uuid.NewV4()
	u.Uuid = uid.String()
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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid string) (err error, user model.SysUsers) {
	var reqUser model.SysUsers
	err = global.GLB_DB.First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info common.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.SysUsers{})
	var userList []model.SysUsers
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Find(&userList).Error
	//err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(reqUser model.SysUsers) (err error, user model.SysUsers) {
	err = global.GLB_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(userid uint64, uuid string, authorityId string) (err error) {
	assignErr := global.GLB_DB.Where("user_id = ? AND authority_id = ?", userid, authorityId).First(&model.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GLB_DB.Where("uuid = ?", uuid).First(&model.SysUsers{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint64, authorityIds []string) (err error) {
	return global.GLB_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]model.SysUserAuthority{}, "user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}

		useAuthority := []model.SysUserAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, model.SysUserAuthority{
				UserId:      id,
				AuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("user_id = ?", id).First(&model.SysUsers{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}
