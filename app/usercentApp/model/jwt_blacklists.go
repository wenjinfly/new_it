// Package models
package model

import "time"

// JwtBlacklists
// 说明:
// 表名:jwt_blacklists
type JwtBlacklists struct {
	Id        uint64     `gorm:"column:id;primaryKey" json:"Id"`     //type:BIGINT UNSIGNED   comment:       version:2023-03-12 23:21
	Jwt       string     `gorm:"column:jwt" json:"Jwt"`              //type:string            comment:jwt    version:2023-03-12 23:21
	CreatedAt *time.Time `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time        comment:       version:2023-03-12 23:21
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time        comment:       version:2023-03-12 23:21
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time        comment:       version:2023-03-12 23:21

}

// TableName 表名:jwt_blacklists，。
// 说明:
func (s *JwtBlacklists) TableName() string {
	return "jwt_blacklists"
}
