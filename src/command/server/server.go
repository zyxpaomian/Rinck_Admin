package main

import (
	"flag"
	"util/config"
	"util/log"
	"util/mysql"
	
	"fmt"
	"service/http_service"
	"service/crontab"
	"time"
	"runtime"
	"math/rand"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	var confPath = flag.String("confPath", "conf/default.ini","load conf file")
	flag.Parse()

	//配置文件初始化
	config.GlobalConf.CfgInit(*confPath)
	fmt.Println("配置文件初始化完成....")


	//日志初始化
	log.InitLog()
	fmt.Println("日志文件初始化完成....")

	//DB初始化
	mysql.DB.DbInit()
	fmt.Println("数据库初始化完成....")

	// resultlist, _ := cronctrl.GetCronTaskList()
	// fmt.Println(resultlist)

	//调度任务执行
	crontab.GlobalCronTaskList.CronInit()

	//内部事务调度，如数据库更新等等
	crontab.GlobalInternalCronTaskList.CronInit()


	//启动HTTP
	fmt.Println("开始启动HTTP服务")
	
	router := http_service.InitRouter()
	router.Run(":8000")
	//fmt.Println("启动HTTP服务器成功")
	//fmt.Println(http_srv.ListenAndServe())
	//select {}
}