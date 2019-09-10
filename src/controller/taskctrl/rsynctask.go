package taskctrl

import (
	"dao/taskdao"
	"util/log"
	ce "util/error"
	"structs"
	"time"
	"util/config"
	"net/http"
    "io/ioutil"
    "encoding/json"
	"bytes"
	"github.com/bitly/go-simplejson"
)


func GetRsyncTaskList() ([]*structs.RsyncTask, error){
	tasklist := []*structs.RsyncTask{}
	tasklist, err := taskdao.RsyncTaskDao.GetAllRsyncTasks()
	if err != nil {
		log.Errorln("获取异步任务清单失败")
		return nil, ce.GetRsyncTaskError()
	}
	return tasklist, nil
}

func GetRsyncTaskRecords() ([]*structs.TaskRecord, error){
	tasklist := []*structs.TaskRecord{}
	tasklist, err := taskdao.RsyncTaskDao.GetRsyncTaskRecords()
	if err != nil {
		log.Errorln("获取异步任务清单失败")
		return nil, ce.GetRsyncTaskError()
	}
	return tasklist, nil
}


func GetRsyncTaskResult(resultid string) ([]*structs.TaskResult, error ){
	taskresultlist := []*structs.TaskResult{}
	url := config.GlobalConf.GetStr("api", "server_task_result_api")
	data := make(map[string]interface{})
	data["resultid"] = resultid

	bytesData, err := json.Marshal(data)
	if err != nil {
		log.Errorln("获取任务执行结果失败")
		return nil, ce.GetRsyncTaskResult()
	}
	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Errorln("获取任务执行结果失败")
		return nil, ce.GetRsyncTaskResult()
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Errorln("获取任务执行结果失败")
		return nil, ce.GetRsyncTaskResult()
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln("获取任务执行结果失败")
		return nil, ce.GetRsyncTaskResult()
	}
	resultjson, err := simplejson.NewJson([]byte(string(respBytes)))
	if err != nil {
		log.Errorln("获取任务执行结果JSON失败")
		return nil, ce.GetRsyncTaskResult()
	}

	failedrows, _ := resultjson.Get("message").Get("failed").Array()
	for _, failrow := range failedrows {
		taskresult := &structs.TaskResult{}
		rowdata := failrow.(map[string]interface{})
		taskresult.Taskip = rowdata["host"].(string)
		taskresult.Taskoutput = rowdata["msg"].(string)
		taskresult.Celeryid = resultid
		taskresult.Taskstatus = "failed"
		taskresultlist = append(taskresultlist,taskresult)
	}


	successrows, _ := resultjson.Get("message").Get("success").Array()
	for _, successrow := range successrows {
		taskresult := &structs.TaskResult{}
		rowdata := successrow.(map[string]interface{})
		taskresult.Taskip = rowdata["host"].(string)
		taskresult.Taskoutput = rowdata["msg"].(string)
		taskresult.Celeryid = resultid
		taskresult.Taskstatus = "success"
		taskresultlist = append(taskresultlist,taskresult)

	}


	unreachablerows, _ := resultjson.Get("message").Get("unreachable").Array()
	for _, unreachablerow := range unreachablerows {
		taskresult := &structs.TaskResult{}
		rowdata := unreachablerow.(map[string]interface{})
		taskresult.Taskip = rowdata["host"].(string)
		taskresult.Taskoutput = rowdata["msg"].(string)
		taskresult.Celeryid = resultid
		taskresult.Taskstatus = "unreachable"
		taskresultlist = append(taskresultlist,taskresult)
	}

	return taskresultlist, nil


}

func UpdateRsyncUnfinishRecords() {
	recordlist := []*structs.TaskRecord{}
	recordlist, err := taskdao.RsyncTaskDao.GetRsyncUnfinishRecords()
	if err != nil {
		log.Errorln("获取未完成任务清单失败")
	}
	for record:=0;record < len(recordlist); record++ {
		resultid := recordlist[record].Celeryid
		//exectime := recordlist[record].Exectime
		exectime, _ := time.ParseInLocation("2006-01-02 15:04:05", recordlist[record].Exectime, time.Local) 
		// execstate := recordlist[record].Execstate
		// fmt.Println(time.Now().Unix())
		// fmt.Println(exectime.Unix())

		// 超过6小时的任务置为超时
		if (time.Now().Unix() - exectime.Unix()) > 21600 {
			updateid, err := taskdao.RsyncTaskDao.UpdateTaskRecord("timeout",resultid)
			if err != nil || updateid == -1{
				log.Errorln("更新task纪录失败")
			}
		} else {
			url := config.GlobalConf.GetStr("api", "server_task_state_api")
			data := make(map[string]interface{})
			data["resultid"] = resultid

			bytesData, err := json.Marshal(data)
			if err != nil {
				log.Errorln("更新任务状态失败")
			}
			reader := bytes.NewReader(bytesData)

			request, err := http.NewRequest("POST", url, reader)
			if err != nil {
				log.Errorln("更新任务状态失败")
			}
			request.Header.Set("Content-Type", "application/json;charset=UTF-8")
			client := http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				log.Errorln("更新任务状态失败")
			}
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Errorln("更新任务状态失败")
			}
			//fmt.Println(string(respBytes))
			resultjson, _ := simplejson.NewJson([]byte(string(respBytes)))
			resultstr := resultjson.Get("message").MustString()

			if resultstr == "finished" {
				updateid, err := taskdao.RsyncTaskDao.UpdateTaskRecord(resultstr,resultid)
				if err != nil || updateid == -1{
					log.Errorln("更新task纪录失败")
				}
			}		
		}
    }
}
