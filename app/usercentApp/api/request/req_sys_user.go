package request

// User login structure
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码,md5(user+password+new_it)
}
