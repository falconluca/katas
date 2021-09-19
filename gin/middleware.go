package gin

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("greeting", "Luca")

		c.Next()

		latency := time.Since(t)
		log.Println(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

func MiddlewareEntry() {
	r := gin.New()
	r.Use(Logger())

	// Gin 内置中间件
	// gin.Logger()
	// gin.Recovery()
	// gin.CustomRecovery(handle gin.RecoveryFunc)
	// gin.BasicAuth()

	// 开源中间件
	// gin-jwt
	// gin-swagger
	// cors
	// sessions
	// authz
	// pprof
	// go-gin-prometheus *
	// gzip
	// gin-limit *
	// requestid *

	// 中间件的作用范围
	r.Use(gin.Logger(), gin.Recovery()) // 所有HTTP请求
	v1 := r.Group("/v1").Use(gin.BasicAuth(gin.Accounts{"luca": "f"}))
	v1.GET("greetings", nil).Use(gin.BasicAuth(gin.Accounts{"luca": "f"}))

	// 自定义中间件
	r.GET("/greetings", func(c *gin.Context) {
		name := c.MustGet("greeting").(string)
		log.Println(name)
	})

	// 可以使用Gin中间件实现认证、requestId、跨域

	r.Run(":8080")
}
