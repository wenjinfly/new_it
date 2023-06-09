package api

import (
	"net/http"
	"new_it/app/usercentApp/api/request"
	"new_it/app/usercentApp/api/response"
	"new_it/app/usercentApp/model"
	"new_it/app/usercentApp/service"
	"new_it/common"
	"new_it/errorcode"
	"new_it/utils"
)

type MenuInfoApi struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /menu/getViewMenu [post]
func (a *MenuInfoApi) GetViewMenu(w http.ResponseWriter, r *http.Request) {

	//拿不到token或是uuid则说明 认证异常
	token, err := common.HttpRequestGetJWTToken(r)
	if err != nil {
		common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	authid, errs := utils.GetUserAuthorityIDFromJWT(token)
	if errs != nil {
		common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(errs.Error()))
		return
	}

	if err, menus := service.MenusServices.GetMenuTree(authid); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
	} else {
		if menus == nil {
			menus = []model.ViewAuthorityMenu{}
		}
		common.HttpOKResponse(w, response.SysMenusResponse{Menus: menus})
	}

}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单列表"
// @Router /menu/getBaseMenuTree [post]
func (a *MenuInfoApi) GetBaseMenuTree(w http.ResponseWriter, r *http.Request) {
	if menus, err := service.MenusServices.GetBaseMenuTree(); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, response.SysBaseMenusResponse{Menus: menus})
	}
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

	//获取menu
	menu, err2 := service.MenusServices.GetBaseMenuById(authorityMenu.MenuId)
	if err2 != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err2.Error()))
		return
	}

	var menus []model.SysBaseMenus
	menus = append(menus, menu)

	if len(menu.Children) > 0 {
		menus = append(menus, menu.Children...)
	}

	if err := service.MenusServices.AddMenuAuthority(menus, authorityMenu.AuthorityId); err != nil {
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
// @Router /menu/getMenusByAuthority [post]
func (a *MenuInfoApi) GetMenusByAuthority(w http.ResponseWriter, r *http.Request) {
	var param request.GetAuthorityId
	err := common.HttpRequest2Struct(r, &param)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if viewAuthMenus, err := service.MenusServices.GetMenuByAuthority(&param); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))
	} else {
		common.HttpOKResponse(w, response.SysMenusResponse{Menus: viewAuthMenus})
	}
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
	var idInfo request.GetByMenuId
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.MenusServices.DeleteBaseMenu(idInfo.MenuId); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("删除失败-"+err.Error()))
	} else {
		common.HttpOKResponse(w, nil)
	}
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
	var menu model.SysBaseMenus
	err := common.HttpRequest2Struct(r, &menu)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if err := service.MenusServices.UpdateBaseMenu(menu); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("更新失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, nil)
	}
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
	var idInfo request.GetByMenuId
	err := common.HttpRequest2Struct(r, &idInfo)
	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if menu, err := service.MenusServices.GetBaseMenuById(idInfo.MenuId); err != nil {

		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("获取失败-"+err.Error()))

	} else {
		common.HttpOKResponse(w, response.SysBaseMenuResponse{Menu: menu})
	}
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
