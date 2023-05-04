package api

import (
	"net/http"
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/model"
	"new_it/app/taskApp/service"
	"new_it/common"
	"new_it/errorcode"
	"new_it/utils"
)

type UserTaskRelationAPI struct{}

func (us *UserTaskRelationAPI) AddRelation(w http.ResponseWriter, r *http.Request) {
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

	var task model.UserTaskRelation

	err = common.HttpRequest2Struct(r, &task)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	task.ApplyUserId = userid

	if err := service.UserTaskService.AddRelation(task); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
	} else {
		common.HttpOKResponse(w, nil)
	}
}

func (us *UserTaskRelationAPI) DeleteRelation(w http.ResponseWriter, r *http.Request) {
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

	var param request.ParamUserTaskRelation
	err = common.HttpRequest2Struct(r, &param)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	param.UserId = userid

	if err = service.UserTaskService.DeleteRelation(param); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("删除失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

func (us *UserTaskRelationAPI) UpdateRelationStatus(w http.ResponseWriter, r *http.Request) {
	//
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

	var param request.ParamUserTaskStatus
	err = common.HttpRequest2Struct(r, &param)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if param.CommStatus == "" {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("需要TaskStatus参数有值"))
		return
	}

	param.UserId = userid

	if err = service.UserTaskService.UpdateRelationStatus(param); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

func (us *UserTaskRelationAPI) GetRelationListByUserId(w http.ResponseWriter, r *http.Request) {
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

	if list, total, err := service.UserTaskService.GetRelationListByUserId(pageInfo); err != nil {
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
