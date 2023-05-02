package model

import "time"

// Contract  电子合约信息。
// 说明:
// 表名:contract
// group: Contract
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Contract
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Contract
// version:2023-04-02 16:41
type Contract struct {
	ContractId   uint64     `gorm:"column:contract_id;primaryKey" json:"ContractId"`    //type:             comment:合同ID              version:2023-04-02 16:41
	TaskUserId   uint64     `gorm:"column:task_user_id;primaryKey" json:"TaskUserId"`   //type:             comment:发布者用户ID        version:2023-04-02 16:41
	ApplyUserId  uint64     `gorm:"column:apply_user_id;primaryKey" json:"ApplyUserId"` //type:             comment:申请者用户ID        version:2023-04-02 16:41
	TaskId       uint64     `gorm:"column:task_id;primaryKey" json:"TaskId"`            //type:             comment:任务ID                version:2023-04-02 00:05
	TaskReward   float64    `gorm:"column:task_reward" json:"TaskReward"`               //type:*float64     comment:任务酬金            version:2023-04-02 16:41
	TaskBegin    *time.Time `gorm:"column:task_begin" json:"TaskBegin"`                 //type:*time.Time   comment:任务开始时间        version:2023-04-02 16:41
	TaskEnd      *time.Time `gorm:"column:task_end" json:"TaskEnd"`                     //type:*time.Time   comment:任务结束时间        version:2023-04-02 16:41
	TaskStandard string     `gorm:"column:task_standard" json:"TaskStandard"`           //type:string       comment:验收标准            version:2023-04-02 16:41
	ContractMd5  string     `gorm:"column:contract_md5" json:"ContractMd5"`             //type:string       comment:电子合约文件指纹    version:2023-04-02 16:41
	ContractUrl  string     `gorm:"column:contract_url" json:"ContractUrl"`             //type:string       comment:电子合约下载URL     version:2023-04-02 16:41
	CreatedBy    string     `gorm:"column:created_by" json:"CreatedBy"`                 //type:string       comment:创建人              version:2023-04-02 16:41
	CreatedAt    *time.Time `gorm:"column:created_at" json:"CreatedAt"`                 //type:*time.Time   comment:创建时间            version:2023-04-02 16:41
	UpdatedBy    string     `gorm:"column:updated_by" json:"UpdatedBy"`                 //type:string       comment:更新人              version:2023-04-02 16:41
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"UpdatedAt"`                 //type:*time.Time   comment:更新时间            version:2023-04-02 16:41
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"DeletedAt"`                 //type:*time.Time        comment:                version:2023-03-12 00:10

}

// TableName 表名:contract，电子合约信息。
// 说明:
func (c *Contract) TableName() string {
	return "contract"
}
