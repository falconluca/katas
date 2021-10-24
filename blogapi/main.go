package main

import (
	"blogapi/pkg/setting"
	"blogapi/routers"
	"fmt"
	"net/http"
)

func main() {
	//r := gin.Default()
	//r.GET("/healthz", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "success",
	//	})
	//})
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
