package setting

import (
	"fmt"

	"github.com/Unknwon/goconfig"
)

const (
	DEFAULT_SECTION = "QRCODE"
	AppConfPath     = "./conf/app.ini"

	REDIS_SECTION = "REDIS"
)

var (
	AppVer           string
	AppLogLevel      int
	AppPort          string
	AppUri           string
	AppName          string
	IndexFileUri     string
	QRCodeSaveDir    string
	AppLogFileFormat string
	StaticDir        string
	Expire           int64

	RedisAddr     string
	RedisPwd      string
	RedisDb       int
	RedisPoolSize int
)

func loadConfig() {
	Cfg, err := goconfig.LoadConfigFile(AppConfPath)
	if err != nil {
		panic(fmt.Sprint("Fail to load configuration file: " + err.Error()))
	}
	AppName = Cfg.MustValue(DEFAULT_SECTION, "AppName", "QRCode")
	AppPort = ":" + Cfg.MustValue(DEFAULT_SECTION, "AppPort", "8080")
	AppVer = Cfg.MustValue(DEFAULT_SECTION, "AppVer", "0.1.1.0814")
	AppUri = Cfg.MustValue(DEFAULT_SECTION, "AppUri", "http://localhost:8080")
	AppLogFileFormat = Cfg.MustValue(DEFAULT_SECTION, "AppLogFileFormat", "ByHour")
	AppLogLevel = Cfg.MustInt(DEFAULT_SECTION, "AppLogLevel", 0)
	IndexFileUri = Cfg.MustValue(DEFAULT_SECTION, "IndexFileUri", "./index.html")
	QRCodeSaveDir = Cfg.MustValue(DEFAULT_SECTION, "QRCodeSaveDir", "./public/")
	StaticDir = Cfg.MustValue(DEFAULT_SECTION, "StaticDir", "./static/")
	Expire = Cfg.MustInt64(DEFAULT_SECTION, "Expire",3600000000000)
	RedisAddr = Cfg.MustValue(REDIS_SECTION, "Addr", "localhost:6379")
	RedisPwd = Cfg.MustValue(REDIS_SECTION, "Password", "")
	RedisDb = Cfg.MustInt(REDIS_SECTION, "DB", 0)
	RedisPoolSize = Cfg.MustInt(REDIS_SECTION, "PoolSize", 5)
}

func init() {
	loadConfig()
}
