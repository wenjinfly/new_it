package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"new_it/app/usercentApp/api/request"
	"new_it/app/usercentApp/api/response"
	"new_it/app/usercentApp/model"
	"new_it/app/usercentApp/service"
	"new_it/common"
	"new_it/global"
	"new_it/utils"
)

type UsercentApi struct{}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]
func (us *UsercentApi) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	//1 获取传入的用户信息
	var l request.Login_t

	v2, ok := r.Header["Content-Type"]
	if !ok {
		fmt.Println("Content-Type", v2)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	fmt.Println("body :", string(body))
	if err := json.Unmarshal(body, &l); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}
	fmt.Println("body struct:", l)

	//2 从数据库查看是否存在
	var token string
	var ExpiresAt int64
	var msg []byte

	u := &model.SysUsers{UserName: l.Username, Password: l.Password}
	if err, user := service.UserServices.Login(u); err != nil {
		//response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		fmt.Println(user)
		token, ExpiresAt, _ = getNewToken(user)

		user.Password = "xxx"
		user.IdentityCard = "111"

		res := response.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: ExpiresAt * 1000,
		}

		msg, _ = json.Marshal(res)
	}

	//3 返回token
	//fmt.Println("token:", token)
	w.WriteHeader(http.StatusOK)

	w.Write(msg)

}

func getNewToken(user *model.SysUsers) (string, int64, error) {

	j := &utils.JWT{SigningKey: []byte(global.GLB_CFG_INFO.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(common.BaseClaims{
		UUID:        user.Uuid,
		ID:          user.UserId,
		Username:    user.UserName,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		//global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		//response.FailWithMessage("获取token失败", c)
		//return
	}

	return token, claims.StandardClaims.ExpiresAt, err
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func (us *UsercentApi) Register(w http.ResponseWriter, r *http.Request) {
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.ChangePasswordStruct true "用户名, 原密码, 新密码"
// @Success 200 {object} response.Response{msg=string} "用户修改密码"
// @Router /user/changePassword [post]
func (us *UsercentApi) ChangePassword(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router /user/getUserList [post]
func (us *UsercentApi) GetUserList(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 更改用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthority [post]
func (us *UsercentApi) SetUserAuthority(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuthorities true "用户UUID, 角色ID"
// @Success 200 {object} response.Response{msg=string} "设置用户权限"
// @Router /user/setUserAuthorities [post]
func (us *UsercentApi) SetUserAuthorities(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /user/deleteUser [delete]
func (us *UsercentApi) DeleteUser(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/setUserInfo [put]
func (us *UsercentApi) SetUserInfo(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "设置用户信息"
// @Router /user/SetSelfInfo [put]
func (us *UsercentApi) SetSelfInfo(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取用户信息"
// @Router /user/getUserInfo [get]
func (us *UsercentApi) GetUserInfo(w http.ResponseWriter, r *http.Request) {}

// @Tags SysUser
// @Summary 重置用户密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body system.SysUser true "ID"
// @Success 200 {object} response.Response{msg=string} "重置用户密码"
// @Router /user/resetPassword [post]
func (us *UsercentApi) ResetPassword(w http.ResponseWriter, r *http.Request) {}
