package service

import (
	"errors"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

type MenusService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.SysBaseMenu
//@return: error

func (m *MenusService) AddBaseMenu(menu model.SysBaseMenus) error {
	if !errors.Is(global.GLB_DB.Where("name = ?", menu.Name).First(&model.SysBaseMenus{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GLB_DB.Create(&menu).Error
}
