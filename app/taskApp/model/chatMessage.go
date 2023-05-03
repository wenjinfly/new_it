package model

import (
	"new_it/common"
	"time"
)

// 聊天交流记录
type ChatMessage struct {
	ChatId     uint64            `gorm:"column:chat_id;primaryKey" json:"ChatId"`
	FromId     uint64            `gorm:"column:form_id;primaryKey" json:"FromId"`  //发送人的userid
	FromUuid   string            `gorm:"column:form_uuid;size:64" json:"FromUUID"` //发送人的uuid
	FromName   string            `gorm:"column:form_name;size:64" json:"FromName"` //发送人的名称
	ToId       uint64            `gorm:"column:to_id;primaryKey" json:"ToId"`      //接收人的userid
	ToUuid     string            `gorm:"column:to_uuid;size:64" json:"ToUUID"`     //接收人的uuid
	ToName     string            `gorm:"column:to_name;size:64" json:"ToName"`     //接收人的名称
	SendTime   *common.LocalTime `gorm:"column:send_time" json:"SendTime"`         //消息的发送时间
	Content    string            `gorm:"column:content;type:text" json:"Content"`  //消息内容
	ChatType   int               `gorm:"column:chat_type" json:"ChatType"`         // 消息类型：1是普通文本，2是图片，3是表情
	ClassType  int               `gorm:"column:class_type" json:"ClassType"`       //'消息分类：1是用户聊天，2是task下记录，3是群组聊天'
	ChatStatus int               `gorm:"column:chat_status" json:"ChatStatus"`     //消息的状态：0未读消息；1：已读消息，2消息已撤销
	TaskID     uint64            `gorm:"column:task_id" json:"TaskID"`             //任务ID 有的话
	GroupId    uint64            `gorm:"column:group_id" json:"GroupId"`           //群ID 有的话
	GroupName  string            `gorm:"column:group_name" json:"GroupName"`       //群名称

	CreatedAt *time.Time `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time   comment:创建时间            version:2023-04-02 16:41
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time   comment:更新时间            version:2023-04-02 16:41
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time        comment:                version:2023-03-12 00:10

}

// TableName 表名:contract，电子合约信息。
// 说明:
func (c *ChatMessage) TableName() string {
	return "chat_message"
}
