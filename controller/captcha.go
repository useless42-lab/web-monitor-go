package controller

import (
	"WebMonitor/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id          string
	CaptchaType string
	VerifyValue string
	DriverDigit base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// 获取base64验证码基本配置
func GetCaptchaConfig() *configJsonBody {
	configJsonBody := &configJsonBody{
		CaptchaType: "digit",
		DriverDigit: base64Captcha.DriverDigit{Height: 70, Width: 155, Length: 5, MaxSkew: 1, DotCount: 100},
		Id:          "",
		VerifyValue: "",
	}
	return configJsonBody
}

// base64Captcha create http handler
func GenerateCaptchaHandler(context *gin.Context) {
	param := GetCaptchaConfig()
	driver := &param.DriverDigit
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	response.Success(context, 200, body)
}

// base64Captcha verify http handler
func VerifyCaptcha(id string, captcha string) bool {
	if store.Verify(id, captcha, true) {
		return true
	} else {
		return false
	}
}
