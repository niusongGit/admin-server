package routers

import (
	"admin-server/internal/controller/admin"
	"admin-server/internal/controller/announcement"
	"admin-server/internal/controller/banner"
	"admin-server/internal/controller/blogroll"
	"admin-server/internal/controller/comment"
	"admin-server/internal/controller/competition"
	"admin-server/internal/controller/competitionlinks"
	"admin-server/internal/controller/expert"
	"admin-server/internal/controller/feedback"
	"admin-server/internal/controller/payment"
	"admin-server/internal/controller/permission"
	"admin-server/internal/controller/post"
	"admin-server/internal/controller/sensitiveword"
	"admin-server/internal/controller/sport_type"
	"admin-server/internal/controller/system"
	"admin-server/internal/controller/user"
	"admin-server/internal/controller/usermember"
	"admin-server/internal/controller/version"
	"admin-server/internal/middleware"
	"admin-server/internal/qcsocket"
	"admin-server/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	group *gin.RouterGroup
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	//r.StaticFS("/static", http.Dir("./static"))

	group = r.Group("/admin")
	group.Use(middleware.Authmiddleware())

	noAuthGroup := r.Group("/admin")

	admin := admin.NewAdmin()
	noAuthGroup.POST("/add", admin.Add)
	noAuthGroup.POST("/login", admin.Login)
	group.POST("/info", admin.Info)
	group.POST("/changepw", admin.ChangePassword)
	group.POST("/list", admin.List)
	group.POST("/update", admin.Update)
	group.POST("/getrouts", func(ctx *gin.Context) {
		routers := r.Routes()
		res := make([]map[string]string, 0, len(routers))
		for _, v := range routers {
			res = append(res, map[string]string{
				"method": v.Method,
				"path":   v.Path,
			})
		}
		response.Success(ctx, res, "成功")
	})

	sysapi := permission.NewSysApi()
	sysapiGroup := group.Group("/sys_api").Use(middleware.CasbinHandler()) //权限控制
	sysapiGroup.POST("/info", sysapi.Info)
	sysapiGroup.POST("/list", sysapi.List)
	sysapiGroup.POST("/add", sysapi.Add)
	sysapiGroup.POST("/update", sysapi.Update)
	sysapiGroup.POST("/del", sysapi.Del)

	role := permission.NewRole()
	roleGroup := group.Group("/role")
	roleGroup.POST("/info", role.Info)
	roleGroup.POST("/list", role.List)
	roleGroup.POST("/add", role.Add)
	roleGroup.POST("/update", role.Update)
	roleGroup.POST("/del", role.Del)
	roleGroup.POST("/get_role_policy_path", role.GetRolePolicyPath)
	roleGroup.POST("/set_role_policy_path", role.SetRolePolicyPath)

	//group.POST("/upload", middleware.Authmiddleware(), upload.NewUpload().Upload)

	sportType := sport_type.NewSportType()
	sportTypeGroup := group.Group("/sport_type")
	sportTypeGroup.POST("/info", sportType.Info)
	sportTypeGroup.POST("/list", sportType.List)
	sportTypeGroup.POST("/add", sportType.Add)
	sportTypeGroup.POST("/update", sportType.Update)

	banner := banner.NewBanner()
	bannerGroup := group.Group("/banner")
	bannerGroup.POST("/info", banner.Info)
	bannerGroup.POST("/list", banner.List)
	bannerGroup.POST("/add", banner.Add)
	bannerGroup.POST("/update", banner.Update)
	bannerGroup.POST("/del", banner.Del)

	blogroll := blogroll.NewBlogroll()
	blogrollGroup := group.Group("/blogroll")
	blogrollGroup.POST("/info", blogroll.Info)
	blogrollGroup.POST("/list", blogroll.List)
	blogrollGroup.POST("/add", blogroll.Add)
	blogrollGroup.POST("/update", blogroll.Update)
	blogrollGroup.POST("/del", blogroll.Del)

	announcementObj := announcement.NewAnnouncement()
	announcementGroup := group.Group("/announcement")
	announcementGroup.POST("/info", announcementObj.Info)
	announcementGroup.POST("/list", announcementObj.List)
	announcementGroup.POST("/add", announcementObj.Add)
	announcementGroup.POST("/update", announcementObj.Update)
	announcementGroup.POST("/del", announcementObj.Del)

	competitionTypeObj := competition.NewCompetitionType()
	competitionTypeGroup := group.Group("/competition_type")
	competitionTypeGroup.POST("/info", competitionTypeObj.Info)
	competitionTypeGroup.POST("/list", competitionTypeObj.List)
	competitionTypeGroup.POST("/add", competitionTypeObj.Add)
	competitionTypeGroup.POST("/update", competitionTypeObj.Update)
	competitionTypeGroup.POST("/del", competitionTypeObj.Del)

	competitionObj := competition.NewCompetition()
	competitionGroup := group.Group("/competition")
	competitionGroup.POST("/info", competitionObj.Info)
	competitionGroup.POST("/list", competitionObj.List)
	competitionGroup.POST("/add", competitionObj.Add)
	competitionGroup.POST("/update", competitionObj.Update)
	competitionGroup.POST("/del", competitionObj.Del)
	competitionGroup.POST("/status", competitionObj.StatusUpdate)
	competitionGroup.POST("/finish", competitionObj.Finish)

	playRuleTemplate := competition.NewPlayRuleTemplate()
	playRuleTempGroup := group.Group("/play_rule_template")
	playRuleTempGroup.POST("/info", playRuleTemplate.Info)
	playRuleTempGroup.POST("/list", playRuleTemplate.List)
	playRuleTempGroup.POST("/add", playRuleTemplate.Add)
	playRuleTempGroup.POST("/update", playRuleTemplate.Update)
	playRuleTempGroup.POST("/del", playRuleTemplate.Del)
	playRuleTempGroup.POST("/status", playRuleTemplate.StatusUpdate)

	postObj := post.NewPost()
	postGroup := group.Group("/post")
	postGroup.POST("/info", postObj.Info)
	postGroup.POST("/list", postObj.List)
	postGroup.POST("/update", postObj.Update)
	postGroup.POST("/del", postObj.Del)
	postGroup.POST("/audit", postObj.Audit)
	postGroup.POST("/result", postObj.Result)
	postGroup.POST("/essence", postObj.Essence)
	postGroup.POST("/log", postObj.Log)
	postGroup.POST("/top", postObj.Top)
	postGroup.POST("/add", postObj.Add)
	postGroup.POST("/updatead", postObj.UpdateAd)

	versionObj := version.NewVersion()
	versionGroup := group.Group("/version")
	versionGroup.POST("/info", versionObj.Info)
	versionGroup.POST("/list", versionObj.List)
	versionGroup.POST("/add", versionObj.Add)
	versionGroup.POST("/update", versionObj.Update)
	versionGroup.POST("/del", versionObj.Del)

	paymentWayObj := payment.NewPaymentWay()
	paymentWayGroup := group.Group("/payment_way")
	paymentWayGroup.POST("/info", paymentWayObj.Info)
	paymentWayGroup.POST("/list", paymentWayObj.List)
	paymentWayGroup.POST("/add", paymentWayObj.Add)
	paymentWayGroup.POST("/update", paymentWayObj.Update)
	paymentWayGroup.POST("/del", paymentWayObj.Del)

	orderObj := payment.NewOrder()
	orderGroup := group.Group("/order")
	orderGroup.POST("/info", orderObj.Info)
	orderGroup.POST("/list", orderObj.List)
	orderGroup.POST("/status", orderObj.Status)
	orderGroup.POST("/sync", orderObj.Sync)

	withdrawAccountObj := payment.NewWithdrawAccount()
	withdrawAccountGroup := group.Group("/withdraw_account")
	withdrawAccountGroup.POST("/info", withdrawAccountObj.Info)
	withdrawAccountGroup.POST("/list", withdrawAccountObj.List)
	withdrawAccountGroup.POST("/update", withdrawAccountObj.Update)

	withdrawApplicationObj := payment.NewWithdrawApplication()
	withdrawApplicationGroup := group.Group("/withdraw_application")
	withdrawApplicationGroup.POST("/info", withdrawApplicationObj.Info)
	withdrawApplicationGroup.POST("/list", withdrawApplicationObj.List)
	withdrawApplicationGroup.POST("/audit", withdrawApplicationObj.Audit)

	pointRecordCtrl := payment.NewPointRecord()
	pointRecordGroup := group.Group("/point_record")
	pointRecordGroup.POST("/info", pointRecordCtrl.Info)
	pointRecordGroup.POST("/list", pointRecordCtrl.List)

	user1 := user.NewUser()
	userGroup := group.Group("/user")
	userGroup.POST("/info", user1.Info)
	userGroup.POST("/list", user1.List)
	userGroup.POST("/disable", user1.Disable)
	userGroup.POST("/update", user1.Update)
	userGroup.POST("/del", user1.Del)
	userGroup.POST("/point_change", user1.PointChange)
	userGroup.POST("/become_admin", user1.BecomeAdmin)
	userGroup.POST("/verfied_code", user1.VerfiedCode)

	expert1 := expert.NewExpert()
	expertGroup := group.Group("/expert")
	expertGroup.POST("/info", expert1.Info)
	expertGroup.POST("/list", expert1.List)
	expertGroup.POST("/update", expert1.Update)

	applicationAudit := expert.NewApplicationAudit()
	expertApplicationAuditGroup := group.Group("/expert_application_audit")
	expertApplicationAuditGroup.POST("/list", applicationAudit.List)
	//expertApplicationAuditGroup.POST("/update", applicationAudit.Update)
	expertApplicationAuditGroup.POST("/approve", applicationAudit.Approve)

	commentCtrl := comment.NewComment()
	commentGroup := group.Group("/comment")
	commentGroup.POST("/info", commentCtrl.Info)
	commentGroup.POST("/list", commentCtrl.List)
	commentGroup.POST("/update", commentCtrl.Update)
	commentGroup.POST("/audit", commentCtrl.Audit)
	commentGroup.POST("/sticky", commentCtrl.Sticky)
	commentGroup.POST("/hot", commentCtrl.Hot)
	commentGroup.POST("/del", commentCtrl.Del)

	usermemberCtrl := usermember.NewUserMember()
	usermemberGroup := group.Group("/user_member")
	usermemberGroup.POST("/info", usermemberCtrl.Info)
	usermemberGroup.POST("/list", usermemberCtrl.List)

	feedbackCtrl := feedback.NewFeedback()
	feedbackGroup := group.Group("/feedback")
	feedbackGroup.POST("/list", feedbackCtrl.List)

	systemObj := system.NewSystem()
	systemGroup := group.Group("/system")
	systemGroup.POST("/info", systemObj.Info)
	systemGroup.POST("/update", systemObj.Update)
	pointAndAmountObj := system.NewPointAndAmount()
	pointAndAmountGroup := group.Group("/pointandamout")
	pointAndAmountGroup.POST("/info", pointAndAmountObj.Info)
	pointAndAmountGroup.POST("/update", pointAndAmountObj.Update)
	smsTemplateObj := system.NewSmsTemplate()
	smsTemplateGroup := group.Group("/smstemplate")
	smsTemplateGroup.POST("/info", smsTemplateObj.Info)
	smsTemplateGroup.POST("/update", smsTemplateObj.Update)

	memberCategoriesObj := system.NewMemberCategories()
	memberCategoriesGroup := group.Group("/membercategories")
	memberCategoriesGroup.POST("/info", memberCategoriesObj.Info)
	memberCategoriesGroup.POST("/update", memberCategoriesObj.Update)

	sensitiveWordObj := sensitiveword.NewSensitiveWord()
	sensitiveWordGroup := group.Group("/sensitive_word")
	sensitiveWordGroup.POST("/info", sensitiveWordObj.Info)
	sensitiveWordGroup.POST("/list", sensitiveWordObj.List)
	sensitiveWordGroup.POST("/add", sensitiveWordObj.Add)
	sensitiveWordGroup.POST("/update", sensitiveWordObj.Update)
	sensitiveWordGroup.POST("/del", sensitiveWordObj.Del)

	clCtrl := competitionlinks.NewCompetitionLinks()
	competitionLinksGroup := group.Group("/competition_links")
	competitionLinksGroup.POST("/info", clCtrl.Info)
	competitionLinksGroup.POST("/list", clCtrl.List)
	competitionLinksGroup.POST("/add", clCtrl.Add)
	competitionLinksGroup.POST("/update", clCtrl.Update)
	competitionLinksGroup.POST("/del", clCtrl.Del)

	noAuthGroup.GET("/websocket", func(c *gin.Context) {
		var upGrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		//升级get请求为webSocket协议
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer ws.Close()
		qcsocket.Ws(c.Request.Context(), ws)
	})

	return r
}
