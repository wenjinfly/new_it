package api

import (
	"net/http"
	dictService "new_it/app/dictApp/service"
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/api/response"
	"new_it/app/taskApp/model"
	"new_it/app/taskApp/service"
	"new_it/common"
	"new_it/errorcode"
	"new_it/utils"
)

type TaskAPI struct{}

// 整体暂时不校验字典编码的有效性

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

func (us *TaskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
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

	var param request.ParamTaskIDStatusPhash
	err = common.HttpRequest2Struct(r, &param)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if param.TaskId == 0 {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("需要Taskid参数有值"))
		return
	}

	param.UserId = userid

	if err = service.TaskService.DeleteTask(param); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("删除失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

func (us *TaskAPI) UpdateTaskInfo(w http.ResponseWriter, r *http.Request) {
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
	if err = service.TaskService.UpdateTaskInfo(task); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}

}

// UpdateTaskStatus
// 更新任务状态
//
//	@receiver us
//	@param w
//	@param r
func (us *TaskAPI) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
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

	var param request.ParamTaskIDStatusPhash
	err = common.HttpRequest2Struct(r, &param)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if !dictService.DictInfoService.CheckDictCodeExist(param.TaskStatus) {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("当前TaskStatus编码不存在"))
		return
	}

	if param.TaskStatus == "" {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("需要TaskStatus参数有值"))
		return
	}

	param.UserId = userid

	if err = service.TaskService.UpdateTaskStatus(param); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

// UpdateTaskPhase
// 更新任务阶段
//
// @receiver us
// @param w
// @param r
func (us *TaskAPI) UpdateTaskPhase(w http.ResponseWriter, r *http.Request) {
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

	var param request.ParamTaskIDStatusPhash
	err = common.HttpRequest2Struct(r, &param)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if param.TaskPhase == "" {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("需要TaskPhase参数有值"))
		return
	}

	param.UserId = userid

	if err = service.TaskService.UpdateTaskPhase(param); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

func (us *TaskAPI) GetTaskByTaskID(w http.ResponseWriter, r *http.Request) {
	var idInfo request.ParamTaskID
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if taskinfo, err := service.TaskService.GetTaskByTaskID(idInfo.TaskId); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, response.ParamTaskResponse{TaskInfo: taskinfo})
	}
}

func (us *TaskAPI) GetTaskListByUserId(w http.ResponseWriter, r *http.Request) {

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

	var pageInfo request.ParamTaskInfoList
	err = common.HttpRequest2Struct(r, &pageInfo)
	pageInfo.UserId = userid

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if list, total, err := service.TaskService.GetTaskListByUserId(pageInfo); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))

	} else {
		common.HttpOKResponse(w, common.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		})
	}
}

func (us *TaskAPI) GetTaskListByName(w http.ResponseWriter, r *http.Request) {
	var pageInfo request.ParamTaskInfoList
	err := common.HttpRequest2Struct(r, &pageInfo)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if list, total, err := service.TaskService.GetTaskListByName(pageInfo); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))

	} else {
		common.HttpOKResponse(w, common.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		})
	}
}
