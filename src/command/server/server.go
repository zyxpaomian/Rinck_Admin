package main

import (
	"flag"
	"util/config"
	"util/log"
	"util/mysql"
	"fmt"
	//"github.com/astaxie/beego"
	//"net/http"
	"service/http_service"
	//"service/salt"
	//"service/mydb"
	//"github.com/robfig/cron"
)


type TestJob struct {
}

func (this TestJob)Run() {
    fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job)Run() {
    fmt.Println("testJob2...")
}

func main(){


    //c := cron.New()

    //AddFunc
    //spec := "*/5 * * * * ?"


    //AddJob方法
    /*c.AddJob(spec, TestJob{})
    c.AddJob(spec, Test2Job{})*/

    //启动计划任务
    /*c.Start()

    //关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()*/
	

	
	//fmt.Println(aaa)




	var confPath = flag.String("confPath", "conf/default.ini","load conf file")
	flag.Parse()

	//配置文件初始化
	config.GlobalConf.CfgInit(*confPath)
	fmt.Println("配置文件初始化完成....")

	//日志初始化
	log.InitLog()
	fmt.Println("日志文件初始化完成....")

	//
	//salt.GetToken()
	//aaaa := salt.ModelExec("hello.haha","local","*","fuck")
	//fmt.Println(aaaa)
	

	//DB初始化
	mysql.DB.DbInit()
	fmt.Println("数据库初始化完成....")

	
	fmt.Println("开始启动HTTP服务")
	
	router := http_service.InitRouter()
	router.Run(":8000")
	//fmt.Println("启动HTTP服务器成功")
	//fmt.Println(http_srv.ListenAndServe())
	//select {}
}