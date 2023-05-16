package dbdata

import (
	"errors"
	"new_it/app/usercentApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

var BaseMenu = new(MENU)

type MENU struct{}

func (m *MENU) TableName() string {
	return "sys_base_menus"
}

func (m *MENU) Initialize() error {

	entities := []model.SysBaseMenus{
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "/home", Name: "home", Component: "views/task/index.vue", Sort: 1, Meta: model.Meta{Title: "主页面", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "/about", Name: "about", Component: "views/about/index.vue", Sort: 2, Meta: model.Meta{Title: "关于我们", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "/admin", Name: "admin", Component: "views/admin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "/menu", Name: "menu", Component: "views/admin/menu/index.vue", Sort: 31, Meta: model.Meta{Title: "菜单管理", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "/role", Name: "role", Component: "views/admin/role/index.vue", Sort: 32, Meta: model.Meta{Title: "角色管理", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "/user", Name: "user", Component: "views/admin/user/index.vue", Sort: 33, Meta: model.Meta{Title: "用户管理", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "/dictType", Name: "dictType", Component: "views/admin/dictionary/dictType.vue", Sort: 34, Meta: model.Meta{Title: "字典类型管理", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "/dictInfo", Name: "dictInfo", Component: "views/admin/dictionary/dictInfo.vue", Sort: 35, Meta: model.Meta{Title: "字典字段管理", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "/person", Name: "person", Component: "views/person/index.vue", Sort: 4, Meta: model.Meta{Title: "个人信息", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 1, Path: "/personalTask", Name: "personalTask", Component: "views/task/personalTask/index.vue", Sort: 11, Meta: model.Meta{Title: "个人任务", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 1, Path: "/taskDetails", Name: "taskDetails", Component: "views/task/taskDetails/index.vue", Sort: 12, Meta: model.Meta{Title: "任务详情", KeepAlive: true}},
		{MenuLevel: 0, Hidden: false, ParentId: 1, Path: "/publishTask", Name: "publishTask", Component: "views/task/publishTask/index.vue", Sort: 13, Meta: model.Meta{Title: "发布任务", KeepAlive: true}},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.New(m.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (m *MENU) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("name = ?", "home").First(&model.SysBaseMenus{}).Error, gorm.ErrRecordNotFound) {
		// 判断是否存在数据
		return false
	}
	return true
}
