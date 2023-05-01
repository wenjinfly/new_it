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
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "MainIndex", Name: "MainIndex", Component: "view/task/index.vue", Sort: 1, Title: "主页面", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 2, Title: "关于我们", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "admin", Name: "admin", Component: "view/superAdmin/index.vue", Sort: 3, Title: "超级管理员", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "MenuManager", Name: "MenuManager", Component: "view/superAdmin/menu/menu.vue", Sort: 31, Title: "菜单管理", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "AuthorityManager", Name: "AuthorityManager", Component: "view/superAdmin/authority/authority.vue", Sort: 32, Title: "角色管理", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 3, Path: "UserManager", Name: "UserManager", Component: "view/superAdmin/user/user.vue", Sort: 33, Title: "用户管理", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "Person", Name: "Person", Component: "view/person/person.vue", Sort: 4, Title: "个人信息", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "PersonalTask", Name: "PersonalTask", Component: "view/person/PersonalTask.vue", Sort: 5, Title: "个人任务", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "TaskDetails", Name: "TaskDetails", Component: "view/task/Details.vue", Sort: 6, Title: "任务详情", KeepAlive: true},
		{MenuLevel: 0, Hidden: false, ParentId: 0, Path: "PublishTask", Name: "PublishTask", Component: "view/task/publishtask.vue", Sort: 7, Title: "发布任务", KeepAlive: true},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.New(m.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (m *MENU) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("path = ?", "MainIndex").First(&model.SysBaseMenus{}).Error, gorm.ErrRecordNotFound) {
		// 判断是否存在数据
		return false
	}
	return true
}
