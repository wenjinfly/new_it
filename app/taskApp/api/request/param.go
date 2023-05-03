package request

import "new_it/common"

type ParamTaskID struct {
	TaskId uint64 `json:"TaskId" form:"TaskId"` //
}

type ParamTaskInfoList struct {
	common.PageInfo
	UserId uint64 `json:"UserId" form:"UserId"`
}
