package controller

import (
	"WebMonitor/tools"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type AuthItemStruct struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
type ResponseAuthStruct struct {
	Code  int            `json:"code"`
	Error string         `json:"error"`
	Data  AuthItemStruct `json:"data"`
}

func OnLoginApi(email string, password string) ResponseAuthStruct {
	path := os.Getenv("AUTH_API") + `v1/auth/login`
	resp, err := http.PostForm(path, url.Values{"email": {email}, "password": {password}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var respJson ResponseAuthStruct
	json.Unmarshal([]byte(string(body)), &respJson)
	return respJson
}

func OnRegisterApi(username string, email string, password string) ResponseAuthStruct {
	path := os.Getenv("AUTH_API") + `v1/auth/register`
	resp, err := http.PostForm(path, url.Values{"username": {username}, "email": {email}, "password": {password}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var respJson ResponseAuthStruct
	json.Unmarshal([]byte(string(body)), &respJson)
	return respJson
}

func OnChangePasswordApi(userId string, oldPassword string, newPassword string, reNewPassword string) ResponseAuthStruct {
	encryptCode := tools.AesEncrypt("changepassword"+userId, os.Getenv("KEY"))
	path := os.Getenv("AUTH_API") + `v1/auth/password/change`
	resp, err := http.PostForm(path, url.Values{"user_id": {userId}, "old_password": {oldPassword}, "new_password": {newPassword}, "re_new_password": {reNewPassword}, "encrypt_code": {encryptCode}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var respJson ResponseAuthStruct
	json.Unmarshal([]byte(string(body)), &respJson)
	return respJson
}

func OnResetPasswordApi(email string, password string) ResponseAuthStruct {
	encryptCode := tools.AesEncrypt("resetpassword"+email, os.Getenv("KEY"))
	path := os.Getenv("AUTH_API") + `v1/auth/password/reset`
	resp, err := http.PostForm(path, url.Values{"email": {email}, "password": {password}, "encrypt_code": {encryptCode}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var respJson ResponseAuthStruct
	json.Unmarshal([]byte(string(body)), &respJson)
	return respJson
}
