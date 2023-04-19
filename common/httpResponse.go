package common

import (
	"encoding/json"
	"net/http"
	"new_it/errorcode"
)

func HttpResponse(w http.ResponseWriter, statusCode int, res Response) {

	msg, _ := json.Marshal(res)

	w.WriteHeader(statusCode)

	w.Write(msg)

}

func HttpErrorResponse(w http.ResponseWriter, errs errorcode.Errno) {
	w.WriteHeader(http.StatusOK)
	res := responseError(errs)
	msg, _ := json.Marshal(res)

	w.Write(msg)
}

func HttpOKResponse(w http.ResponseWriter, data interface{}) {

	w.WriteHeader(http.StatusOK)

	res := responseError(*errorcode.OK)
	newres := res.AddData(data)

	msg, _ := json.Marshal(newres)

	w.Write(msg)
}
