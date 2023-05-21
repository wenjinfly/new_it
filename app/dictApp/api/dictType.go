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

type DictTypeAPI struct{}

// AddDictType
//
//	@receiver us
//	@param w
//	@param r
func (us *DictTypeAPI) AddDictType(w http.ResponseWriter, r *http.Request) {
	var dtype model.DictType

	err := common.HttpRequest2Struct(r, &dtype)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.DictTypeService.AddDictType(dtype); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
	} else {
		common.HttpOKResponse(w, nil)
	}

}

// DeleteDictType
//
//	@receiver us
//	@param w
//	@param r
func (us *DictTypeAPI) DeleteDictType(w http.ResponseWriter, r *http.Request) {
	var idInfo request.ParamDictTypeCode
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	var dtype model.DictType
	dtype.TypeCode = idInfo.TypeCode

	if err := service.DictTypeService.DeleteDictType(dtype); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("删除失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}

}

// UpdateDictType
//
//	@receiver us
//	@param w
//	@param r
func (us *DictTypeAPI) UpdateDictType(w http.ResponseWriter, r *http.Request) {
	var dtype model.DictType
	err := common.HttpRequest2Struct(r, &dtype)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.DictTypeService.UpdateDictType(dtype); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}

}

// GetDictTypeByCode
//
//	@receiver us
//	@param w
//	@param r
func (us *DictTypeAPI) GetDictTypeByCode(w http.ResponseWriter, r *http.Request) {
	var idInfo request.ParamDictTypeCode
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if dictdata, err := service.DictTypeService.GetDictTypeByCode(idInfo.TypeCode); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, response.ParamDicttypeResponse{Dicttype: dictdata})
	}

}

// GetDictTypeList
//
//	@receiver us
//	@param w
//	@param r
func (us *DictTypeAPI) GetDictTypeList(w http.ResponseWriter, r *http.Request) {
	var pageInfo request.ParamDictTypeList
	err := common.HttpRequest2Struct(r, &pageInfo)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if list, total, err := service.DictTypeService.GetDictTypeList(pageInfo); err != nil {
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
