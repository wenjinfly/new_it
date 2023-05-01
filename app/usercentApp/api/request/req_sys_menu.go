package request

import (
	"new_it/app/usercentApp/model"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	MenuId      uint64 `json:"menuId"`      //type:                  comment:菜单ID              version:2023-03-12 22:57
	AuthorityId string `json:"authorityId"` // 角色ID
}

type AddMenuAuthorityInfo2 struct {
	Menus       []model.SysBaseMenus `json:"menus"`
	AuthorityId string               `json:"authorityId"` // 角色ID
}

func DefaultMenu() []model.SysBaseMenus {
	return []model.SysBaseMenus{{
		MenuLevel: 0,
		ParentId:  0,
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Title:     "仪表盘",
		Icon:      "setting",
	}}
}

// GetById Find by id structure
type GetByMenuId struct {
	MenuId uint64 `json:"MenuId" form:"MenuId"` // 主键ID
}
