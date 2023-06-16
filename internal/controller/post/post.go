package post

import (
	"admin-server/internal/response"
	"admin-server/internal/schema"
	"admin-server/internal/service"
	"admin-server/pkg/errmsg"
	"admin-server/pkg/orm"
	"admin-server/pkg/orm/datatypes"
	"admin-server/pkg/validator"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Post struct {
}

func (c Post) Info(ctx *gin.Context) {

	var req = schema.PostIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).GetPost("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Post) List(ctx *gin.Context) {

	var req = schema.PostListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).GetPostList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Post) Add(ctx *gin.Context) {

	var req = schema.PostAddRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostAdd(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) UpdateAd(ctx *gin.Context) {

	var req = schema.PostUpdateAdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	images, err := json.Marshal(req.Images)
	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.MARSHAL_JSON_FAIL, nil, "修改帖子图片数组解析失败："+err.Error())
		return
	}

	err = service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostUpdate([]int64{req.Id}, map[string]interface{}{
		"user_id":         req.UserId,
		"title":           req.Title,
		"expert_analysis": req.ExpertAnalysis,
		"is_top":          req.IsTop,
		"images":          datatypes.JSON(images),
		"is_essence_post": req.IsEssencePost,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) Update(ctx *gin.Context) {

	var req = schema.PostUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	playRules, err := json.Marshal(req.PlayRules)
	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.MARSHAL_JSON_FAIL, nil, "修改帖子玩法规则模版解析规则数组失败："+err.Error())
		return
	}

	images, err := json.Marshal(req.Images)
	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.MARSHAL_JSON_FAIL, nil, "修改帖子图片数组解析失败："+err.Error())
		return
	}

	err = service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostUpdate([]int64{req.Id}, map[string]interface{}{
		"title":           req.Title,
		"expert_analysis": req.ExpertAnalysis,
		"points":          req.Points,
		"readers_num":     req.ReadersNum,
		"comments_num":    req.CommentsNum,
		"likes_num":       req.LikesNum,
		"is_essence_post": req.IsEssencePost,
		"play_rules":      datatypes.JSON(playRules),
		"images":          datatypes.JSON(images),
		"is_guaranteed":   req.IsGuaranteed,
		"is_top":          req.IsTop,
		"status":          req.Status,
		"remark":          req.Remark,
		"created_at": datatypes.XTime{
			time.Unix(req.CreatedAt, 0),
		},
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) Audit(ctx *gin.Context) {

	var req = schema.PostAuditRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostUpdate(req.Ids, map[string]interface{}{
		"status": req.Status,
		"remark": req.Remark,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) Result(ctx *gin.Context) {

	var req = schema.PostResultRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostResult(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) Essence(ctx *gin.Context) {

	var req = schema.PostEssenceRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostUpdate(req.Ids, map[string]interface{}{
		"is_essence_post": req.IsEssencePost,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) Top(ctx *gin.Context) {

	var req = schema.PostTopRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostUpdate(req.Ids, map[string]interface{}{
		"is_top": req.IsTop,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Post) Log(ctx *gin.Context) {

	var req = schema.PostLogListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).GetPostLogList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Post) Del(ctx *gin.Context) {

	var req = schema.PostIdsRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPostService(orm.GetDBWithContext(ctx.Request.Context())).PostDel(req.Ids)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewPost() Post {
	return Post{}
}
