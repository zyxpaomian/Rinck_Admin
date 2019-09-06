package cronctrl

import (
	"dao/crontabmgtdao"
	"util/log"
	ce "util/error"
	"structs"
	"service/crontab"
)


func GetAllCronList() ([]*structs.CronParent, error) {
	parent_result_list := []*structs.CronParent{}
	parent_result_list, err := crontabmgtdao.CronDao.GetCronParentList()
	if err != nil {
		log.Errorln("获取计划任务大类失败")
		return nil, ce.GetCronError()
	}
	for i := 0; i < len(parent_result_list); i++ {
		parentid := parent_result_list[i].Id
		children_result_list, err := crontabmgtdao.CronDao.GetCronChildrenList(parentid)
		if err != nil {
			log.Errorln("获取计划任务子类失败")
			return nil, ce.GetCronError()
		}
		parent_result_list[i].Children = children_result_list
    }
	return parent_result_list, nil
}

func GetParentCronList() ([]*structs.CronParent, error) {
	parent_result_list := []*structs.CronParent{}


	parent_result_list, err := crontabmgtdao.CronDao.GetCronParentList()
	if err != nil {
		log.Errorln("获取计划任务大类失败")
		return nil, ce.GetCronError()
	}
	return parent_result_list, nil
}

func UpdateCronChildren(job_url string, newjobtype string, exec_time string,jobtype string) (int, error){
	updateid, err := crontabmgtdao.CronDao.UpdateCronChildren(job_url, newjobtype, jobtype, exec_time)
	if err != nil || updateid == -1{
		log.Errorln("更新计划任务子任务失败")
		return -1,ce.DBError()
	}

	crontab.GlobalCronTaskList.CronClose()
	crontab.GlobalCronTaskList.CronInit()
	return 1,nil
}

func DelCronChildren(jobtype string) (int, error){
	updateid, err := crontabmgtdao.CronDao.DelCronChildren(jobtype)
	if err != nil || updateid == -1{
		log.Errorln("删除计划任务子任务失败")
		return -1,ce.DBError()
	}

	crontab.GlobalCronTaskList.CronClose()
	crontab.GlobalCronTaskList.CronInit()
	return 1,nil
}

func AddCronChildren(job_url string, childjobtype string, parentjobtype string, exec_time string) (int, error){
	updateid, err := crontabmgtdao.CronDao.AddCronChildren(job_url, childjobtype, parentjobtype, exec_time)
	if err != nil || updateid == -1{
		log.Errorln("添加计划任务子任务失败")
		return -1,ce.DBError()
	}

	crontab.GlobalCronTaskList.CronClose()
	crontab.GlobalCronTaskList.CronInit()
	return 1,nil
}

func GetCronTaskList() ([]*structs.CronTask, error) {
	task_result_list := []*structs.CronTask{}

	task_result_list, err := crontabmgtdao.CronDao.GetTaskList()
	if err != nil {
		log.Errorln("获取计划任务列表失败")
		return nil, ce.GetCronError()
	}
	return task_result_list, nil
}