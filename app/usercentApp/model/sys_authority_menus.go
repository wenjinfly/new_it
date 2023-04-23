// Package models  角色菜单关系表
package model

// 废弃 该表自动创建

// SysAuthorityMenus  角色菜单关系表。
// 说明:
// 表名:sys_authority_menus
type SysAuthorityMenus struct {
	MenuId      uint64 `gorm:"column:menu_id;primaryKey" json:"MenuId"`           //type:         comment:菜单ID    version:2023-03-12 23:15
	AuthorityId string `gorm:"column:authority_id;primaryKey" json:"AuthorityId"` //type:string   comment:角色ID    version:2023-03-12 23:15
}

// TableName 表名:sys_authority_menus，角色菜单关系表。
// 说明:
func (s *SysAuthorityMenus) TableName() string {
	return "sys_authority_menus"
}
