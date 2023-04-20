package common

import (
	"crypto/md5"
	"encoding/hex"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5
//@description: md5 哈希
//@param: str []byte
//@return: string

func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
