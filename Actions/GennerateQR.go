package Actions

import (
	"QRCodeService/setting"
	"crypto/md5"
	"encoding/hex"
	"time"

	"QRCodeService/model"

	"github.com/lunny/tango"
	"github.com/skip2/go-qrcode"
)

var RedisClient *model.RedisClient

func init() {
	RedisClient = model.NewClient()
}

type GennerateQR struct {
	tango.Json
	tango.Ctx
}

func getName(content string) string {
	md5ctx := md5.New()
	md5ctx.Write([]byte(content))
	name := md5ctx.Sum(nil)
	return hex.EncodeToString(name)
}

func (h *GennerateQR) Get() interface{} {
	m := h.Form("msg")
	if m == "" {
		return OutModel{Code: 200, Msg: "没有内容，无法产生"}
	}
	name := getName(m)
	var expTime string
	if !RedisClient.IsMember("Qrcode", name) {
		expTime = time.Now().Add(time.Hour * 2).Format(time.RFC3339)
		RedisClient.SaveQrcode("Qrcode", name, expTime)
	} else {
		expTime = RedisClient.GetExpire("Qrcode", name)
	}
	err := qrcode.WriteFile(m, qrcode.Medium, 256, "public/"+name)
	if err != nil {
		println("出错了" + err.Error())
	}
	uri := setting.AppUri + "/QrCode/" + name
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["url"] = uri
	m1["content"] = m
	m1["expDate"] = expTime
	return OutModel{Code: 200, Data: m1}
}

type OutModel struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
