package jwtauth

import (
	"admin-server/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var jwtKey []byte

type Claims struct {
	AdminId int64
	jwt.StandardClaims
}

func init() {
	jwtKey = []byte(viper.GetString("server.jwtKey"))
}

func ReleaseToken(admin model.Admin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //token有效期为7天
	claims := &Claims{
		AdminId: admin.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //token有效期为7天
			IssuedAt:  time.Now().Unix(),     //发放token的时间为当前时间
			Issuer:    "admin-server",        //这个token是admin-server发放的
			Subject:   "admin token",         //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey) //使用jwtKey这个密钥来生成token
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 从tokenString解析出claims

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err

}
