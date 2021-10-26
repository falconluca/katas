package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PageSize  int

	RuntimeRootPath string

	ImgPrefixUrl string
	ImgSavePath  string
	ImgMaxSize   int
	ImgAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	DbName      string
	TablePrefix string
}

var (
	AppSettings      = &App{}
	ServerSettings   = &Server{}
	DatabaseSettings = &Database{}
)

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载配置文件失败 err: %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppSettings.ImgMaxSize = AppSettings.ImgMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSettings err: %v", err)
	}
	ServerSettings.ReadTimeout = ServerSettings.ReadTimeout * time.Second
	ServerSettings.WriteTimeout = ServerSettings.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSettings err: %v", err)
	}
}
