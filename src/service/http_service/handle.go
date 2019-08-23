package http_service

import (
	"controller/saltrun"
	"github.com/gin-gonic/gin"
	"net/http"
	"controller/userctrl"
	//"structs"
)

type Response struct {
	Status string `json:"Status"`
	Result map[string]string `json:"Result"`
}

func ChangePasswd(c *gin.Context) {
	type reqContent struct {
		Ip	string `json:"ip" binding:"required"`
		Func string `json:"func" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {
		ipList := r.Ip
		funcName := r.Func
		chResult := saltrun.ChangePasswd(ipList,funcName)
		if len(chResult) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"Salt执行失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":chResult,
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func UserAuth(c *gin.Context) {
	type reqContent struct {
		Username string `json:username`
		Password string `json:password`
	}
	var r reqContent
	//result := &structs.UserInfo{}
	
	err := c.ShouldBindJSON(&r)
	if err == nil {
		username := r.Username
		password := r.Password
		result, _ := userctrl.CreateToken(username,password)
		
		if result == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":"Failed",
				"Msg":"用户认证失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				//"Msg":&result.Username,
				"Token":&result.Token,
				"Username":&result.Username,
				"Role":&result.Role,
				})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"参数获取失败",
	})
	}
}

func GetUser(c *gin.Context) {
	resultlist, _ := userctrl.GetUserInfo()
		
	if resultlist == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":"Failed",
			"Msg":"获取用户数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":"Success",
			"Msg": resultlist,
				//"Msg":&result.Username,
			})
		}
	}