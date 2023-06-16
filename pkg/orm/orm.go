package orm

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	apklogger "admin-server/pkg/logger"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func InitDb() {
	//配置文件中变量名：后要有一个空格才能开始写变量的值，否则取不到值
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)

	slowLogger := logger.New(
		//设置Logger
		apklogger.NewMyGormWriter(),
		logger.Config{
			//慢SQL阈值
			SlowThreshold: 200 * time.Millisecond,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	//slowLogger.LogMode(logger.Info)
	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: slowLogger,
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		panic("连接数据库失败，请检查参数：" + err.Error())
	}
	db = db.Set("gorm:table_options", "COLLATE=utf8mb4_general_ci")
	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	err = db.AutoMigrate(
		&model.Admin{}, &model.Announcement{}, &model.Banner{}, &model.Blogroll{},
		&model.Comment{}, &model.CommentLog{},
		&model.Competition{}, &model.CompetitionType{},
		&model.DataDictionary{},
		&model.ExpertApplicationAudit{},
		&model.Feedback{},
		&model.Order{},
		&model.PaymentWay{},
		&model.PlayRuleTemplate{},
		&model.PointRecord{},
		&model.Post{}, &model.PostLog{},
		&model.SportType{},
		&model.User{},
		&model.UserFollowingRef{}, &model.UserFollowerRef{}, &model.UserMember{},
		&model.Version{},
		&model.WithdrawAccount{}, &model.WithdrawApplication{},
		&model.SensitiveWord{},
		&model.CompetitionLinks{},
	)
	if err != nil {
		panic("迁移数据表失败！：" + err.Error())
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	initData()
}

func GetDB() *gorm.DB {
	return db.Debug()
}

func GetDBWithContext(c context.Context) *gorm.DB {
	myDb := db.Debug()
	myDb.Logger = logger.New(
		//设置Logger
		apklogger.NewMyGormWriterWithContext(c),
		logger.Config{
			//慢SQL阈值
			SlowThreshold: 200 * time.Millisecond,
			//设置日志级别，只有Warn以上才会打印sql
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	return myDb
}

func initData() {
	err := GetDB().Transaction(func(tx *gorm.DB) error {

		if err := tx.Where(model.Admin{Id: 1, AdminName: "admin"}).Attrs(model.Admin{Password: "$2a$10$92EgCdFEKsB.v5VU.LAQe.5FkwyfzoszM2xW0BzL5y.O0R.KIxGAK"}).FirstOrCreate(&model.Admin{}).Error; err != nil {
			return err
		}

		if err := tx.Where(model.DataDictionary{DataType: consts.DataDictionaryPropertiesSystem}).Attrs(model.DataDictionary{
			Properties: []byte(`{"amount_to_point": 10, "recent_competition_count": 7, "guaranteed_point_multiple": 2, "near_competition_finish_disable_buy_post_time": 0, "near_competition_finish_disable_send_post_time": 0, "customer_service_telephone": ""}`),
		}).FirstOrCreate(&model.DataDictionary{}).Error; err != nil {
			return err
		}

		if err := tx.Where(model.DataDictionary{DataType: consts.DataDictionaryPointAndAmount}).Attrs(model.DataDictionary{
			Properties: []byte(`[{"point": 1, "amount": 0.1}, {"point": 50, "amount": 5}, {"point": 100, "amount": 9.9}, {"point": 500, "amount": 49.8}, {"point": 1000, "amount": 99}, {"point": 2000, "amount": 198}, {"point": 5000, "amount": 490}]`),
		}).FirstOrCreate(&model.DataDictionary{}).Error; err != nil {
			return err
		}

		if err := tx.Where(model.DataDictionary{DataType: consts.DataDictionarySmsTemplate}).Attrs(model.DataDictionary{
			Properties: []byte(`{"pwd": "c0079a7c663170bdd00aac4e0589d617", "uid": "dt8899", "text": "【球场新星】 您的验证码是: {**}, 5分钟内有效，请及时输入。"}`),
		}).FirstOrCreate(&model.DataDictionary{}).Error; err != nil {
			return err
		}

		if err := tx.Where(model.DataDictionary{DataType: consts.DataDictionaryMemberCategories}).Attrs(model.DataDictionary{
			Properties: []byte(`[
        {
            "name": "月度会员",
            "valid_day": 365,
            "present_price": 0.1,
            "original_price": 10,
            "extend": ""
        },
        {
            "name": "季度会员",
            "valid_day": 120,
            "present_price": 0.2,
            "original_price": 20,
            "extend": ""
        },
        {
            "name": "年度会员",
            "valid_day": 365,
            "present_price": 3,
            "original_price": 30,
            "extend": "改名卡"
        }
    ]`),
		}).FirstOrCreate(&model.DataDictionary{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
