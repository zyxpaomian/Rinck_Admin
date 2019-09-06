package http_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controller/userctrl"
	//"controller/cronctrl"
	//"structs"
)

func UpdateRole(c *gin.Context) {
	type reqContent struct {
		Username	string `json:"username" binding:"required"`
		Role string `json:"role" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {
		username := r.Username
		role := r.Role
		_, err := userctrl.UpdateUserRole(username,role)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"更新用户权限失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"更新用户权限成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func DelUser(c *gin.Context) {
	type reqContent struct {
		Username	string `json:"username" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {
		username := r.Username
		_, err := userctrl.DelUser(username)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"删除用户失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"删除用户成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func ResetPassword(c *gin.Context) {
	type reqContent struct {
		Username	string `json:"username" binding:"required"`
		Password	string `json:"password" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {
		username := r.Username
		password := r.Password
		_, err := userctrl.ResetPassword(username,password)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"重置密码失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"重置密码成功",
				})
		}
   } else{
	c.JSON(http.StatusBadRequest, gin.H{
		"Status":"Failed",
		"Msg":"参数获取失败",
	})
   }
}

func AddUser(c *gin.Context) {
	type reqContent struct {
		Username	string `json:"username" binding:"required"`
		Password	string `json:"password" binding:"required"`
		Role	string `json:"role" binding:"required"`
	}
	var r reqContent
	err := c.ShouldBindJSON(&r)
	if err == nil {
		username := r.Username
		password := r.Password
		role := r.Role
		_, err := userctrl.AddUser(username,password,role)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Failed",
				"Msg":"添加用户失败",
				})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":"Success",
				"Msg":"添加用户成功",
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


	func GetRole(c *gin.Context) {
		resultlist, _ := userctrl.GetRoleList()
			
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