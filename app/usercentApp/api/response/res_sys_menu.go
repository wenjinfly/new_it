package response

import "new_it/app/usercentApp/model"

type SysMenusResponse struct {
	Menus []model.ViewAuthorityMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []model.SysBaseMenus `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu model.SysBaseMenus `json:"menu"`
}
