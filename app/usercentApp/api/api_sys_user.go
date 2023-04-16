package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"new_it/app/usercentApp/api/request"
	"new_it/app/usercentApp/model"
	"new_it/app/usercentApp/service"
	"new_it/common"
	"new_it/global"
	"new_it/utils"
	"strings"
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

	u := &model.SysUsers{UserName: l.Username, Password: l.Password}
	if err, user := service.UserServices.Login(u); err != nil {
		//response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		fmt.Println(user)
		token, _ = getNewToken(user)
	}

	//3 返回token
	fmt.Println("token:", token)
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(token))
}

func getNewToken(user *model.SysUsers) (string, error) {

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

	return token, err
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()       //解析参数，默认是不会解析的
	//这些信息是输出到服务器端的打印信息
	fmt.Println("Request解析")
	//HTTP方法
	fmt.Println("method", r.Method)
	// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
	fmt.Println("RequestURI", r.RequestURI)
	//URL类型,下方分别列出URL的各成员
	fmt.Println("URL_scheme", r.URL.Scheme)
	fmt.Println("URL_opaque", r.URL.Opaque)
	fmt.Println("URL_user", r.URL.User.String())
	fmt.Println("URL_host", r.URL.Host)
	fmt.Println("URL_path", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)
	//协议版本
	fmt.Println("proto", r.Proto)
	fmt.Println("protomajor", r.ProtoMajor)
	fmt.Println("protominor", r.ProtoMinor)
	//HTTP请求的头域
	for k, v := range r.Header {
		// fmt.Println("Header key:" + k)
		for _, vv := range v {
			fmt.Println("header key:" + k + "  value:" + vv)
		}
	}
	//判断是否multipart方式
	is_multipart := false
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			is_multipart = true
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	fmt.Println("body :", string(body))

	//解析body
	if is_multipart == true {
		r.ParseMultipartForm(128)
		fmt.Println("解析方式:ParseMultipartForm")
	} else {
		r.ParseForm()
		fmt.Println("解析方式:ParseForm")
	}
	//body内容长度
	fmt.Println("ContentLength", r.ContentLength)
	//是否在回复请求后关闭连接
	fmt.Println("Close", r.Close)
	//HOSt
	fmt.Println("host", r.Host)
	//form
	fmt.Println("Form", r.Form)
	//postform
	fmt.Println("PostForm", r.PostForm)
	//MultipartForm
	fmt.Println("MultipartForm", r.MultipartForm)
	//解析MultipartForm内的文件
	if is_multipart {
		files := r.MultipartForm.File
		for k, v := range files {
			fmt.Println(k)
			for _, vv := range v {
				fmt.Println(k + ":" + vv.Filename)
			}
		}
	}

	//该请求的来源地址
	fmt.Println("RemoteAddr", r.RemoteAddr)

	fmt.Fprintf(w, "hello astaxie!") //这个写入到w的是输出到客户端的
}
