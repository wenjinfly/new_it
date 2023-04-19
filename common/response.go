package common

import (
	"encoding/json"
	"new_it/errorcode"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

func (res *Response) AddMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

func (res *Response) AddData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func responseError(errs errorcode.Errno) *Response {
	return &Response{
		Code: errs.Code,
		Msg:  errs.Message,
		Data: nil,
	}
}
