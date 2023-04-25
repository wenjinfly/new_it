package api

import (
	"net/http"
	"new_it/app/usercentApp/api/request"
	"new_it/app/usercentApp/model"
	"new_it/app/usercentApp/service"
	"new_it/common"
	"new_it/errorcode"
)

type MenuInfoApi struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getMenu [post]
func (a *MenuInfoApi) GetMenu(w http.ResponseWriter, r *http.Request) {
	/*	if err, menus := service.MenusServices.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
			common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		} else {
			if menus == nil {
				menus = []model.SysBaseMenus{}
			}
			response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
		}
	*/
}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单列表"
// @Router /menu/getBaseMenuTree [post]
func (a *MenuInfoApi) GetBaseMenuTree(w http.ResponseWriter, r *http.Request) {

}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {object} response.Response{msg=string} "增加menu和角色关联关系"
// @Router /menu/addMenuAuthority [post]
func (a *MenuInfoApi) AddMenuAuthority(w http.ResponseWriter, r *http.Request) {
	var authorityMenu request.AddMenuAuthorityInfo
	err := common.HttpRequest2Struct(r, &authorityMenu)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.MenusServices.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

// @Tags AuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取指定角色menu"
// @Router /menu/getMenuAuthority [post]
func (a *MenuInfoApi) GetMenuAuthority(w http.ResponseWriter, r *http.Request) {

}

// @Tags Menu
// @Summary 新增菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "新增菜单"
// @Router /menu/addBaseMenu [post]
func (a *MenuInfoApi) AddBaseMenu(w http.ResponseWriter, r *http.Request) {
	var menu model.SysBaseMenus
	err := common.HttpRequest2Struct(r, &menu)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.MenusServices.AddBaseMenu(menu); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
}

// @Tags Menu
// @Summary 删除菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{msg=string} "删除菜单"
// @Router /menu/deleteBaseMenu [post]
func (a *MenuInfoApi) DeleteBaseMenu(w http.ResponseWriter, r *http.Request) {

}

// @Tags Menu
// @Summary 更新菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "更新菜单"
// @Router /menu/updateBaseMenu [post]
func (a *MenuInfoApi) UpdateBaseMenu(w http.ResponseWriter, r *http.Request) {

}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenuResponse,msg=string} "根据id获取菜单,返回包括系统菜单列表"
// @Router /menu/getBaseMenuById [post]
func (a *MenuInfoApi) GetBaseMenuById(w http.ResponseWriter, r *http.Request) {

}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router /menu/getMenuList [post]
func (a *MenuInfoApi) GetMenuList(w http.ResponseWriter, r *http.Request) {
	var pageInfo common.PageInfo
	err := common.HttpRequest2Struct(r, &pageInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err, menuList, total := service.MenusServices.GetInfoList(); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))

	} else {
		res := common.PageResult{
			List:     menuList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}
		common.HttpOKResponse(w, res)

	}
}
