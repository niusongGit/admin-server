package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func (s *CommentService) GetComment(query interface{}, args ...interface{}) (*schema.CommentItem, error) {

	m := schema.CommentItem{}
	if err := s.db.Model(model.Comment{}).
		Preload("Post").
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Preload("ReplyToUser", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *CommentService) GetCommentList(params schema.CommentListRequest) (*schema.CommentListResponse, error) {

	dbModel := s.db.Model(&model.Comment{})

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

	if params.Content != "" {
		dbModel = dbModel.Where("content LIKE ?", "%"+params.Content+"%")
	}

	if params.Level != 0 {
		dbModel = dbModel.Where("level = ?", params.Level)
	}
	if params.IsHot != nil {
		dbModel = dbModel.Where("is_hot = ?", &params.IsHot)
	}
	if params.IsSticky != nil {
		dbModel = dbModel.Where("is_sticky = ?", &params.IsSticky)
	}

	if len(params.Statuses) > 0 {
		dbModel = dbModel.Where("status in ?", params.Statuses)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.CommentItem
	err := dbModel.Model(model.Comment{}).
		Preload("Post").
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Preload("ReplyToUser", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error

	if err != nil {
		return nil, err
	}

	return &schema.CommentListResponse{total, params.Page, len(list), list}, nil
}

func (s *CommentService) CommentUpdate(ids []int64, data model.Comment) error {

	if err := s.db.Model(&model.Comment{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *CommentService) CommentSticky(ids []int64, isSticky int) error {
	if err := s.db.Model(&model.Comment{}).Where("id in ?", ids).UpdateColumn("is_sticky", isSticky).Error; err != nil {
		return err
	}
	return nil
}

func (s *CommentService) CommentHot(ids []int64, isHot int) error {
	if err := s.db.Model(&model.Comment{}).Where("id in ?", ids).UpdateColumn("is_hot", isHot).Error; err != nil {
		return err
	}
	return nil
}

func (s *CommentService) CommentDel(ids []int64) error {

	return s.db.Transaction(func(tx *gorm.DB) error {
		var commentIds []int64
		tx.Model(&model.Comment{}).Where("id in @cid or parent_id in @cid", map[string]interface{}{"cid": ids}).Pluck("id", &commentIds)

		if err := tx.Where("comment_id in ?", commentIds).Delete(&model.CommentLog{}).Error; err != nil {
			return errors.New("删除评论日志异常：" + err.Error())
		}

		if err := tx.Delete(&model.Comment{}, commentIds).Error; err != nil {
			return errors.New("删除评论异常：" + err.Error())
		}

		return nil
	})

}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{
		db: db,
	}
}
