package custom_string

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-common/ctxdata"
	"time"
)

// CalcMD5 字符串md5加密
func CalcMD5(plainText string) string {
	data := []byte(plainText)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GetJwtToken(secretKey string, seconds, userId int64) (string, error) {
	iat := time.Now().Unix()

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}