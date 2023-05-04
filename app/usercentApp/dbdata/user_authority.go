package dbdata

import (
	"errors"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

var UserAuthority = new(userAuthority)

type userAuthority struct{}

func (a *userAuthority) TableName() string {

	return "sys_user_authority"
}

func (a *userAuthority) Initialize() error {
	entities := []model.SysUserAuthority{
		{UserId: 1, AuthorityId: "666"},
		{UserId: 2, AuthorityId: "777"},
		{UserId: 3, AuthorityId: "666"},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(a.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (a *userAuthority) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("user_id = ? AND authority_id = ?", 2, "777").First(&model.SysUserAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
