package admin

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/response"
	"admin-server/internal/schema"
	"admin-server/internal/service"
	jwtauth "admin-server/pkg/auth"
	"admin-server/pkg/errmsg"
	"admin-server/pkg/orm"
	"admin-server/pkg/validator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Admin struct {
}

func (c Admin) Register(ctx *gin.Context) {

	var req = schema.AdminRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	s := service.NewAdminService(orm.GetDBWithContext(ctx.Request.Context()))
	// 判断帐号是否存在
	if s.IsAdminExist(req.AdminName) {
		response.Fail(ctx, errmsg.ERROR_USERNAME_USED, nil)
		return
	}
	// 创建用户
	// 加密用户密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(ctx, errmsg.ERROR_PASSWORD_ENCRYPTION, nil)
		return
	}

	if err := s.Add(&model.Admin{AdminName: req.AdminName, Password: string(hashedPassword)}); err != nil {
		response.Response(ctx, http.StatusInternalServerError, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "注册成功")
}

func (c Admin) Login(ctx *gin.Context) {
	var req = schema.AdminRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	db := orm.GetDBWithContext(ctx.Request.Context())
	//获取参数

	// 判断手机号是否存在
	var admin model.Admin
	db.Where("admin_name = ?", req.AdminName).First(&admin)
	if admin.Id == 0 {
		//用户不存在
		response.Fail(ctx, errmsg.ERROR_USER_NOT_EXIST, nil)
		return
	}

	// 禁用用户不能登录
	if consts.AdminStatusOff.Equals(admin.Status) {
		response.Fail(ctx, errmsg.ERROR_USER_DISABLE, nil)
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		response.Fail(ctx, errmsg.ERROR_PASSWORD_WRONG, nil)
		return
	}

	// 密码正确发放token给前端ReleseToken
	token, errtoken := jwtauth.ReleaseToken(admin)
	if errtoken != nil {
		response.Fail(ctx, errmsg.ERROR_TOKEN_GENERATE, nil)
		return
	}

	if err := service.NewAdminService(db).UpdateLastLoinTime(admin.Id); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_UPDATE_DATA, nil, err.Error())
		return
	}

	response.Success(ctx, gin.H{"token": token}, "ok")
}

func (c Admin) Info(ctx *gin.Context) {
	admin, _ := ctx.Get("admin")
	adminInfo := admin.(model.Admin)

	response.Success(ctx, gin.H{"admin_id": adminInfo.Id, "admin_name": adminInfo.AdminName, "last_login_time": adminInfo.LastLoginTime.Unix()}, "ok")
}

func (c Admin) ChangePassword(ctx *gin.Context) {

	var req = schema.AdminChangePasswordRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, "新密码确认两次密码不一致！")
		return
	}

	admin, _ := ctx.Get("admin")
	adminInfo := admin.(model.Admin)

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(adminInfo.Password), []byte(req.OldPassword)); err != nil {
		response.Fail(ctx, errmsg.ERROR_PASSWORD_WRONG, nil)
		return
	}

	// 加密用户密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(ctx, errmsg.ERROR_PASSWORD_ENCRYPTION, nil)
		return
	}

	if err := service.NewAdminService(orm.GetDBWithContext(ctx.Request.Context())).ChangePassword(adminInfo.Id, string(hashedPassword)); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_UPDATE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewAdmin() Admin {
	return Admin{}
}
