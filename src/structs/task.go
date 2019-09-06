package structs

type SyncTask struct {
	Taskname string `json:taskname`
	Taskurl string `json:taskurl`
	Taskargs string `json:taskargs`
	Taskmod string `json:taskmod`
}


type RsyncTask struct {
	Taskname string `json:taskname`
	Taskurl string `json:taskurl`
	Taskargs string `json:taskargs`
	Taskmod string `json:taskmod`
}

type TaskRecord struct {
	Taskname string `json:taskname`
	Celeryid string `json:celeryid`
	Execuser string `json:execuser`
	Exectime string `json:exectime`
	Execstate string `json:execstate`
	Exectype string `json:exectype`
}

