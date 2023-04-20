package model

import "time"

// Package models  用户表

// SysUsers  用户表。
// 说明:
// 表名:sys_users
// group: SysUsers
type SysUsers struct {
	UserId       uint64     `gorm:"column:user_id;primaryKey" json:"UserId"`  //type:BIGINT UNSIGNED   comment:用户ID          version:2023-03-12 00:10
	CreatedAt    *time.Time `gorm:"column:created_at" json:"CreatedAt"`       //type:*time.Time        comment:                version:2023-03-12 00:10
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"UpdatedAt"`       //type:*time.Time        comment:                version:2023-03-12 00:10
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"DeletedAt"`       //type:*time.Time        comment:                version:2023-03-12 00:10
	UserName     string     `gorm:"column:user_name" json:"UserName"`         //type:string            comment:用户登录名      version:2023-03-12 00:10
	Uuid         string     `gorm:"column:uuid" json:"Uuid"`                  //type:string            comment:用户uuid        version:2023-03-12 00:10
	HeaderImg    string     `gorm:"column:header_img" json:"HeaderImg"`       //type:string            comment:用户头像        version:2023-03-12 00:10
	Password     string     `gorm:"column:password" json:"Password"`          //type:string            comment:用户登录密码    version:2023-03-12 00:10
	NickName     string     `gorm:"column:nick_name" json:"NickName"`         //type:string            comment:用户昵称        version:2023-03-12 00:10
	AuthorityId  string     `gorm:"column:authority_id" json:"AuthorityId"`   //type:string            comment:用户角色ID      version:2023-03-12 00:10
	Phone        string     `gorm:"column:phone" json:"Phone"`                //type:string            comment:用户手机号      version:2023-03-12 00:10
	Email        string     `gorm:"column:email" json:"Email"`                //type:string            comment:用户邮箱        version:2023-03-12 00:10
	Gender       string     `gorm:"column:gender" json:"Gender"`              //type:string            comment:性别            version:2023-03-12 00:10
	IdentityCard string     `gorm:"column:identity_card" json:"IdentityCard"` //type:string            comment:身份证号        version:2023-03-12 00:10
	Name         string     `gorm:"column:name" json:"Name"`                  //type:string            comment:真实姓名        version:2023-03-12 00:10
}

// TableName 表名:sys_users，用户表。
// 说明:
func (s *SysUsers) TableName() string {
	return "sys_users"
}
