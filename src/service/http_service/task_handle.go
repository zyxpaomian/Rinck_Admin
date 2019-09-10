package http_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controller/taskctrl"

)

func GetSyncTaskList(c *gin.Context) {
	resultlist, _ := taskctrl.GetSyncTaskList()
			
	if resultlist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"获取同步任务失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Msg": resultlist,
		})
	}
}

func GetRsyncTaskList(c *gin.Context) {
	resultlist, _ := taskctrl.GetRsyncTaskList()
			
	if resultlist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"获取异步任务失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Msg": resultlist,
		})
	}
}

func GetRsyncTaskRecords(c *gin.Context) {
	resultlist, _ := taskctrl.GetRsyncTaskRecords()
			
	if resultlist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"获取异步任务失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Msg": resultlist,
		})
	}
}


func GetRsyncTaskResults(c *gin.Context) {
	type reqContent struct {
		Resultid string `json:"resultid" binding:"required"`
	}

	var r reqContent
	err := c.ShouldBindJSON(&r)

	if err == nil {
		resultid := r.Resultid
		resultlist, _ := taskctrl.GetRsyncTaskResult(resultid)				
		if resultlist == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":"Failed",
				"Msg":"获取异步任务结果失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg": resultlist,
			})
		}
	}
}


func AddSyncTask(c *gin.Context) {
	type reqContent struct {
		Taskname string `json:"taskname" binding:"required"`
		Taskurl string `json:"taskurl" binding:"required"`
		Taskargs string `json:"taskargs" binding:"required"`
		Taskmod string `json:"taskmod" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {

		taskname := r.Taskname
		taskurl := r.Taskurl
		taskargs := r.Taskargs
		taskmod := r.Taskmod

		_, err := taskctrl.AddSyncTask(taskname, taskurl, taskargs, taskmod)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"添加任务失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"添加任务成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func AddSyncTaskRecord(c *gin.Context) {
	type reqContent struct {
		Taskname string `json:"taskname" binding:"required"`
		Celeryid string `json:"celeryid" binding:"required"`
		Execuser string `json:"execuser" binding:"required"`
		Execstate string `json:"execstate" binding:"required"`
		Exectype string `json:"exectype" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {

		taskname := r.Taskname
		celeryid := r.Celeryid
		execuser := r.Execuser
		execstate := r.Execstate
		exectype := r.Exectype

		_, err := taskctrl.AddSyncTaskRecord(taskname, celeryid, execuser, execstate, exectype)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"添加任务纪录失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"添加任务纪录成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}