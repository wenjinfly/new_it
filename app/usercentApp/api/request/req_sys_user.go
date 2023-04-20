package request

// User login structure
type Login_t struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码,md5(user+password+new_it)
}

// User register structure
type Register struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NickName    string `json:"nickname" gorm:"default:'ITUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}
