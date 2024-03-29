package mEncrypt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 生成token
type Claims struct {
	Message string
	jwt.StandardClaims
}

type NewTokenOpt struct {
	SecretKey string    // key
	ExpiresAt time.Time // 过期时间
	Message   string
	Issuer    string // 签名颁发者
	Subject   string // 签名主题
}

type TokenObj struct {
	SecretKey []byte    // key
	ExpiresAt time.Time // 过期时间
	Message   string
	Issuer    string // 签名颁发者
	Subject   string // 签名主题
}

func NewToken(opt NewTokenOpt) *TokenObj {
	var tokenObj TokenObj
	tokenObj.SecretKey = []byte(opt.SecretKey)
	tokenObj.ExpiresAt = opt.ExpiresAt
	tokenObj.Message = opt.Message
	tokenObj.Issuer = opt.Issuer
	tokenObj.Subject = opt.Subject

	if len(tokenObj.Issuer) < 1 {
		tokenObj.Issuer = "goTools"
	}
	if len(tokenObj.Subject) < 1 {
		tokenObj.Subject = "Token"
	}
	return &tokenObj
}

func (Obj *TokenObj) Generate() string {
	claims := Claims{
		Message: Obj.Message,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Obj.ExpiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    Obj.Issuer,
			Subject:   Obj.Subject,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(Obj.SecretKey)
	if err != nil {
		return ""
	}
	return token
}

func ParseToken(tokenString string, SecretKey string) (Claims, bool) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		Claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil || !token.Valid {
		return *Claims, false
	}

	return *Claims, token.Valid
}
