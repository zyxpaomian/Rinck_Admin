package taskdao

import (
	"util/mysql"
	ce "util/error"
	"util/log"
	"structs"
)

type RsyncDao struct {
}

var RsyncTaskDao RsyncDao

func (rsyncdao *RsyncDao) GetAllRsyncTasks() ([]*structs.RsyncTask ,error) {
	resultlist := []*structs.RsyncTask{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select taskname,taskurl,taskargs from tasks_model where taskmod REGEXP '^rsync';")
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
		result := &structs.RsyncTask{}
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


func (rsyncdao *RsyncDao) GetRsyncTaskRecords() ([]*structs.TaskRecord ,error) {
	resultlist := []*structs.TaskRecord{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select taskname ,celeryid, execuser, DATE_FORMAT(exectime,'%Y-%m-%d %H:%i:%S'), execstate, exectype from task_record where exectype REGEXP '^rsync' order by exectime desc;")
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
		result := &structs.TaskRecord{}
		err := rows.Scan(&result.Taskname, &result.Celeryid, &result.Execuser, &result.Exectime, &result.Execstate, &result.Exectype)
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

func (rsyncdao *RsyncDao) UpdateTaskRecord(state string,celeryid string) (int,error) {
	updateid, err := mysql.DB.SimpleUpdate("update task_record set execstate = ? where celeryid = ?;",state,celeryid)
	if err != nil || updateid == -1{
		log.Errorln("更新task纪录失败")
		return -1,ce.DBError()
	}
	return updateid, nil
}

func (rsyncdao *RsyncDao) GetRsyncUnfinishRecords() ([]*structs.TaskRecord ,error) {
	resultlist := []*structs.TaskRecord{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select celeryid, DATE_FORMAT(exectime,'%Y-%m-%d %H:%i:%S'), execstate from task_record where exectype REGEXP '^rsync' and execstate = 'processing';")
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
		result := &structs.TaskRecord{}
		err := rows.Scan(&result.Celeryid, &result.Exectime, &result.Execstate)
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
