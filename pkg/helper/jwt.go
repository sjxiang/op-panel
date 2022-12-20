package helper

import (
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sjxiang/op-panel/pkg/constants"
)


type UserClaims struct {
	UserIdentity string `json:"user_identity"`

	jwt.RegisteredClaims
}

func GenerateToken(userIdentitiy string) (string, error) {

	// 签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		UserIdentity: userIdentitiy,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(constants.TokenExpireDuation),
			},
		},
	})
	
	// 密钥
	tokenString, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*UserClaims, error) {
	uc := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, errors.New("token is invalid")
	}

	return uc, nil
}