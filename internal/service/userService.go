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
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) GetUser(query interface{}, args ...interface{}) (*model.User, error) {

	m := model.User{}

	if err := s.db.Where(query, args...).
		First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *UserService) GetUserList(params schema.UserListRequest) (*schema.UserListResponse, error) {

	dbModel := s.db.Model(&model.User{})
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

	if params.IsAdmin != nil {
		dbModel = dbModel.Where("is_admin = ?", *params.IsAdmin)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []model.User

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.UserListResponse{total, params.Page, len(list), &list}, nil
}

func (s *UserService) UserUpdate(params schema.UserUpdateRequest) error {

	if err := s.db.Model(&model.User{}).Where("id = ?", params.Id).Updates(&model.User{
		Name:   params.Name,
		Gender: params.Gender,
		Bio:    params.Bio,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) UserDisable(params schema.UserDisableRequest) error {
	status := consts.UserStatusNormal
	if params.Disable {
		status = consts.UserStatusDisabled
	}

	if err := s.db.Model(&model.User{}).Where("id = ?", params.Id).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) PointChange(params schema.UserPointChangeRequest) error {

	var change string
	var objType string
	if params.Type == "income" {
		change = "+"
		objType = "topup"
	} else if params.Type == "expenditure" {
		change = "-"
		objType = "deduct"
	} else {
		return errors.New("类型错误")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&model.User{}).Where("id = ?", params.Id).Update("points", gorm.Expr("points "+change+" ?", params.Point)).Error; err != nil {
			return err
		}

		objByte, err := json.Marshal(schema.UserPointRecordObject{
			Id:      params.Id,
			Type:    objType,
			AdminId: params.AdminId,
		})
		if err != nil {
			return errors.New("解析结构体失败：" + err.Error())
		}

		if err := tx.Create(&model.PointRecord{
			UserId: params.Id,
			Type:   params.Type,
			Object: datatypes.JSON(objByte),
			Point:  params.Point,
		}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *UserService) UserBecomeAdmin(params schema.UserBecomeAdminRequest) error {

	if err := s.db.Model(&model.User{}).Where("id = ?", params.Id).Update("is_admin", *params.IsAdmin).Error; err != nil {
		return err
	}

	return nil
}

func (s *UserService) UserDel(ids []int64) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		// ----------------------------评论相关

		var commentIds []int64
		tx.Model(&model.Comment{}).Where("user_id in @uid or reply_to_id in @uid", map[string]interface{}{"uid": ids}).Pluck("id", &commentIds)
		if len(commentIds) > 0 {
			if err := NewCommentService(tx).CommentDel(commentIds); err != nil {
				return err
			}
		}

		if err := tx.Delete(&model.CommentLog{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除评论日志异常1：" + err.Error())
		}

		// ----------------贴子相关

		var postIds []int64
		tx.Model(&model.Post{}).Where("user_id in ?", ids).Pluck("id", &postIds)

		if len(postIds) > 0 {
			if err := NewPostService(tx).PostDel(postIds); err != nil {
				return err
			}
		}

		if err := tx.Delete(&model.PostLog{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除贴子日志异常1：" + err.Error())
		}

		// --------------------------专家相关
		if err := tx.Delete(&model.ExpertApplicationAudit{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除专家审核申请异常：" + err.Error())
		}

		// --------------------------反馈
		if err := tx.Delete(&model.Feedback{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除反馈异常：" + err.Error())
		}

		// --------------------------订单
		if err := tx.Delete(&model.Order{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除订单异常：" + err.Error())
		}

		// --------------------------积分记录
		if err := tx.Delete(&model.PointRecord{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除积分记录异常：" + err.Error())
		}

		// --------------------------关注和粉丝相关
		if err := tx.Delete(&model.UserFollowingRef{}, "user_id in @uid or following_uid in @uid", map[string]interface{}{"uid": ids}).Error; err != nil {
			return errors.New("删除用户和关注用户的关系异常：" + err.Error())
		}

		if err := tx.Delete(&model.UserFollowerRef{}, "user_id in @uid or follower_uid in @uid", map[string]interface{}{"uid": ids}).Error; err != nil {
			return errors.New("删除用户粉丝关系异常：" + err.Error())
		}

		// --------------------------会员
		if err := tx.Delete(&model.UserMember{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除会员异常：" + err.Error())
		}

		// --------------------------提现相关
		if err := tx.Delete(&model.WithdrawAccount{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除提现帐户异常：" + err.Error())
		}

		if err := tx.Delete(&model.WithdrawApplication{}, "user_id in ?", ids).Error; err != nil {
			return errors.New("删除提现申请记录异常：" + err.Error())
		}

		if err := tx.Delete(&model.User{}, ids).Error; err != nil {
			return errors.New("删除用户异常：" + err.Error())
		}

		return nil
	})
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}
