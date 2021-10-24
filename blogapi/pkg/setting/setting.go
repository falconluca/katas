package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration // TODO 查看time文档
	WriteTimeout time.Duration

	JwtSecret string
	PageSize  int
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载配置文件失败 err: %v", err)
	}

	loadCommon()
	loadServer()
	loadApp()
}

func loadApp() {
	sec := getSection("app")

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func loadServer() {
	sec := getSection("server")

	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func getSection(section string) *ini.Section {
	sec, err := Cfg.GetSection(section) // TODO Cfg.GetSection和Cfg.Section的区别
	if err != nil {
		log.Fatalf("加载区失败 err: %v", err)
	}
	return sec
}

func loadCommon() {
	RunMode = Cfg.Section("").
		Key("RUN_MODE").
		MustString("debug")
}
