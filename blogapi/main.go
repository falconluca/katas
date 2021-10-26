package main

import (
	"blogapi/models"
	"blogapi/pkg/logging"
	"blogapi/pkg/setting"
	"blogapi/routers"
	"fmt"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSettings.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSettings.ReadTimeout,
		WriteTimeout:   setting.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
