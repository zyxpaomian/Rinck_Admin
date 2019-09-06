package crontab

import (
	"github.com/robfig/cron"
	"structs"
	"dao/crontabmgtdao"
	"util/log"
	"util/tools"
)

type CronTaskList struct {
	Crontabtask *cron.Cron
	TaskList []*structs.CronTask
}

var GlobalCronTaskList CronTaskList

func (c *CronTaskList) CronInit() {
	task_result_list, err := crontabmgtdao.CronDao.GetTaskList()
	if err != nil {
		log.Errorf("获取计划任务列表失败")
	}
	
	c.TaskList = task_result_list

	c.Crontabtask = cron.New()

	for i := 0; i < len(c.TaskList); i++ {
		exectime := c.TaskList[i].Exec_time
		joburl := c.TaskList[i].Job_url
		c.Crontabtask.AddFunc(exectime, func() {
			tools.HttpGetClient(joburl)
		})
	}
	c.Crontabtask.Start()
}

func (c *CronTaskList) CronClose() {
	c.Crontabtask.Stop()
}
