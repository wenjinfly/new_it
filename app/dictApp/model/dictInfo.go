package model

import "time"

type DictInfo struct {
	Code      string     `gorm:"column:code;primaryKey" json:"code"`          //type:             comment:编码                version:2023-04-02 00:05
	Name      string     `gorm:"column:name" json:"name"`                     //type:string       comment:名称              version:2023-04-02 00:05
	CNName    string     `gorm:"column:cn_name" json:"cn_name"`               //type:string       comment:中文名          version:2023-04-02 00:05
	Fixed     bool       `gorm:"column:fixed" json:"fixed"`                   //type:BIT               comment:附加属性            version:2023-03-12 22:57
	TypeCode  string     `gorm:"column:type_code;comment:编码" json:"TypeCode"` //type:             comment:编码                version:2023-04-02 00:05
	TypeDict  DictType   `gorm:"foreignKey:TypeCode;references:TypeCode;comment:字典类型" json:"DictType"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time   comment:创建时间              version:2023-04-02 00:05
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time   comment:更新时间              version:2023-04-02 00:05
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time        comment:                version:2023-03-12 00:10
}

// TableName 表名:Task，任务单。
// 说明:
func (t *DictInfo) TableName() string {
	return "dict_info"
}
