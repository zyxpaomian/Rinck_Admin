package crontabmgtdao

import (
	"util/mysql"
	ce "util/error"
	"util/log"
	"structs"
)

type CronTabDao struct {
}

var CronDao CronTabDao

//获取子任务详情
func (crontabdao *CronTabDao) GetCronChildrenList(parentid int64) ([]*structs.CronChildren ,error) {
	resultlist := []*structs.CronChildren{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select level,exec_time,job_url,jobtype from cron_jobs_children_cfg where parentid = ?")
	if err != nil {
		tx.Rollback()
		log.Errorln("MySQL 获取TX失败: ",err.Error())
		return nil, ce.DBError()
	}
	rows, err := stmt.Query(parentid)
	if err != nil {
		log.Errorln("MySQL 查询失败: ",err.Error())
		stmt.Close()
		return nil, ce.DBError()
	}
	for rows.Next() {
		result := &structs.CronChildren{}
		err := rows.Scan(&result.Level,&result.Exectime,&result.Job_url,&result.Jobtype)
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

// 获取任务父ID
func (crontabdao *CronTabDao) GetCronParentList() ([]*structs.CronParent ,error) {
	resultlist := []*structs.CronParent{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select id, level, job_url, jobtype from cron_jobs_parent_cfg")
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
		result := &structs.CronParent{}
		err := rows.Scan(&result.Id,&result.Level,&result.Job_url,&result.Jobtype)
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

//子任务更新
func (crontabdao *CronTabDao) UpdateCronChildren(job_url string, newjobtype string, exec_time string, jobtype string) (int ,error) {
	updateid, err := mysql.DB.SimpleUpdate("update cron_jobs_children_cfg set job_url = ?, jobtype = ? ,exec_time = ? where jobtype = ?", job_url,newjobtype,exec_time,jobtype)
	// fmt.Println(err)
	if err != nil || updateid == 0{
			log.Errorln("更新计划任务子任务失败")
			return -1,ce.DBError()
		}
		return updateid, nil
}

//子任务删除
func (crontabdao *CronTabDao) DelCronChildren(jobtype string) (int ,error) {
	updateid, err := mysql.DB.SimpleUpdate("delete from cron_jobs_children_cfg where jobtype = ?",jobtype)
	// fmt.Println(err)
	if err != nil || updateid == 0{
			log.Errorln("删除计划任务子任务失败")
			return -1,ce.DBError()
		}
		return updateid, nil
}


//子任务添加
func (crontabdao *CronTabDao) AddCronChildren(job_url string, childjobtype string, parentjobtype string, exec_time string ) (int ,error) {
	updateid, err := mysql.DB.SimpleInsert("insert into cron_jobs_children_cfg (level,job_url,jobtype,parentid,exec_time) values (2,?,?,(select id from cron_jobs_parent_cfg where jobtype = ?),?)",job_url, childjobtype, parentjobtype, exec_time)
	// fmt.Println(err)
	if err != nil || updateid == 0{
			log.Errorln("添加计划任务子任务失败")
			return -1,ce.DBError()
		}
		return updateid, nil
}

func (crontabdao *CronTabDao) GetTaskList() ([]*structs.CronTask ,error) {
	resultlist := []*structs.CronTask{}
	tx := mysql.DB.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return nil, ce.DBError()
	}
	stmt, err := tx.Prepare("select job_url,exec_time from cron_jobs_children_cfg;")
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
		result := &structs.CronTask{}
		err := rows.Scan(&result.Job_url,&result.Exec_time)
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

