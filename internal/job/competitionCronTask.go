package job

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/pkg/logger"
	"admin-server/pkg/orm"
	"admin-server/pkg/orm/datatypes"
	"admin-server/pkg/utils"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type CompetitionCronTask struct{}

func (c *CompetitionCronTask) CompetitionStart() {
	// 向日志中添加traceId，在每个请求中加入追踪编号
	traceId := fmt.Sprintf("CronTask-%s-%v", utils.RandomString(20), time.Now().UnixMilli())
	ctx := context.Background()
	logger, ctx := logger.AddCtxWithTraceId(ctx, traceId)

	if err := orm.GetDBWithContext(ctx).Model(&model.Competition{}).Where("status = ? and start_time <= ?", consts.CompetitionStatusNotStarted, datatypes.XTime{time.Now()}).Update("status", consts.CompetitionStatusStarted).Error; err != nil {
		logger.Sugar().With(zap.String("log_from", "job")).Error(err)
	}

	return
}

func NewCompetitionCronTask() *CompetitionCronTask {
	return &CompetitionCronTask{}
}
