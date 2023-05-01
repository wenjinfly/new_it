package dbdata

import (
	"errors"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

var Authority = new(USERAuthority)

type USERAuthority struct{}

func (a *USERAuthority) TableName() string {
	return "sys_authorities"
}

func (a *USERAuthority) Initialize() error {
	entities := []model.SysAuthorities{
		{AuthorityId: "666", AuthorityName: "研发人员", ParentId: "0", DefaultRouter: "MainIndex"},
		{AuthorityId: "777", AuthorityName: "管理员角色", ParentId: "0", DefaultRouter: "admin"},
		{AuthorityId: "888", AuthorityName: "普通用户", ParentId: "0", DefaultRouter: "MainIndex"},
		{AuthorityId: "999", AuthorityName: "测试角色", ParentId: "0", DefaultRouter: "MainIndex"},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(a.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (a *USERAuthority) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("authority_id = ? ", "666").First(&model.SysAuthorities{}).Error, gorm.ErrRecordNotFound) {
		// 判断是否存在数据
		return false
	}
	return true
}
