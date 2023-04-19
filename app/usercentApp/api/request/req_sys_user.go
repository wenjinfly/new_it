package request

// User login structure
type Login_t struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码,md5(user+password+new_it)
}

// User register structure
type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'ITUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}
