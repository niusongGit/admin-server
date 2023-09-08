package middleware

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/pkg/casbinrbac"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		admin, _ := c.Get("admin")
		adminInfo := admin.(model.Admin)
		//获取uri
		//path := c.Request.URL.Path
		//fmt.Printf(path)
		sub := fmt.Sprintf(consts.SubjectRole, adminInfo.RoleId)
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		success, _ := casbinrbac.GetCasbinEnforcer().Enforce(sub, obj, act)
		if !success {
			c.JSON(http.StatusForbidden, gin.H{"code": 1008, "msg": "该用户无权限"})
			c.Abort()
			return
		} else {
			c.Next()
		}

	}
}
