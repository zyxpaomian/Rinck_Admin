package http_service

import (
	//"controller/saltrun"
	"github.com/gin-gonic/gin"
	"net/http"
	//"controller/userctrl"
	"controller/cronctrl"
	//"structs"
)


func GetAllCronList(c *gin.Context) {
	resultlist, _ := cronctrl.GetAllCronList()
			
	if resultlist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"获取计划任务数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Msg": resultlist,
					//"Msg":&result.Username,
		})
	}
}

func GetParentCronList(c *gin.Context) {
	resultlist, _ := cronctrl.GetParentCronList()
			
	if resultlist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"获取父任务数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Msg": resultlist,
					//"Msg":&result.Username,
		})
	}
}

func UpdateCronChildren(c *gin.Context) {
	type reqContent struct {
		Job_url string `json:"job_url" binding:"required"`
		Jobtype string `json:"jobtype" binding:"required"`
		Newjobtype string `json:"newjobtype" binding:"required"`
		Exec_time string `json:"exec_time" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {
		job_url := r.Job_url
		newjobtype := r.Newjobtype
		jobtype := r.Jobtype
		exec_time := r.Exec_time
		_, err := cronctrl.UpdateCronChildren(job_url, newjobtype, jobtype, exec_time)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"更新子计划任务失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"更新子计划任务成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func DelCronChildren(c *gin.Context) {
	type reqContent struct {
		Jobtype string `json:"jobtype" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {

		jobtype := r.Jobtype

		_, err := cronctrl.DelCronChildren(jobtype)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"删除子计划任务失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"删除子计划任务成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func AddCronChildren(c *gin.Context) {
	type reqContent struct {
		Job_url string `json:"job_url" binding:"required"`
		Childjobtype string `json:"childjobtype" binding:"required"`
		Parentjobtype string `json:"parentjobtype" binding:"required"`
		Exec_time string `json:"exec_time" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {

		job_url := r.Job_url
		childjobtype := r.Childjobtype
		parentjobtype := r.Parentjobtype
		exec_time := r.Exec_time

		_, err := cronctrl.AddCronChildren(job_url, childjobtype, parentjobtype, exec_time)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"添加子计划任务失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"添加子计划任务成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}
