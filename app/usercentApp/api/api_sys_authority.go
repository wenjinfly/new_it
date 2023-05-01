package api

import (
	"net/http"
	"new_it/app/usercentApp/api/response"
	"new_it/app/usercentApp/model"
	"new_it/app/usercentApp/service"
	"new_it/common"
	"new_it/errorcode"
)

type AuthorityInfoApi struct{}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "创建角色,返回包括系统角色详情"
// @Router /authority/createAuthority [post]
func (a *AuthorityInfoApi) CreateAuthority(w http.ResponseWriter, r *http.Request) {

	//暂时不验证token
	var authority model.SysAuthorities

	err := common.HttpRequest2Struct(r, &authority)
	if err != nil {
		common.HttpOKErrorResponse(w, *errorcode.ErrBindParam)
		return
	}

	if err, authBack := service.AuthorityServices.CreateAuthority(authority); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("创建失败"+err.Error()))
	} else {

		common.HttpOKResponse(w, response.SysAuthorityResponse{Authority: authBack})

	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "拷贝角色,返回包括系统角色详情"
// @Router /authority/copyAuthority [post]
func (a *AuthorityInfoApi) CopyAuthority(w http.ResponseWriter, r *http.Request) {
	/*
		var copyInfo systemRes.SysAuthorityCopyResponse
		_ = c.ShouldBindJSON(&copyInfo)
		if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if err, authBack := authorityService.CopyAuthority(copyInfo); err != nil {
			global.GVA_LOG.Error("拷贝失败!", zap.Error(err))
			response.FailWithMessage("拷贝失败"+err.Error(), c)
		} else {
			response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
		}*/
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {object} response.Response{msg=string} "删除角色"
// @Router /authority/deleteAuthority [post]
func (a *AuthorityInfoApi) DeleteAuthority(w http.ResponseWriter, r *http.Request) {
	/*var authority system.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}*/
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {object} response.Response{data=systemRes.SysAuthorityResponse,msg=string} "更新角色信息,返回包括系统角色详情"
// @Router /authority/updateAuthority [post]
func (a *AuthorityInfoApi) UpdateAuthority(w http.ResponseWriter, r *http.Request) {
	/*	var auth system.SysAuthority
		_ = c.ShouldBindJSON(&auth)
		if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if err, authority := authorityService.UpdateAuthority(auth); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败"+err.Error(), c)
		} else {
			response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
		}*/
}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取角色列表,返回包括列表,总数,页码,每页数量"
// @Router /authority/getAuthorityList [post]
func (a *AuthorityInfoApi) GetAuthorityList(w http.ResponseWriter, r *http.Request) {
	var pageInfo common.PageInfo
	err := common.HttpRequest2Struct(r, &pageInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err, list, total := service.AuthorityServices.GetAuthorityInfoList(pageInfo); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))
	} else {
		res := common.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}
		common.HttpOKResponse(w, res)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "设置角色资源权限"
// @Success 200 {object} response.Response{msg=string} "设置角色资源权限"
// @Router /authority/setDataAuthority [post]
func (a *AuthorityInfoApi) SetDataAuthority(w http.ResponseWriter, r *http.Request) {
	/*	var auth system.SysAuthority
		_ = c.ShouldBindJSON(&auth)
		if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if err := authorityService.SetDataAuthority(auth); err != nil {
			global.GVA_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败"+err.Error(), c)
		} else {
			response.OkWithMessage("设置成功", c)
		}*/
}
