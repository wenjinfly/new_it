package common

import (
	"github.com/dgrijalva/jwt-go"
	//uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	//UUID        uuid.UUID
	UUID        string
	ID          uint64
	Username    string
	AuthorityId string
}
