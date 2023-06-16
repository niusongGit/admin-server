package crontask

import (
	"admin-server/internal/job"
	"github.com/robfig/cron/v3"
)

type CronTask struct {
	Cj *cron.Cron
}

func InitCron() {
	ct := &CronTask{
		Cj: cron.New(),
	}
	ct.RegisterTask()
	ct.Cj.Start()
}

func (ct *CronTask) RegisterTask() {
	ct.Cj.AddFunc("@every 1m", job.NewCompetitionCronTask().CompetitionStart)
}
