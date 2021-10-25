package types

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TOKEN_SECRETKEY = "s3:AuthorizationToken"
)

type S3AuthorizationToken struct {
	jwt.StandardClaims
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Path      string `json:"path"`
}

func CreateToken(accessKey, secretKey, path string, expiresIn int64) (tokenString string, err error) {
	claims := &S3AuthorizationToken{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Duration(expiresIn) * time.Second).Unix()),
			Issuer:    accessKey,
		},
		accessKey,
		secretKey,
		path,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(TOKEN_SECRETKEY))
	return
}

func ParseToken(tokenSrt string) (claims jwt.Claims, err error) {
	token, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(TOKEN_SECRETKEY), nil
	})

	claims = token.Claims
	return
}

func GetTokenDataFromJwtClaims(claims jwt.Claims) (map[string]interface{}, error) {
	var tokenMap map[string]interface{}
	switch claims.(type) {
	case jwt.MapClaims:
		tokenMap = claims.(jwt.MapClaims)
	default:
		return tokenMap, errors.New("get token error")
	}
	return tokenMap, nil
}

func ValidNamespace(validNamespaces, namespace string) bool {
	ns := strings.Split(validNamespaces, ",")
	for _, n := range ns {
		if n == namespace {
			return true
		}
	}

	return false
}
