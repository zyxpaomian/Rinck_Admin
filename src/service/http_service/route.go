package http_service

import (
	"github.com/gin-gonic/gin"
	"util/config"
	"os"
)

func InitRouter() *gin.Engine {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	logdir := config.GlobalConf.GetStr("server","logdir")
	logfile := config.GlobalConf.GetStr("server","httplogname")
	file, err := os.OpenFile(logdir + "/" + logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic("打开Gin HTTP日志文件失败")
	}
	gin.DefaultWriter = file

	router := gin.Default()
	// 用户相关接口
	router.POST("/api/v1/user/userauth", UserAuth)
	router.GET("/api/v1/user/userinfo",GetUser)
	router.GET("/api/v1/user/rolelist",GetRole)
	router.POST("/api/v1/user/roleupdate",UpdateRole)
	router.POST("/api/v1/user/userdel",DelUser)
	router.POST("/api/v1/user/useradd",AddUser)
	router.POST("/api/v1/user/resetpassword",ResetPassword)
	// 计划任务相关
	router.GET("/api/v1/crontab/getjoblist",GetAllCronList)
	router.GET("/api/v1/crontab/getparentlist",GetParentCronList)
	router.POST("/api/v1/crontab/updatechildjob",UpdateCronChildren)
	router.POST("/api/v1/crontab/delchildjob",DelCronChildren)
	router.POST("/api/v1/crontab/addchildjob",AddCronChildren)
	// 任务执行相关
	router.GET("/api/v1/task/getsynctask",GetSyncTaskList)
	router.GET("/api/v1/task/getrsynctask",GetRsyncTaskList)
	router.GET("/api/v1/task/getrsynctaskrecords",GetRsyncTaskRecords)
	router.POST("/api/v1/task/addtask",AddSyncTask)
	router.POST("/api/v1/task/addtaskrecord",AddSyncTaskRecord)
	return router
}

