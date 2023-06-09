package model

import "time"

// Package models  角色表

// SysAuthorities  角色表。
// 说明:
// 表名:sys_authorities
type SysAuthorities struct {
	AuthorityId   string           `gorm:"column:authority_id;not null;unique;primaryKey" json:"AuthorityId"` //type:string       comment:角色ID      version:2023-03-12 22:39
	AuthorityName string           `gorm:"column:authority_name" json:"AuthorityName"`                        //type:string       comment:角色名      version:2023-03-12 22:39
	ParentId      string           `gorm:"column:parent_id" json:"ParentId"`                                  //type:string       comment:父角色ID    version:2023-03-12 22:39
	DefaultRouter string           `gorm:"column:default_router" json:"DefaultRouter"`                        //type:string       comment:默认菜单    version:2023-03-12 22:39
	CreatedAt     *time.Time       `gorm:"column:created_at" json:"CreatedAt"`                                //type:*time.Time   comment:            version:2023-03-12 22:39
	UpdatedAt     *time.Time       `gorm:"column:updated_at" json:"UpdatedAt"`                                //type:*time.Time   comment:            version:2023-03-12 22:39
	DeletedAt     *time.Time       `sql:"index"`
	SysBaseMenus  []SysBaseMenus   `json:"menus" gorm:"many2many:sys_authority_menus;joinForeignKey:AuthorityId;joinReferences:MenuId"`
	Users         []SysUsers       `json:"-" gorm:"many2many:sys_user_authority;joinForeignKey:AuthorityId;joinReferences:UserId"`
	Children      []SysAuthorities `json:"children" gorm:"-"` //辅助变量用于生成角色树
}

// TableName 表名:sys_authorities，角色表
// 说明:
func (s *SysAuthorities) TableName() string {
	return "sys_authorities"
}
