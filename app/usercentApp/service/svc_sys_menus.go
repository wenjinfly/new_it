package service

import (
	"errors"
	"new_it/app/usercentApp/api/request"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

type MenusService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

func (m *MenusService) getMenuTreeMap(authorityId string) (err error, treeMap map[int][]model.SysBaseMenus) {
	var allMenus []model.SysBaseMenus
	treeMap = make(map[int][]model.SysBaseMenus)
	err = global.GLB_DB.Where("authority_id = ?", authorityId).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddBaseMenu
//@description: 添加基础路由
//@param: menu model.SysBaseMenu
//@return: error

func (m *MenusService) AddBaseMenu(menu model.SysBaseMenus) error {
	if !errors.Is(global.GLB_DB.Where("name = ?", menu.Name).First(&model.SysBaseMenus{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name,请修改name")
	}
	return global.GLB_DB.Create(&menu).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuTree
//@description: 获取动态菜单树
//@param: authorityId string
//@return: err error, menus []model.SysMenu

func (m *MenusService) GetMenuTree(authorityId string) (err error, menus []model.SysBaseMenus) {
	err, menuTree := m.getMenuTreeMap(authorityId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = m.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: getChildrenList
//@description: 获取子菜单
//@param: menu *model.SysMenu, treeMap map[string][]model.SysMenu
//@return: err error

func (m *MenusService) getChildrenList(menu *model.SysBaseMenus, treeMap map[int][]model.SysBaseMenus) (err error) {
	menu.Children = treeMap[int(menu.MenuId)]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
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
	menu.Children = treeMap[int(menu.MenuId)]
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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: AddMenuAuthority
//@description: 为角色增加menu树
//@param: menus []model.SysBaseMenu, authorityId string
//@return: err error

func (m *MenusService) AddMenuAuthority(menus []model.SysBaseMenus, authorityId string) (err error) {
	var auth model.SysAuthorities
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServices.SetMenuAuthority(&auth)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBaseMenu
//@description: 删除基础路由
//@param: id float64
//@return: err error

func (m *MenusService) DeleteBaseMenu(menu_id uint64) (err error) {
	err = global.GLB_DB.Where("parent_id = ?", menu_id).First(&model.SysBaseMenus{}).Error
	if err != nil {
		var menu model.SysBaseMenus

		db := global.GLB_DB.Preload("SysAuthoritys").Where("menu_id = ?", menu_id).First(&menu)
		if db.Error != nil {
			return db.Error
		}

		//删除关联表
		if len(menu.SysAuthoritys) > 0 {
			err = global.GLB_DB.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
			if err != nil {
				return
			}
			//尝试删除
			err = db.Delete(&menu).Error

		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBaseMenu
//@description: 更新路由
//@param: menu model.SysBaseMenu
//@return: err error

func (m *MenusService) UpdateBaseMenu(menu model.SysBaseMenus) (err error) {
	var oldMenu model.SysBaseMenus
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = global.GLB_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("menu_id = ?", menu.MenuId).Find(&oldMenu)

		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("menu_id <> ? AND name = ?", menu.MenuId, menu.Name).First(&model.SysBaseMenus{}).Error, gorm.ErrRecordNotFound) {

				return errors.New("存在相同name修改失败")
			}
		}

		txErr := db.Updates(upDateMap).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBaseMenuById
//@description: 返回当前选中menu
//@param: id uint64
//@return: menu model.SysBaseMenus, err error

func (m *MenusService) GetBaseMenuById(id uint64) (menu model.SysBaseMenus, err error) {
	err = global.GLB_DB.Where("menu_id = ?", id).First(&menu).Error
	if err != nil {
		return
	}

	var allMenus []model.SysBaseMenus
	err = global.GLB_DB.Where("parent_id = ?", menu.MenuId).Order("sort").Find(&allMenus).Error
	menu.Children = allMenus

	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMenuAuthority
//@description: 查看当前角色树
//@param: info *request.GetAuthorityId
//@return: err error, menus []model.ViewAuthorityMenu

// 查询的视图
func (menuService *MenusService) GetMenuByAuthority(info *request.GetAuthorityId) (menus []model.ViewAuthorityMenu, err error) {
	err = global.GLB_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	return menus, err
}
