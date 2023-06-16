package middleware

import (
	"admin-server/internal/model"
	jwtauth "admin-server/pkg/auth"
	"admin-server/pkg/orm"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//认证token中间件

func Authmiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = ctx.DefaultQuery("x-token", "")
		} else {
			tokenString = tokenString[7:]
		}

		//validate token formate
		//如果这个token为空或都不是以Bearer格式开头说明不是一个合格格式的token
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1007, "msg": "TOKEN格式错误,请重新登陆"})
			ctx.Abort() //将这一次请求抛弃
			return
		}

		token, claims, err := jwtauth.ParseToken(tokenString)
		if err != nil || !token.Valid { //token.Valid解析后的token是否有效是一个bool类型
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1005, "msg": "TOKEN不正确或己过期,请重新登陆"})
			ctx.Abort() //将这一次请求抛弃
			return
		}

		//验证通过后获取claims中的AdminId
		adminId := claims.AdminId
		DB := orm.GetDBWithContext(ctx.Request.Context())
		var admin model.Admin
		DB.First(&admin, adminId)

		//用户
		if admin.Id == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1008, "msg": "该用户无权限"})
			ctx.Abort() //将这一次请求抛弃
			return
		}

		//用户存在将用户信息写入上下文
		ctx.Set("admin", admin)
		ctx.Next()
	}

}
