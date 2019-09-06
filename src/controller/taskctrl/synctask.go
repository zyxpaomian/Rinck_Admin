package taskctrl

import (
	"dao/taskdao"
	"util/log"
	ce "util/error"
	"structs"
)

func GetSyncTaskList() ([]*structs.SyncTask, error){
	tasklist := []*structs.SyncTask{}
	tasklist, err := taskdao.SyncTaskDao.GetAllSyncTasks()
	if err != nil {
		log.Errorln("获取同步任务清单失败")
		return nil, ce.GetSyncTaskError()
	}
	return tasklist, nil
}

func AddSyncTask(taskname string, taskurl string, taskargs string, taskmod string) (int, error){		
	insertid, err := taskdao.SyncTaskDao.AddTask(taskname, taskurl, taskargs, taskmod)
	if err != nil || insertid == -1{
		log.Errorln("添加任务失败")
		return -1,ce.DBError()
	}
	return 1,nil
}

func AddSyncTaskRecord(taskname string, celeryid string, execuser string, execstate string,exectype string) (int, error){		
	insertid, err := taskdao.SyncTaskDao.InsertTaskState(taskname, celeryid, execuser, execstate, exectype)
	if err != nil || insertid == -1{
		log.Errorln("添加任务纪录失败")
		return -1,ce.DBError()
	}
	return 1,nil
}