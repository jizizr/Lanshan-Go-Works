package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const tokenExpireDuration = time.Hour * 24 * 7
const secretKey = "REDROCK114514"

type MyClaims struct {
	Uid int64
	jwt.RegisteredClaims
}

func GenToken(uid int64) (string, error) {
	claim := MyClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secretKey))
}
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
