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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseMenuTreeMap
//@description: 获取路由总树map
//@return: err error, treeMap map[int][]model.SysBaseMenu

func (m *MenusService) getBaseMenuTreeMap() (err error, treeMap map[int][]model.SysBaseMenus) {
	var allMenus []model.SysBaseMenus
	treeMap = make(map[int][]model.SysBaseMenus)
	err = global.GLB_DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getBaseChildrenList
//@description: 获取菜单的子菜单
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func (m *MenusService) getBaseChildrenList(menu *model.SysBaseMenus, treeMap map[int][]model.SysBaseMenus) (err error) {
	menu.Children = treeMap[int(menu.MenusId)]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetInfoList
//@description: 获取路由分页
//@return: err error, list interface{}, total int64

func (m *MenusService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []model.SysBaseMenus
	err, treeMap := m.getBaseMenuTreeMap()
	menuList = treeMap[0]

	total = int64(len(menuList))

	for i := 0; i < len(menuList); i++ {
		err = m.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}
