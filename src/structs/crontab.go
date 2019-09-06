package structs

type CronChildren struct {
	Level int64 `json:level`
	Exectime string `json:exectime`
	Job_url string `json:job_url`
	Jobtype string `json:jobtype`
}

type CronParent struct {
	Id int64 `json:id`
	Level int64 `json:level`
	Job_url string `json:job_url`
	Jobtype string `json:jobtype`
	Children []*CronChildren `json:children`
}

type CronTask struct {
	Exec_time string `json:exectime`
	Job_url string `json:job_url`
}


