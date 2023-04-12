// Package models  用户角色关系表
package model

// SysUserAuthority  用户角色关系表。
// 说明:
// 表名:sys_user_authority
// group: SysUserAuthority
type SysUserAuthority struct {
	Id          uint64 `gorm:"column:id;primaryKey" json:"Id"`                    //type:         comment:关系ID    version:2023-03-12 23:08
	UserId      uint64 `gorm:"column:user_id;primaryKey" json:"UserId"`           //type:         comment:用户ID    version:2023-03-12 23:08
	AuthorityId string `gorm:"column:authority_id;primaryKey" json:"AuthorityId"` //type:string   comment:角色ID    version:2023-03-12 23:08
}

// TableName 表名:sys_user_authority，用户角色关系表。
// 说明:
func (s *SysUserAuthority) TableName() string {
	return "sys_user_authority"
}
