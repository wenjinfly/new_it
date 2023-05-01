package model

import "time"

type DictType struct {
	TypeCode   string     `gorm:"column:type_code;primaryKey" json:"TypeCode"` //type:             comment:编码                version:2023-04-02 00:05
	TypeName   string     `gorm:"column:type_name" json:"TypeName"`            //type:string       comment:名称              version:2023-04-02 00:05
	TypeCNName string     `gorm:"column:type_cn_name" json:"TypeCNName"`       //type:string       comment:中文名          version:2023-04-02 00:05
	CreatedAt  *time.Time `gorm:"column:created_at" json:"CreatedAt"`          //type:*time.Time   comment:创建时间              version:2023-04-02 00:05
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"UpdatedAt"`          //type:*time.Time   comment:更新时间              version:2023-04-02 00:05
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"DeletedAt"`          //type:*time.Time        comment:                version:2023-03-12 00:10
}

// TableName 表名:Task，任务单。
// 说明:
func (t *DictType) TableName() string {
	return "dict_type"
}
