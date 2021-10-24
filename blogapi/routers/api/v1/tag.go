package v1

import (
	"blogapi/models"
	"blogapi/pkg/e"
	"blogapi/pkg/setting"
	"blogapi/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context) {
	// 组装数据访问参数
	maps := make(map[string]interface{})
	name := c.Query("name") // TODO 处理?name=test&state=1这种请求参数
	//name = c.DefaultQuery("name", "") // 支持默认值
	if name != "" {
		maps["name"] = name
	}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	// 调用数据访问
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	// 响应请求
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
