package testexample

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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
