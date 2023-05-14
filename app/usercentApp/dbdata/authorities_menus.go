package dbdata

import (
	"errors"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

func (a *authoritiesMenus) TableName() string {

	return "sys_authority_menus"
}

func (a *authoritiesMenus) Initialize() error {
	entities := []model.SysAuthorityMenus{
		//研发人员
		{MenuId: 1, AuthorityId: "666"},
		{MenuId: 2, AuthorityId: "666"},
		{MenuId: 3, AuthorityId: "666"},
		{MenuId: 4, AuthorityId: "666"},
		{MenuId: 5, AuthorityId: "666"},
		{MenuId: 6, AuthorityId: "666"},
		{MenuId: 7, AuthorityId: "666"},
		{MenuId: 8, AuthorityId: "666"},
		{MenuId: 9, AuthorityId: "666"},
		{MenuId: 10, AuthorityId: "666"},
		//管理员角色
		{MenuId: 2, AuthorityId: "777"},
		{MenuId: 3, AuthorityId: "777"},
		{MenuId: 4, AuthorityId: "777"},
		{MenuId: 5, AuthorityId: "777"},
		{MenuId: 6, AuthorityId: "777"},

		//普通用户
		{MenuId: 1, AuthorityId: "888"},
		{MenuId: 2, AuthorityId: "888"},
		{MenuId: 7, AuthorityId: "888"},
		{MenuId: 8, AuthorityId: "888"},
		{MenuId: 9, AuthorityId: "888"},
		{MenuId: 10, AuthorityId: "888"},

		//测试角色
		{MenuId: 1, AuthorityId: "999"},
		{MenuId: 2, AuthorityId: "999"},
		{MenuId: 3, AuthorityId: "999"},
		{MenuId: 4, AuthorityId: "999"},
		{MenuId: 5, AuthorityId: "999"},
		{MenuId: 6, AuthorityId: "999"},
		{MenuId: 7, AuthorityId: "999"},
		{MenuId: 8, AuthorityId: "999"},
		{MenuId: 9, AuthorityId: "999"},
		{MenuId: 10, AuthorityId: "999"},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(a.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (a *authoritiesMenus) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("menu_id = ? AND authority_id = ?", 1, "666").First(&model.SysAuthorityMenus{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
