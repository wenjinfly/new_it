package api

import (
	"net/http"
	"new_it/app/dictApp/api/request"
	"new_it/app/dictApp/api/response"
	"new_it/app/dictApp/model"
	"new_it/app/dictApp/service"
	"new_it/common"
	"new_it/errorcode"
)

type DictInfoAPI struct{}

// AddDictInfo
//
//	@receiver us
//	@param w
//	@param r
func (us *DictInfoAPI) AddDictInfo(w http.ResponseWriter, r *http.Request) {
	var dtype model.DictInfo

	err := common.HttpRequest2Struct(r, &dtype)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.DictInfoService.AddDictInfo(dtype); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
	} else {
		common.HttpOKResponse(w, nil)
	}
}

// DeleteDictInfo
//
//	@receiver us
//	@param w
//	@param r
func (us *DictInfoAPI) DeleteDictInfo(w http.ResponseWriter, r *http.Request) {
	var idInfo request.ParamDictInfoCode
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	var dtype model.DictInfo
	dtype.Code = idInfo.Code

	if err := service.DictInfoService.DeleteDictInfo(dtype); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("删除失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

// UpdateDictInfo
//
//	@receiver us
//	@param w
//	@param r
func (us *DictInfoAPI) UpdateDictInfo(w http.ResponseWriter, r *http.Request) {
	var dtype model.DictInfo
	err := common.HttpRequest2Struct(r, &dtype)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.DictInfoService.UpdateDictInfo(dtype); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

// GetDictInfoByCode
//
//	@receiver us
//	@param w
//	@param r
func (us *DictInfoAPI) GetDictInfoByCode(w http.ResponseWriter, r *http.Request) {
	var idInfo request.ParamDictInfoCode
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if dictdata, err := service.DictInfoService.GetDictInfoByCode(idInfo.Code); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, response.ParamDictInfoResponse{Dictinfo: dictdata})
	}

}

// GetDictInfoList
//
//	@receiver us
//	@param w
//	@param r
func (us *DictInfoAPI) GetDictInfoListByTypeCode(w http.ResponseWriter, r *http.Request) {
	var pageInfo request.ParamDictInfoList
	err := common.HttpRequest2Struct(r, &pageInfo)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if pageInfo.TypeCode == "" {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("TypeCode不能为空"))
		return
	}

	if list, total, err := service.DictInfoService.GetDictInfoListByTypeCode(pageInfo); err != nil {
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
