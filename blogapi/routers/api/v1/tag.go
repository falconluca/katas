package v1

import (
	"blogapi/models"
	"blogapi/pkg/e"
	"blogapi/pkg/setting"
	"blogapi/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

	"github.com/astaxie/beego/validation"
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
	data["list"] = models.GetTags(util.GetPage(c), setting.AppSettings.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	// 响应请求
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {object} gin.H
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name") // req.getParameter("name");
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	vaild := validation.Validation{}
	vaild.Required(name, "name").Message("名称不能为空")
	vaild.MaxSize(name, 12, "name").Message("名称最长为12个字符")
	vaild.Required(createdBy, "created_by").Message("创建人不能为空")
	vaild.MaxSize(createdBy, 8, "created_by").Message("创建人最长为8个字符")
	vaild.Range(state, 0, 1, "state").Message("状态只允许0或1")

	var code int
	if vaild.HasErrors() {
		code = e.INVALID_PARAMS
	} else if models.ExistTagByName(name) {
		code = e.ERROR_EXIST_TAG
	} else {
		models.AddTag(name, state, createdBy)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {object} gin.H
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 8, "modified_by").Message("修改人最长为8个字符")
	valid.MaxSize(name, 12, "name").Message("名称做长为12个字符")

	var code int
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
	} else if models.ExistTagById(id) {
		code = e.SUCCESS
		data := make(map[string]interface{})
		data["modified_by"] = modifiedBy
		if name != "" {
			data["name"] = name
		}
		if state != -1 { // -1是初始值, 说明用户没有设置state
			data["state"] = state
		}

		models.EditTag(id, data)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("Id必须大于0")

	var code int
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
	} else if models.ExistTagById(id) {
		models.DeleteTag(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
