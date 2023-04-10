package model

type JWT struct {
	SigningKey  string // jwt签名
	ExpiresTime int64  // 过期时间
	BufferTime  int64  // 缓冲时间
	Issuer      string // 签发者
}
