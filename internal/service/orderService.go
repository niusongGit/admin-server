package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/goredis"
	"admin-server/pkg/utils"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func (s *OrderService) GetOrder(query interface{}, args ...interface{}) (*schema.OrderUser, error) {

	m := schema.OrderUser{}

	if err := s.db.Model(model.Order{}).Preload("User").Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *OrderService) GetOrderList(params schema.OrderListRequest) (*schema.OrderListResponse, error) {

	dbModel := s.db.Model(&model.Order{})

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

	if params.OrderNo != "" {
		dbModel = dbModel.Where("order_no = ?", params.OrderNo)
	}

	if params.ThirdPartyOrderNo != "" {
		dbModel = dbModel.Where("third_party_order_no = ?", params.ThirdPartyOrderNo)
	}

	if params.TopUpType != "" {
		dbModel = dbModel.Where("top_up_type = ?", params.TopUpType)
	}

	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.OrderUser

	err := dbModel.Preload("User").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.OrderListResponse{total, params.Page, len(list), list}, nil
}

func (s *OrderService) OrderUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.Order{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *OrderService) SyncOrder(adminId int64, query interface{}, args ...interface{}) error {

	m := model.Order{}
	if err := s.db.Model(model.Order{}).Where(query, args...).First(&m).Error; err != nil {
		return err
	}
	if m.Status != 0 {
		return errors.New("此订单不需要同步")
	}
	val := fmt.Sprintf(consts.RedisListValueForOrder, m.OrderNo, fmt.Sprintf(consts.StrAdmin, adminId))
	fmt.Println(val)
	err := goredis.GetRedisDB().LPush(context.Background(), consts.RedisKeyForOrder, val).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}
