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
	router.POST("/api/v1/changepasswd", ChangePasswd)
	router.POST("/api/v1/user/userauth", UserAuth)
	router.GET("/api/v1/user/userinfo",GetUser)
	return router
}

