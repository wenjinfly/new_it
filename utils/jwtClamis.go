package utils

import (
	"new_it/common"
	"new_it/global"
)

func GetUserIDFromJWT(token string) (uint64, error) {

	cl, errs := GetClaims(token)
	if errs != nil {
		return 0, errs
	}

	return cl.ID, errs
}

func GetUserUUIDFromJWT(token string) (string, error) {

	cl, errs := GetClaims(token)
	if errs != nil {
		return "", errs
	}

	return cl.UUID, errs
}

func GetUserNameFromJWT(token string) (string, error) {

	cl, errs := GetClaims(token)
	if errs != nil {
		return "", errs
	}

	return cl.Username, errs
}

func GetUserAuthorityIDFromJWT(token string) (string, error) {
	cl, errs := GetClaims(token)
	if errs != nil {
		return "", errs
	}

	return cl.AuthorityId, errs
}

func GetClaims(token string) (*common.CustomClaims, error) {
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GLB_LOG.Error("从token中获取从jwt解析信息失败, 请检查请求token是否获取正确", err)
	}
	return claims, err
}
