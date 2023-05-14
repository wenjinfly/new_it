// Package models  菜单
package model

import "time"

// SysBaseMenus  菜单。
// 说明:
// 表名:sys_base_menus
type SysBaseMenus struct {
	MenuId        uint64                            `gorm:"column:menu_id;primaryKey" json:"MenuId"` //type:                  comment:菜单ID              version:2023-03-12 22:57
	MenuLevel     int                               `gorm:"column:menu_level" json:"menuLevel"`      //type:BIGINT UNSIGNED   comment:                    version:2023-03-12 22:57
	ParentId      uint64                            `gorm:"column:parent_id" json:"parentId"`        //type:string            comment:父菜单ID            version:2023-03-12 22:57
	Path          string                            `gorm:"column:path" json:"path"`                 //type:string            comment:路由path            version:2023-03-12 22:57
	Name          string                            `gorm:"column:name" json:"name"`                 //type:string            comment:路由name            version:2023-03-12 22:57
	Hidden        bool                              `gorm:"column:hidden" json:"hidden"`             //type:BIT               comment:是否在列表隐藏      version:2023-03-12 22:57
	Component     string                            `gorm:"column:component" json:"component"`       //type:string            comment:对应前端文件路径    version:2023-03-12 22:57
	Sort          int                               `gorm:"column:sort" json:"sort"`                 //type:int            comment:排序标记            version:2023-03-12 22:57
	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
	CreatedAt     *time.Time                        `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time        comment:                    version:2023-03-12 22:57
	UpdatedAt     *time.Time                        `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time        comment:                    version:2023-03-12 22:57
	DeletedAt     *time.Time                        `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time        comment:                    version:2023-03-12 22:57
	SysAuthoritys []SysAuthorities                  `json:"authoritys" gorm:"many2many:sys_authority_menus;joinForeignKey:MenuId;joinReferences:AuthorityId"`
	Children      []SysBaseMenus                    `json:"children" gorm:"-"` //辅助变量用于生成菜单树
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
}

// TableName 表名:sys_base_menus，菜单。
// 说明:
func (s *SysBaseMenus) TableName() string {
	return "sys_base_menus"
}
