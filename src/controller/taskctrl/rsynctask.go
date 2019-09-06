package taskctrl

import (
	"dao/taskdao"
	"util/log"
	ce "util/error"
	"structs"
	// "fmt"
	"time"
	"util/config"
	"net/http"
    "io/ioutil"
    "encoding/json"
	"bytes"
	"github.com/bitly/go-simplejson"
	// "unsafe"
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
			//fmt.Println(resultstr)


			// fmt.Println(resultjson.Get("message").MustString())
			//fmt.Println(resultjson)
			//personArr, err := js.Get("person").Array()
			//fmt.Println(len(personArr))
			//byte数组直接转成string，优化内存
			// str := (*string)(unsafe.Pointer(&respBytes))
			// fmt.Println(*str)
			
		}
    }
}
