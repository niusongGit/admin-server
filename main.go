package main

import (
	"admin-server/internal/consts"
	"admin-server/internal/middleware"
	"admin-server/internal/model"
	"admin-server/internal/routers"
	"admin-server/pkg/crontask"
	"admin-server/pkg/goredis"
	"admin-server/pkg/logger"
	"admin-server/pkg/orm"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate/locales/zhcn"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
)

var configFile = flag.String("f", "application", "the config file")

func main() {

	flag.Parse()

	InitConfig()
	logger.InitLogger() // 创建记录日志的文件

	orm.InitDb()

	rdb := goredis.InitRedis()
	defer rdb.Close()

	err := initData()
	if err != nil {
		panic(err)
	}

	// for all Validation.
	// NOTICE: 必须在调用 validate.New() 前注册, 它只需要一次调用。
	zhcn.RegisterGlobal()

	r := gin.Default()

	r.Use(middleware.Recover)
	r.Use(middleware.LoggerToFile(), middleware.Cors())
	r = routers.CollectRoute(r)

	crontask.InitCron()

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()                 //获取当前工作目录
	viper.SetConfigName(*configFile)         //设置文件名
	viper.SetConfigType("yml")               //设置文件类型
	viper.AddConfigPath(workDir + "/config") //设置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initData() (err error) {
	var words []interface{}
	err = orm.GetDB().Model(&model.SensitiveWord{}).Where("status = ?", 0).Pluck("word", &words).Error
	if err != nil {
		return err
	}
	if len(words) == 0 {
		return nil
	}

	err = goredis.GetRedisDB().SAdd(context.Background(), consts.RedisSensitiveWord, words...).Err()
	if err != nil {
		return err
	}
	return
}
