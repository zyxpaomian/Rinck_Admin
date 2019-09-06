package taskdao

import (
	"util/mysql"
	ce "util/error"
	"util/log"
	"structs"
)

type SyncDao struct {
}

var SyncTaskDao SyncDao

func (syncdao *SyncDao) GetAllSyncTasks() ([]*structs.SyncTask ,error) {
	resultlist := []*structs.SyncTask{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select taskname,taskurl,taskargs from tasks_model where taskmod REGEXP '^sync';")
	if err != nil {
		tx.Rollback()
		log.Errorln("MySQL 获取TX失败: ",err.Error())
		return nil, ce.DBError()
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Errorln("MySQL 查询失败: ",err.Error())
		stmt.Close()
		return nil, ce.DBError()
	}
	for rows.Next() {
		result := &structs.SyncTask{}
		err := rows.Scan(&result.Taskname,&result.Taskurl,&result.Taskargs)
		if err != nil {
			log.Errorln("MySQL 查询失败: ",err.Error())
			rows.Close()
			stmt.Close()
			tx.Rollback()
			return nil,ce.DBError()
		} else {
			resultlist = append(resultlist, result)
		}
	}
	rows.Close()
	stmt.Close()
	tx.Commit()
	return resultlist, nil
}

func (syncdao *SyncDao) AddTask(taskname string, taskurl string, taskargs string, taskmod string) (int ,error) {
	updateid, err := mysql.DB.SimpleUpdate("insert into tasks_model (taskname,taskurl,taskargs,taskmod) values (?,?,?,?);", taskname, taskurl, taskargs, taskmod)
	// fmt.Println(err)
	if err != nil || updateid == 0{
			log.Errorln("添加同步任务失败")
			return -1,ce.DBError()
		}
		return updateid, nil
}

func (syncdao *SyncDao) InsertTaskState(taskname string, celeryid string, execuser string, execstate string, exectype string) (int ,error) {
	updateid, err := mysql.DB.SimpleUpdate("insert into task_record (taskname,celeryid,execuser,exectime,execstate,exectype) values (?,?,?,NOW(),?,?);", taskname, celeryid, execuser, execstate, exectype)
	// fmt.Println(err)
	if err != nil || updateid == 0{
			log.Errorln("添加同步任务失败")
			return -1,ce.DBError()
		}
		return updateid, nil
}

