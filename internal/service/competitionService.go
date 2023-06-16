package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/orm/datatypes"
	"admin-server/pkg/utils"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

type CompetitionService struct {
	db *gorm.DB
}

func (s *CompetitionService) GetCompetition(query interface{}, args ...interface{}) (*schema.CompetitionSportTypeCompetitionType, error) {

	m := schema.CompetitionSportTypeCompetitionType{}

	if err := s.db.Where(query, args...).Preload("SportType").Preload("CompetitionType").First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *CompetitionService) GetCompetitionList(params schema.CompetitionListRequest) (*schema.CompetitionListResponse, error) {

	dbModel := s.db.Model(&model.Competition{})

	if params.SportTypeId > 0 {
		dbModel = dbModel.Where("sport_type_id = ?", params.SportTypeId)
	}

	if params.CompetitionTypeId > 0 {
		dbModel = dbModel.Where("competition_type_id = ?", params.CompetitionTypeId)
	}

	if params.Title != "" {
		dbModel = dbModel.Where("title LIKE ?", "%"+params.Title+"%")
	}

	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.CompetitionSportTypeCompetitionType

	err := dbModel.Preload("SportType").Preload("CompetitionType").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.CompetitionListResponse{total, params.Page, len(list), list}, nil
}

func (s *CompetitionService) CompetitionAdd(params schema.CompetitionAddRequest) error {
	m := model.Competition{
		SportTypeId:       params.SportTypeId,
		CompetitionTypeId: params.CompetitionTypeId,
		StartTime: datatypes.XTime{
			time.Unix(params.StartTime, 0),
		},
		EndTime: datatypes.XTime{
			time.Unix(params.EndTime, 0),
		},
		Title:        params.Title,
		TemplateCode: params.TemplateCode,
		Template:     params.Template,
		Status:       params.Status,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *CompetitionService) CompetitionUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.Competition{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *CompetitionService) CompetitionFinish(params schema.CompetitionFinishRequest) error {

	data := map[string]interface{}{
		"status":             consts.CompetitionStatusEnd,
		"practical_end_time": datatypes.XTime{time.Now()},
	}

	if len(params.PlayRules) > 0 {
		playRulesByte, err := json.Marshal(params.PlayRules)
		if err != nil {
			return errors.New("结束比赛规则数组json序列化失败：" + err.Error())
		}
		data["finish_play_rules"] = datatypes.JSON(playRulesByte)
	}

	if len(params.Template) > 0 {
		data["template"] = params.Template
	}

	res := s.db.Model(&model.Competition{}).Where("id = ? and status <> ?", params.Id, consts.CompetitionStatusEnd).Updates(data)

	if res.RowsAffected == 0 || res.Error != nil {
		return errors.New("更新结束比赛失败")
	}

	return nil
}

func (s *CompetitionService) CompetitionDel(ids []int64) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		var postIds []int64
		tx.Model(&model.Post{}).Where("competition_id in ?", ids).Pluck("id", &postIds)

		if len(postIds) > 0 {
			if err := NewPostService(tx).PostDel(postIds); err != nil {
				return err
			}
		}

		if err := tx.Delete(&model.Competition{}, ids).Error; err != nil {
			return errors.New("删除比赛异常：" + err.Error())
		}

		return nil
	})
}

func NewCompetitionService(db *gorm.DB) *CompetitionService {
	return &CompetitionService{
		db: db,
	}
}
