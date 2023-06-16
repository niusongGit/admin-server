package competition

import (
	"admin-server/internal/consts"
	"admin-server/internal/response"
	"admin-server/internal/schema"
	"admin-server/internal/service"
	"admin-server/pkg/errmsg"
	"admin-server/pkg/orm"
	"admin-server/pkg/orm/datatypes"
	"admin-server/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Competition struct {
}

func (c Competition) Info(ctx *gin.Context) {

	var req = schema.CompetitionIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).GetCompetition("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Competition) List(ctx *gin.Context) {

	var req = schema.CompetitionListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).GetCompetitionList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Competition) Add(ctx *gin.Context) {

	var req = schema.CompetitionAddRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).CompetitionAdd(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Competition) Update(ctx *gin.Context) {

	var req = schema.CompetitionUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	olddata, err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).GetCompetition("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	data := map[string]interface{}{
		"sport_type_id":       req.SportTypeId,
		"competition_type_id": req.CompetitionTypeId,
		"start_time": datatypes.XTime{
			time.Unix(req.StartTime, 0),
		},
		"end_time":      datatypes.XTime{time.Unix(req.EndTime, 0)},
		"title":         req.Title,
		"template_code": req.TemplateCode,
		"template":      req.Template,
	}

	if olddata.Status != consts.CompetitionStatusEnd {
		data["status"] = req.Status
	}

	err = service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).CompetitionUpdate([]int64{req.Id}, data)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Competition) StatusUpdate(ctx *gin.Context) {

	var req = schema.CompetitionStatusUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).CompetitionUpdate(req.Ids, map[string]interface{}{
		"status": req.Status,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Competition) Finish(ctx *gin.Context) {

	var req = schema.CompetitionFinishRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).CompetitionFinish(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Competition) Del(ctx *gin.Context) {

	var req = schema.CompetitionIdsRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewCompetitionService(orm.GetDBWithContext(ctx.Request.Context())).CompetitionDel(req.Ids)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewCompetition() Competition {
	return Competition{}
}
