package Actions

import (
	"QRCodeService/setting"
	"os"
	"time"
)

func DestroyQRcode() {
	elem, t := RedisClient.GetDestroy()
	//fmt.Println(elem, t)
	du := t.Sub(time.Now())
	if du < 0 {
		du = 0
	}
	if elem != "" && du == 0 {
		RedisClient.DelQrcode(elem)
		os.Remove(setting.QRCodeSaveDir + elem)
	}
	time.AfterFunc(du, DestroyQRcode)
}
