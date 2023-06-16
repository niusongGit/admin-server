package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func (s *PostService) GetPost(query interface{}, args ...interface{}) (*schema.PostCompetitionUser, error) {

	m := schema.PostCompetitionUser{}

	if err := s.db.Model(model.Post{}).Preload("Competition.SportType").Preload("Competition.CompetitionType").Preload("User").Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *PostService) GetPostList(params schema.PostListRequest) (*schema.PostListResponse, error) {

	dbModel := s.db.Model(&model.Post{})

	if params.IsAdPost > 0 {
		dbModel = dbModel.Where("competition_id = ?", 0)
	} else {
		dbModel = dbModel.Where("competition_id > ?", 0)

		competitionWhere := make(map[string]interface{})

		if params.SportTypeId > 0 {
			competitionWhere["sport_type_id"] = params.SportTypeId
		}

		if params.CompetitionTypeId != nil {
			competitionWhere["competition_type_id"] = *params.CompetitionTypeId
		}

		if params.CompetitionTitle != "" {
			competitionWhere["title"] = params.CompetitionTitle
		}

		if params.HomeTeam != "" {
			competitionWhere["home_team"] = params.HomeTeam
		}

		if params.AwayTeam != "" {
			competitionWhere["away_team"] = params.AwayTeam
		}

		if len(competitionWhere) > 0 {
			competitionIds := make([]int64, 0)
			s.db.Model(&model.Competition{}).Where(competitionWhere).Pluck("id", &competitionIds)
			dbModel = dbModel.Where("competition_id in ?", competitionIds)
		}
	}

	userWhere := make(map[string]interface{})
	if params.UserName != "" {
		userWhere["name"] = params.UserName
	}

	if params.Phone != "" {
		userWhere["phone"] = params.Phone
	}

	if len(userWhere) > 0 {
		userIds := make([]int64, 0)
		s.db.Model(&model.User{}).Where(userWhere).Pluck("id", &userIds)
		dbModel = dbModel.Where("user_id in ?", userIds)
	}

	if params.Title != "" {
		dbModel = dbModel.Where("title LIKE ?", "%"+params.Title+"%")
	}

	if params.PredictedResult != "" {
		dbModel = dbModel.Where("predicted_result = ?", params.PredictedResult)
	}

	if len(params.IsGuaranteed) > 0 {
		dbModel = dbModel.Where("is_guaranteed in ?", params.IsGuaranteed)
	}

	if len(params.IsTop) > 0 {
		dbModel = dbModel.Where("is_top in ?", params.IsTop)
	}

	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.PostCompetitionUser

	err := dbModel.Preload("Competition.SportType").Preload("Competition.CompetitionType").Preload("User").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.PostListResponse{total, params.Page, len(list), list}, nil
}

