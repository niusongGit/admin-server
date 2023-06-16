package competition

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
)

type PlayRuleTemplate struct {
}

func (c PlayRuleTemplate) Info(ctx *gin.Context) {

	var req = schema.PlayRuleTemplateIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPlayRuleTemplateService(orm.GetDBWithContext(ctx.Request.Context())).GetPlayRuleTemplate("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c PlayRuleTemplate) List(ctx *gin.Context) {

	var req = schema.PlayRuleTemplateListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPlayRuleTemplateService(orm.GetDBWithContext(ctx.Request.Context())).GetPlayRuleTemplateList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c PlayRuleTemplate) Add(ctx *gin.Context) {

	var req = schema.PlayRuleTemplateAddRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}
	ser := service.NewPlayRuleTemplateService(orm.GetDBWithContext(ctx.Request.Context()))

	_, err := ser.GetPlayRuleTemplate("sport_type_id = ? and name = ? and code = ? and type = ?", req.SportTypeId, req.Name, req.Code, req.Type)
	if err == nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROR_DATA_DUPLICATION, nil, "该规则已重复")
		return
	}

	err = ser.PlayRuleTemplateAdd(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c PlayRuleTemplate) Update(ctx *gin.Context) {

	var req = schema.PlayRuleTemplateUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	ser := service.NewPlayRuleTemplateService(orm.GetDBWithContext(ctx.Request.Context()))
	_, err := ser.GetPlayRuleTemplate("sport_type_id = ? and name = ? and code = ? and type = ? and id <> ?", req.SportTypeId, req.Name, req.Code, req.Type, req.Id)
	if err == nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROR_DATA_DUPLICATION, nil, "该规则已重复")
		return
	}

	choicesByte, err := json.Marshal(req.Choices)
	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.MARSHAL_JSON_FAIL, nil, "添加玩法规则模版解析规则数组失败："+err.Error())
		return
	}

	err = ser.PlayRuleTemplateUpdate([]int64{req.Id}, map[string]interface{}{
		"sport_type_id":         req.SportTypeId,
		"name":                  req.Name,
		"code":                  req.Code,
		"type":                  req.Type,
		"choices":               datatypes.JSON(choicesByte),
		"post_content_template": req.PostContentTemplate,
		"status":                req.Status,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c PlayRuleTemplate) StatusUpdate(ctx *gin.Context) {

	var req = schema.PlayRuleTemplateStatusUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPlayRuleTemplateService(orm.GetDBWithContext(ctx.Request.Context())).PlayRuleTemplateUpdate(req.Ids, map[string]interface{}{
		"status": req.Status,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c PlayRuleTemplate) Del(ctx *gin.Context) {

	var req = schema.PlayRuleTemplateIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPlayRuleTemplateService(orm.GetDBWithContext(ctx.Request.Context())).PlayRuleTemplateDel(req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewPlayRuleTemplate() PlayRuleTemplate {
	return PlayRuleTemplate{}
}
