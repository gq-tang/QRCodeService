package main

import (
	"QRCodeService/Actions"
	"QRCodeService/setting"
	"io"
	"os"

	"github.com/lunny/log"
	"github.com/lunny/tango"
)

func main() {
	defer Actions.RedisClient.Close()
	var s log.ByType
	switch setting.AppLogFileFormat {
	case "ByHour":
		s = log.ByHour
	case "ByMonth":
		s = log.ByMonth
	default:
		s = log.ByDay
	}
	w := log.NewFileWriter(log.FileOptions{
		ByType: s,
		Dir:    "./logs",
	})
	log.Std.SetOutput(io.MultiWriter(w, os.Stdout))
	log.Std.SetPrefix("[" + setting.AppName + "_" + setting.AppVer + "]")
	log.Std.SetOutputLevel(setting.AppLogLevel)
	t := tango.Classic(log.Std)
	t.Get("/static/*", tango.Dir(setting.StaticDir))
	t.Any("/GetQR", new(Actions.GennerateQR))
	t.Get("/", tango.File(setting.IndexFileUri))
	t.Get("/QrCode/:name", tango.Dir(setting.QRCodeSaveDir))
	t.Run(setting.AppPort)
}