func (s *PostService) PostAdd(params schema.PostAddRequest) error {
	images, err := json.Marshal(params.Images)
	if err != nil {
		return errors.New("添加帖子图片数组解析失败：" + err.Error())
	}
	m := model.Post{
		UserId:         params.UserId,
		Title:          params.Title,
		ExpertAnalysis: params.ExpertAnalysis,
		IsEssencePost:  params.IsEssencePost,
		IsTop:          params.IsTop,
		Images:         images,
		PlayRules:      []byte("{}"),
		Status:         consts.AuditPass.Int(),
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *PostService) PostUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.Post{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *PostService) PostResult(params schema.PostResultRequest) error {

	return s.db.Transaction(func(tx *gorm.DB) error {
		var post model.Post

		if err := tx.Where("id = ? and predicted_result = ? and status = ?", params.Id, consts.PredictedResultUnknown, consts.AuditPass).First(&post).Error; err != nil {
			return errors.New("只能操作审核通过并未公部结果的帖子！")
		}

		var competition model.Competition

		if err := tx.Where("id = ? and status = ?", post.CompetitionId, consts.CompetitionStatusEnd).First(&competition).Error; err != nil {
			return errors.New("帖子所属比赛未结束！")
		}

		res := tx.Model(&model.Post{}).Where("id = ?", params.Id).Update("predicted_result", params.PredictedResult)
		if res.RowsAffected == 0 || res.Error != nil {
			return errors.New("更新贴子结果失败")
		}

		var list []*model.PostLog

		if err := tx.Model(&model.PostLog{}).Where("status = ? and operation_type = ? and post_id = ?", 0, "purchase", params.Id).Find(&list).Error; err != nil {
			return errors.New("查询帖子日志失败：" + err.Error())
		}

		system, err := NewSystemService(s.db).GetSystem()
		if err != nil {
			return errors.New("获取系统配置失败：" + err.Error())
		}

		for _, v := range list {
			if v.IsGuaranteed == 1 { //担保
				point := post.Points * float64(system.GuaranteedPointMultiple)
				amount := point / float64(system.AmountToPoint)
				if params.PredictedResult == consts.PredictedResultRed { //已中
					// 给专家发两倍积分兑换的金额
					if err := tx.Model(&model.User{}).Where("id =?", post.UserId).Update("withdraw_amount", gorm.Expr("withdraw_amount + ?", amount)).Error; err != nil {
						return errors.New("给专家发金额失败：" + err.Error())
					}
					// 用户冻结的积分扣除2倍
					if err := tx.Model(&model.User{}).Where("id =?", v.UserId).Update("frozen_points", gorm.Expr("frozen_points - ?", amount)).Error; err != nil {
						return errors.New("用户冻结积分扣除失败：" + err.Error())
					}
				} else {
					// 用户冻结的积分扣除2倍,积分加2倍
					if err := tx.Model(&model.User{}).Where("id =?", v.UserId).Updates(map[string]interface{}{
						"frozen_points": gorm.Expr("frozen_points - ?", point),
						"points":        gorm.Expr("points + ?", point),
					}).Error; err != nil {
						return errors.New("用户冻结积分扣除失败：" + err.Error())
					}
				}
			} else {
				// 给专家发单倍钱
				if err := tx.Model(&model.User{}).Where("id =?", post.UserId).Update("withdraw_amount", gorm.Expr("withdraw_amount + ?", post.Points/float64(system.AmountToPoint))).Error; err != nil {
					return errors.New("给专家发金额失败：" + err.Error())
				}
				// 给用户冻结的积分扣除1倍
				if err := tx.Model(&model.User{}).Where("id =?", v.UserId).Update("frozen_points", gorm.Expr("frozen_points - ?", post.Points)).Error; err != nil {
					return errors.New("用户冻结积分扣除失败：" + err.Error())
				}

			}
			// 标记记录已处理过
			if err := tx.Model(&model.PostLog{}).Where("id =?", v.Id).Update("status", 1).Error; err != nil {
				return errors.New("更新标记记录失败：" + err.Error())
			}
		}

		return nil
	})
}

func (s *PostService) GetPostLogList(params schema.PostLogListRequest) (*schema.PostLogListResponse, error) {

	dbModel := s.db.Model(&model.PostLog{})

	if params.PostId > 0 {
		dbModel = dbModel.Where("post_id = ?", params.PostId)
	}

	if params.OperationType != "" {
		dbModel = dbModel.Where("operation_type = ?", params.OperationType)
	}

	userWhere := make(map[string]interface{})
	if params.UserName != "" {
		userWhere["name"] = params.UserName
	}

	if params.Phone != "" {
		userWhere["phone"] = params.Phone
	}

	if len(userWhere) > 0 {
		userIds := make([]int64, 0)
		s.db.Model(&model.User{}).Where(userWhere).Pluck("id", &userIds)
		dbModel = dbModel.Where("user_id in ?", userIds)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.PostLogUser

	err := dbModel.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.PostLogListResponse{total, params.Page, len(list), list}, nil
}

func (s *PostService) PostDel(ids []int64) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		var commentIds []int64
		tx.Model(&model.Comment{}).Where("post_id in ?", ids).Pluck("id", &commentIds)

		if len(commentIds) > 0 {
			if err := NewCommentService(tx).CommentDel(commentIds); err != nil {
				return err
			}
		}

		if err := tx.Where("post_id in ?", ids).Delete(&model.PostLog{}).Error; err != nil {
			return errors.New("删除贴子日志异常：" + err.Error())
		}

		if err := tx.Delete(&model.Post{}, ids).Error; err != nil {
			return errors.New("删除贴子异常：" + err.Error())
		}
		return nil
	})

}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		db: db,
	}
}
