package common

import (
	"encoding/json"
	"net/http"
	"new_it/errorcode"
)

// 自定义http的状态和，业务状态
func HttpStatusResponse(w http.ResponseWriter, statusCode int, res Response) {

	msg, _ := json.Marshal(res)
	w.WriteHeader(statusCode)
	w.Write(msg)
}

// http协议错误
func HttpErrorErrorResponse(w http.ResponseWriter, statusCode int, errs errorcode.Errno) {
	res := responseError(errs)
	HttpStatusResponse(w, statusCode, *res)
}

// http协议正常，但是业务有错误
func HttpOKErrorResponse(w http.ResponseWriter, errs errorcode.Errno) {
	res := responseError(errs)
	HttpStatusResponse(w, http.StatusOK, *res)
}

// http协议正常，业务也正常
func HttpOKResponse(w http.ResponseWriter, data interface{}) {
	//err
	res := responseError(*errorcode.OK)
	//data
	newres := res.AddData(data)

	HttpStatusResponse(w, http.StatusOK, newres)
}
