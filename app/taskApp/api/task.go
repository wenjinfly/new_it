package api

import (
	"net/http"
	"new_it/app/taskApp/model"
	"new_it/app/taskApp/service"
	"new_it/common"
	"new_it/errorcode"
	"new_it/utils"
)

type TaskAPI struct{}

// AddTask 新增一个任务
//
//	@receiver us
//	@param w
//	@param r
func (us *TaskAPI) AddTask(w http.ResponseWriter, r *http.Request) {
	//拿不到token或是uuid则说明 认证异常
	token, err := common.HttpRequestGetJWTToken(r)
	if err != nil {
		common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}
	//通过token拿到userid
	userid, errs := utils.GetUserIDFromJWT(token)
	if errs != nil {
		common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(errs.Error()))
		return
	}

	var task model.Task

	err = common.HttpRequest2Struct(r, &task)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	task.UserId = userid

	if err := service.TaskService.AddTask(task); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
	} else {
		common.HttpOKResponse(w, nil)
	}
}
