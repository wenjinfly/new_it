// Package models  菜单
package model

import "time"

// SysBaseMenus  菜单。
// 说明:
// 表名:sys_base_menus
type SysBaseMenus struct {
	MenuId        uint64           `gorm:"column:menu_id;primaryKey" json:"MenuId"` //type:                  comment:菜单ID              version:2023-03-12 22:57
	MenuLevel     int              `gorm:"column:menu_level" json:"MenuLevel"`      //type:BIGINT UNSIGNED   comment:                    version:2023-03-12 22:57
	ParentId      int              `gorm:"column:parent_id" json:"ParentId"`        //type:string            comment:父菜单ID            version:2023-03-12 22:57
	Path          string           `gorm:"column:path" json:"Path"`                 //type:string            comment:路由path            version:2023-03-12 22:57
	Name          string           `gorm:"column:name" json:"Name"`                 //type:string            comment:路由name            version:2023-03-12 22:57
	Hidden        bool             `gorm:"column:hidden" json:"Hidden"`             //type:BIT               comment:是否在列表隐藏      version:2023-03-12 22:57
	Component     string           `gorm:"column:component" json:"Component"`       //type:string            comment:对应前端文件路径    version:2023-03-12 22:57
	Sort          int              `gorm:"column:sort" json:"Sort"`                 //type:int            comment:排序标记            version:2023-03-12 22:57
	KeepAlive     bool             `gorm:"column:keep_alive" json:"KeepAlive"`      //type:BIT               comment:附加属性            version:2023-03-12 22:57
	DefaultMenu   bool             `gorm:"column:default_menu" json:"DefaultMenu"`  //type:BIT               comment:附加属性            version:2023-03-12 22:57
	Title         string           `gorm:"column:title" json:"Title"`               //type:string            comment:附加属性            version:2023-03-12 22:57
	Icon          string           `gorm:"column:icon" json:"Icon"`                 //type:string            comment:附加属性            version:2023-03-12 22:57
	CloseTab      bool             `gorm:"column:close_tab" json:"CloseTab"`        //type:BIT               comment:附加属性            version:2023-03-12 22:57
	CreatedAt     *time.Time       `gorm:"column:created_at" json:"CreatedAt"`      //type:*time.Time        comment:                    version:2023-03-12 22:57
	UpdatedAt     *time.Time       `gorm:"column:updated_at" json:"UpdatedAt"`      //type:*time.Time        comment:                    version:2023-03-12 22:57
	DeletedAt     *time.Time       `gorm:"column:deleted_at" json:"DeletedAt"`      //type:*time.Time        comment:                    version:2023-03-12 22:57
	SysAuthoritys []SysAuthorities `json:"authoritys" gorm:"many2many:sys_authority_menus;joinForeignKey:MenuId;joinReferences:AuthorityId"`
	Children      []SysBaseMenus   `json:"children" gorm:"-"` //辅助变量用于生成菜单树
}

// TableName 表名:sys_base_menus，菜单。
// 说明:
func (s *SysBaseMenus) TableName() string {
	return "sys_base_menus"
}
