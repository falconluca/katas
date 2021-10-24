package routers

import (
	"blogapi/pkg/setting"
	v1 "blogapi/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 根据id更新标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 根据id删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
