package crontab

import (
	"github.com/robfig/cron"
	//"structs"
	//"util/log"
	"controller/taskctrl"
)

type InternalCronTaskList struct {
	InternalCrontabtask *cron.Cron
}

var GlobalInternalCronTaskList InternalCronTaskList

func (i *InternalCronTaskList) CronInit() {
	i.InternalCrontabtask = cron.New()
	i.InternalCrontabtask.Start()

	i.InternalCrontabtask.AddFunc("*/10 * * * * * ", func() {
		taskctrl.UpdateRsyncUnfinishRecords()
	})
	i.InternalCrontabtask.Start()
}

func (i *InternalCronTaskList) CronClose() {
	i.InternalCrontabtask.Stop()
}

func (i *InternalCronTaskList) CronStart() {
	i.InternalCrontabtask.Start()
}
