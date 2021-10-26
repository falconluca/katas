package util

import (
	"blogapi/pkg/setting"
	"github.com/gin-gonic/gin"
	// https://gowalker.org/github.com/unknwon/com
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0

	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSettings.PageSize
	}
	return result
}
