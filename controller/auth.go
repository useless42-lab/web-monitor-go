package controller

import (
	"WebMonitor/cache"
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"WebMonitor/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	_ "github.com/joho/godotenv/autoload"
)

type UserLoginInfo struct {
	Token string `json:"token"`
}

type AuthForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type PasswordForm struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ResetPasswordForm struct {
	Token    string `json:"token"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

func (form AuthForm) ValidateLoginForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required.Error("邮箱不能为空"), is.Email.Error("邮箱格式有误")),
		validation.Field(&form.Password, validation.Required.Error("密码不能为空")),
		validation.Field(&form.Captcha, validation.Required.Error("验证码不能为空")),
	)
}

func (form AuthForm) ValidateRegisterForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Username, validation.Required.Error("用户名不能为空")),
		validation.Field(&form.Email, validation.Required.Error("邮箱不能为空"), is.Email.Error("请输入邮箱")),
		validation.Field(&form.Password, validation.Required.Error("密码不能为空")),
		validation.Field(&form.Captcha, validation.Required.Error("验证码不能为空")),
	)
}

func (form PasswordForm) ValidateChangePasswordForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.OldPassword, validation.Required.Error("旧密码不能为空")),
		validation.Field(&form.NewPassword, validation.Required.Error("新密码不能为空")),
	)
}

func (form AuthForm) ValidateForgetForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required.Error("邮箱不能为空"), is.Email.Error("邮箱格式有误")),
		validation.Field(&form.Captcha, validation.Required.Error("验证码不能为空")),
	)
}

func (form ResetPasswordForm) ValidateResetPasswordForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Token, validation.Required.Error("用户名不能为空")),
		validation.Field(&form.Password, validation.Required.Error("密码不能为空")),
		validation.Field(&form.Captcha, validation.Required.Error("验证码不能为空")),
	)
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	captchaId := c.PostForm("captcha_id")
	loginForm := AuthForm{
		Email:    email,
		Password: password,
		Captcha:  captcha,
	}
	err := loginForm.ValidateLoginForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	err1 := VerifyCaptcha(captchaId, captcha)
	if !err1 {
		response.Error(c, 4000, "验证码有误")
		return
	}
	result := OnLoginApi(email, password)
	if result.Error == "" && result.Code == 200 {
		userId, _ := strconv.ParseInt(result.Data.Id, 10, 64)
		isExist := models.CheckIsExistUserId(userId)
		if !isExist {
			// 2022年5月1日前注册的用户，默认三个月试用计划
			targetTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-05-01 00:00:00", time.Local)
			if time.Now().After(targetTime) {
				service.InitUserService(userId, result.Data.Username, result.Data.Email, time.Now().AddDate(0, 1, 0))
			} else {
				service.InitUserService(userId, result.Data.Username, result.Data.Email, time.Now().AddDate(0, 3, 0))
			}
			// service.InitUserService(userId, result.Data.Username, result.Data.Email)
		}
		token := utils.GetTokenFromUser(userId)
		if token == "" {
			token = utils.GetRandomStr(88)
		}
		utils.SetTokenToUser(userId, token)
		utils.SetUserToToken(token, userId)
		response.Success(c, 200, UserLoginInfo{
			token,
		})
	} else {
		response.Error(c, 4001, "用户名密码错误")
	}
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	captchaId := c.PostForm("captcha_id")
	registerForm := AuthForm{
		Username: username,
		Email:    email,
		Password: password,
		Captcha:  captcha,
	}
	err := registerForm.ValidateRegisterForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	err1 := VerifyCaptcha(captchaId, captcha)
	if !err1 {
		response.Error(c, 4000, "验证码有误")
		return
	}
	result := OnRegisterApi(username, email, password)
	if result.Error == "" && result.Code == 200 {
		userId, _ := strconv.ParseInt(result.Data.Id, 10, 64)
		isExist := models.CheckIsExistUserId(userId)
		if !isExist {
			// 2022年5月1日前注册的用户，默认三个月试用计划
			targetTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-05-01 00:00:00", time.Local)
			if time.Now().After(targetTime) {
				service.InitUserService(userId, result.Data.Username, result.Data.Email, time.Now().AddDate(0, 1, 0))
			} else {
				service.InitUserService(userId, result.Data.Username, result.Data.Email, time.Now().AddDate(0, 3, 0))
			}
			// service.InitUserService(userId, result.Data.Username, result.Data.Email)
		}
		response.Success(c, 200, "")
	} else {
		response.Error(c, result.Code, result.Error)
	}
}

func ChangePassword(c *gin.Context) {
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")
	passwordForm := PasswordForm{
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}
	err := passwordForm.ValidateChangePasswordForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	if c.PostForm("re_new_password") != newPassword {
		response.Error(c, 4000, "新密码不一致")
		return
	}
	userId := strconv.FormatInt(c.GetInt64("userId"), 10)
	result := OnChangePasswordApi(userId, oldPassword, newPassword, c.PostForm("re_new_password"))
	if result.Error == "" && result.Code == 200 {
		token := utils.GetTokenFromUser(c.GetInt64("userId"))
		utils.DelTokenToUser(userId)
		utils.DelUserToToken(token)
		response.Success(c, 200, "")
	} else {
		response.Error(c, 4001, "密码有误")
	}
}

func CheckAuth(c *gin.Context) {
	response.Success(c, 200, "")
}

func GenerateResetPasswordLink(c *gin.Context) {
	email := c.PostForm("email")
	token := utils.GetRandomStr(32)
	captcha := c.PostForm("captcha")
	captchaId := c.PostForm("captcha_id")
	forgetForm := AuthForm{
		Email:   email,
		Captcha: captcha,
	}
	err := forgetForm.ValidateForgetForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	err1 := VerifyCaptcha(captchaId, captcha)
	if !err1 {
		response.Error(c, 4000, "验证码有误")
		return
	}
	resetStatus := cache.Get("reset:email:" + email)
	if resetStatus == "" {
		cache.Set("reset:email:"+email, token, 90)
		cache.Set("auth:reset:"+token, email, 5*60)
		go utils.SendResetPasswordMail(email, token)
		response.Success(c, 200, token)
	} else {
		response.Error(c, 5000, "请稍后再试")
	}
}

func ResetPassword(c *gin.Context) {
	token := c.PostForm("token")
	password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	captchaId := c.PostForm("captcha_id")

	resetForm := ResetPasswordForm{
		Token:    token,
		Password: password,
		Captcha:  captcha,
	}
	err := resetForm.ValidateResetPasswordForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	err1 := VerifyCaptcha(captchaId, captcha)
	if !err1 {
		response.Error(c, 4000, "验证码有误")
		return
	}
	email := cache.Get("auth:reset:" + token)
	result := OnResetPasswordApi(email, password)
	if result.Error == "" && result.Code == 200 {
		cache.Del("auth:reset:" + token)
		userId := strconv.FormatInt(c.GetInt64("userId"), 10)
		token := utils.GetTokenFromUser(c.GetInt64("userId"))
		utils.DelTokenToUser(userId)
		utils.DelUserToToken(token)
		response.Success(c, 200, "")
	} else {
		response.Error(c, 4000, "授权码过期")
	}
}
