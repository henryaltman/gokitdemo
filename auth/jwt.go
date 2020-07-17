package auth

import (
	"context"
	"errors"
	"fmt"
	kitJwt "github.com/go-kit/kit/auth/jwt"
	"gokitdemo/router"
	"gokitdemo/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//secret key
var secretKey = []byte("abcd1234!@#$")

// ArithmeticCustomClaims 自定义声明
type BasicCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

// jwtKeyFunc 返回密钥
func JwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return secretKey, nil
}

// Sign 生成token
func GeneratorToken(uid int) (string, error) {

	//为了演示方便，设置两分钟后过期
	expAt := time.Now().Add(time.Duration(2) * time.Minute).Unix()

	// 创建声明
	claims := BasicCustomClaims{
		UserId: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "system",
		},
	}
	//创建token，指定加密算法为HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

//解析token
func GetUidFromContext(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	basicCustomClaims := BasicCustomClaims{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		basicCustomClaims.UserId = int(claims["user_id"].(float64))
	}
	return basicCustomClaims.UserId, err
}

//验证token
func VerifyToken(ctx context.Context) (bool, error) {
	httpPath := ctx.Value(HttpPATH)
	httpPathString := fmt.Sprintf("%v", httpPath)
	httpPathString = util.Ucfirst(httpPathString)
	if _, ok := router.RouterWithoutToken[httpPathString]; ok {
		return true, nil
	}
	jWTtoken := ctx.Value(kitJwt.JWTTokenContextKey)
	tokenString := fmt.Sprintf("%v", jWTtoken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	fmt.Println("tokenString", tokenString)
	if err != nil {
		fmt.Println("VerifyToken err", err.Error())
		return false, err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}
	return false, errors.New("token verify faild")
}
