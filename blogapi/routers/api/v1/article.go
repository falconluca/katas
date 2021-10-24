package v1

import (
	"blogapi/models"
	"blogapi/pkg/e"
	"blogapi/pkg/setting"
	"blogapi/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

// GetArticle 获取文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")

	var code int
	var data interface{}
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		logValidError(valid.Errors)
	} else if models.ExistArticleById(id) {
		data = models.GetArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles 获取文章列表
func GetArticles(c *gin.Context) {
	v := validation.Validation{}
	var state = -1
	stateArg := c.Query("state")
	if stateArg != "" {
		v.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var tagId = -1
	tagIdArg := c.Query("tag_id")
	if tagIdArg != "" {
		v.Min(tagId, 1, "tag_id").Message("标签id必须大于0")
	}

	var code int
	data := make(map[string]interface{})
	if v.HasErrors() {
		// 处理请求参数错误
		code = e.INVALID_PARAMS
		logValidError(v.Errors)
	} else {
		// 正常流程
		code = e.SUCCESS

		maps := make(map[string]interface{})
		maps["state"] = com.StrTo(stateArg).MustInt()
		maps["tag_id"] = com.StrTo(tagIdArg).MustInt()

		data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	v := validation.Validation{}
	v.Min(tagId, 1, "tag_id").Message("标签id必须大于0")
	v.Required(title, "title").Message("标题不能为空")
	v.Required(desc, "desc").Message("简述不能为空")
	v.Required(content, "content").Message("内容不能为空")
	v.Required(createdBy, "created_by").Message("创建人不能为空")
	v.Range(state, 0, 1, "staet").Message("状态只允许为0或1")

	var code int
	if v.HasErrors() {
		code = e.INVALID_PARAMS
		logValidError(v.Errors)
	} else if models.ExistTagById(tagId) {
		newArticle := make(map[string]interface{})
		newArticle["tag_id"] = tagId
		newArticle["title"] = title
		newArticle["desc"] = desc
		newArticle["content"] = content
		newArticle["created_by"] = createdBy
		newArticle["state"] = state

		models.AddArticle(newArticle)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	var state = -1
	if arg := c.Query("state"); arg != "" { // TODO c.Query和c.Param的区别
		state = com.StrTo(arg).MustInt()
	}

	v := validation.Validation{}
	if state != -1 {
		v.Range(state, 0, 1, "state").Message("状态只允许为0或1")
	}
	v.Min(id, 1, "id").Message("id必须大于0")
	v.MaxSize(title, 100, "title").Message("标题最长为100个字符")
	v.MaxSize(desc, 255, "desc").Message("简述最长为255个字符")
	v.MaxSize(content, 65535, "content").Message("内容最长为65535个字符")
	v.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	v.MaxSize(modifiedBy, 20, "modified_by").Message("修改人最长为100个字符")

	var code int
	if v.HasErrors() {
		logValidError(v.Errors)
		code = e.INVALID_PARAMS
	} else if models.ExistArticleById(id) {
		if models.ExistTagById(tagId) {
			article := make(map[string]interface{})
			article["tag_id"] = tagId
			article["title"] = title
			article["desc"] = desc
			article["content"] = content
			article["modified_by"] = modifiedBy

			models.EditArticle(id, article)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	v := validation.Validation{}
	v.Min(id, 1, "id").Message("id必须大于0")

	var code int
	if v.HasErrors() {
		logValidError(v.Errors)
		code = e.INVALID_PARAMS
	} else if models.ExistArticleById(id) {
		models.DeleteArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func logValidError(errors []*validation.Error) {
	for _, err := range errors {
		log.Printf("err.key: %s, err.msg: %s", err.Key, err.Message)
	}
}
