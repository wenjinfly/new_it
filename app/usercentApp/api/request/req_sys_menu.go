package request

import (
	"new_it/app/usercentApp/model"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
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
