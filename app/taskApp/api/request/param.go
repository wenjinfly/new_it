package request

import "new_it/common"

type ParamTaskID struct {
	TaskId uint64 `json:"TaskId" form:"TaskId"` //
}

type ParamTaskIDStatusPhash struct {
	TaskId     uint64 `json:"TaskId" form:"TaskId"` //
	TaskStatus string `json:"TaskStatus"`           //任务状态
	TaskPhase  string `json:"TaskPhase"`            //任务阶段
	UserId     uint64 //通过token获取
}

type ParamTaskInfoList struct {
	common.PageInfo
	UserId uint64 //通过token获取
}

type ParamUserTaskStatus struct {
	TaskId     uint64 `json:"TaskId" form:"TaskId"` //
	CommStatus string `json:"CommStatus"`           //任务状态
	UserId     uint64 //通过token获取
}

type ParamUserTaskRelation struct {
	Id     uint64 `json:"Id"` //type:  comment:关系id          version:2023-04-02 16:08
	UserId uint64 //通过token获取
}

type ParamGetChatMessage struct {
	common.PageInfo
	FromId    uint64 `json:"FromId"`    //发送人的userid
	ToId      uint64 `json:"ToId"`      //接收人的userid
	TaskID    uint64 `json:"TaskID"`    //任务ID 有的话
	GroupId   uint64 `json:"GroupId"`   //群ID 有的话
	ClassType int    `json:"ClassType"` //'消息分类：1是用户聊天，2是task下记录，3是群组聊天'

	ChatUuid string
}
