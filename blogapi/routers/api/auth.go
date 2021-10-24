package api

import (
	"blogapi/models"
	"blogapi/pkg/e"
	"blogapi/pkg/logging"
	"blogapi/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 校验用户名和密码
	v := validation.Validation{}
	auth := &auth{Username: username, Password: password}
	if ok, _ := v.Valid(auth); !ok {
		for _, err := range v.Errors {
			logging.Info(err.Key, err.Message)
		}
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": make(map[string]interface{}),
		})
		return
	}

	// 验证用户在系统中是否存在
	isExist := models.CheckAuth(username, password)
	if !isExist {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_AUTH,
			"msg":  e.GetMsg(e.ERROR_AUTH),
			"data": make(map[string]interface{}),
		})
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		logging.Info(err)
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_AUTH_TOKEN,
			"msg":  e.GetMsg(e.ERROR_AUTH_TOKEN),
			"data": make(map[string]interface{}),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": map[string]string{
			"token": token,
		},
	})
}
