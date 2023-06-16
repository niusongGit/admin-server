package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/goredis"
	"admin-server/pkg/utils"
	"context"
	"gorm.io/gorm"
)

type SensitiveWordService struct {
	db *gorm.DB
}

func (s *SensitiveWordService) GetSensitiveWord(query interface{}, args ...interface{}) (*model.SensitiveWord, error) {

	m := model.SensitiveWord{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *SensitiveWordService) GetSensitiveWordList(params schema.SensitiveWordListRequest) (*schema.SensitiveWordListResponse, error) {

	dbModel := s.db.Model(&model.SensitiveWord{})

	if params.Word != "" {
		dbModel = dbModel.Where("word LIKE ?", "%"+params.Word+"%")
	}
	if len(params.Statuses) > 0 {
		dbModel = dbModel.Where("status in ?", params.Statuses)
	}
	if params.Type != 0 {
		dbModel = dbModel.Where("type = ?", params.Type)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*model.SensitiveWord

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.SensitiveWordListResponse{total, params.Page, len(list), list}, nil
}

func (s *SensitiveWordService) SensitiveWordAdd(params schema.SensitiveWordAddRequest) error {
	m := model.SensitiveWord{
		Word: params.Word,
		Type: params.Type,
	}
	err := s.db.Create(&m).Error
	if err != nil {
		return err
	}

	err = goredis.GetRedisDB().SAdd(context.Background(), consts.RedisSensitiveWord, params.Word).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *SensitiveWordService) SensitiveWordUpdate(params schema.SensitiveWordUpdateRequest) error {
	oldSw := &model.SensitiveWord{}
	err := s.db.First(oldSw, "id = ?", params.Id).Error
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	if params.Word != "" {
		m["word"] = params.Word
	}
	if params.Type != 0 {
		m["type"] = params.Type
	}

	if params.Status != nil {
		m["status"] = *params.Status
	}
	err = s.db.Model(&model.SensitiveWord{}).Where("id = ?", params.Id).Updates(m).Error
	if err != nil {
		return err
	}

	newSw := &model.SensitiveWord{}
	err = s.db.First(newSw, "id = ?", params.Id).Error
	if err != nil {
		return err
	}
	if params.Status != nil && *params.Status == 1 && newSw.Status == 1 {
		err = goredis.GetRedisDB().SRem(context.Background(), consts.RedisSensitiveWord, newSw.Word).Err()
		if err != nil {
			return err
		}
		return nil
	}
	if params.Status != nil && *params.Status == 0 && newSw.Status == 0 {
		err = goredis.GetRedisDB().SAdd(context.Background(), consts.RedisSensitiveWord, newSw.Word).Err()
		if err != nil {
			return err
		}
		return nil
	}

	if oldSw.Word != newSw.Word {
		err = goredis.GetRedisDB().SRem(context.Background(), consts.RedisSensitiveWord, oldSw.Word).Err()
		if err != nil {
			return err
		}
		err = goredis.GetRedisDB().SAdd(context.Background(), consts.RedisSensitiveWord, newSw.Word).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SensitiveWordService) SensitiveWordDel(id int64) error {
	sw := &model.SensitiveWord{}
	err := s.db.Take(sw, "id = ?", id).Error
	if err != nil {
		return err
	}

	err = s.db.Delete(&model.SensitiveWord{}, id).Error
	if err != nil {
		return err
	}
	err = goredis.GetRedisDB().SRem(context.Background(), consts.RedisSensitiveWord, sw.Word).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewSensitiveWordService(db *gorm.DB) *SensitiveWordService {
	return &SensitiveWordService{db: db}
}
