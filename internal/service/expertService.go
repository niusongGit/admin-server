package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type ExpertService struct {
	db *gorm.DB
}

func (s *ExpertService) GetExpert(query interface{}, args ...interface{}) (*schema.ExpertItem, error) {
	m := schema.ExpertItem{}

	if err := s.db.Model(model.User{}).
		Preload("SportType", func(db *gorm.DB) *gorm.DB { return db.Model(model.SportType{}) }).
		Where(query, args...).
		First(&m, "is_expert = 1").Error; err != nil {
		return nil, err
	}

	predictedRatioM, err := s.GetRecentPredictedRatio([]int64{m.Id})
	if err != nil {
		return nil, err
	}
	m.RecentHitRatio = predictedRatioM[m.Id]
	return &m, nil
}

func (s *ExpertService) GetExpertList(params schema.ExpertListRequest) (*schema.ExpertListResponse, error) {

	dbModel := s.db.Model(&model.User{}).
		Preload("SportType", func(db *gorm.DB) *gorm.DB { return db.Model(model.SportType{}) }).
		Where("is_expert = 1")
	if params.Id > 0 {
		dbModel = dbModel.Where("id = ?", params.Id)
	}
	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if len(params.Phone) > 0 {
		dbModel = dbModel.Where("phone = ?", params.Phone)
	}
	if len(params.Statuses) > 0 {
		dbModel = dbModel.Where("status in ?", params.Statuses)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.ExpertItem

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).
		Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return &schema.ExpertListResponse{total, params.Page, len(list), list}, nil
	}

	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	predictedRatioM, err := s.GetRecentPredictedRatio(ids)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		v.RecentHitRatio = predictedRatioM[v.Id]
	}

	return &schema.ExpertListResponse{total, params.Page, len(list), list}, nil
}

func (s *ExpertService) ExpertUpdate(params schema.ExpertUpdateRequest) error {

	if err := s.db.Model(model.User{}).Where("id = ?", params.Id).Updates(&model.Expert{
		Name:        params.Name,
		SportTypeId: params.SportTypeId,
	}).Error; err != nil {
		return err
	}

	return nil
}

func NewExpertService(db *gorm.DB) *ExpertService {
	return &ExpertService{
		db: db,
	}
}

type predictedData struct {
	Uid   int64
	Red   int
	Total int
}

func (s *ExpertService) GetRecentPredictedRatio(uids []int64) (m map[int64]float32, err error) {
	systemConf, err := NewSystemService(s.db).GetSystem()
	if err != nil {
		return nil, err
	}

	m = make(map[int64]float32)
	pm := make(map[int64]*predictedData)

	for _, uid := range uids {
		rows, err := s.db.Raw(`SELECT p.predicted_result  FROM post p JOIN competition c ON p.competition_id = c.id 
                           WHERE p.user_id = ? AND c.status = ? ORDER BY created_at DESC LIMIT ?`,
			uid, consts.CompetitionStatusEnd, systemConf.RecentCompetitionCount).Rows()
		if err != nil {
			return nil, err
		}
		pd := &predictedData{
			Uid: uid,
		}
		pm[uid] = pd
		for rows.Next() {
			predictedResult := ""
			rows.Scan(&predictedResult)
			pd.Total++
			if predictedResult == consts.PredictedResultRed {
				pd.Red++
			}
		}

	}

	for uid, v := range pm {
		if v.Red == 0 {
			continue
		}
		m[uid] = float32(v.Red) / float32(v.Total)
	}
	return m, nil
}
