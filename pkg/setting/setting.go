package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)


var (
	Cfg *ini.File

	RunModule string
	HttpPort int
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
	JwtSecret string
	PageSize int
)

func init() {
	var err error

	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Faile to parse 'conf/app.ini'", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunModule = Cfg.Section("").Key("RUN_MODULE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	RunModule = Cfg.Section("").Key("RUN_MODULE").MustString("debug")
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}