package request

import "new_it/common"

type ParamDictTypeCode struct {
	TypeCode string `json:"TypeCode" form:"TypeCode"` //
}

type ParamDictInfoCode struct {
	Code string `json:"Code" form:"Code"` //
}

type ParamDictInfoList struct {
	common.PageInfo
	TypeCode string `json:"TypeCode" form:"TypeCode"` //
}
