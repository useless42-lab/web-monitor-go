package utils

import (
	"WebMonitor/cache"
	"strconv"
)

// 设置token对应用户
func SetTokenToUser(userId int64, token string) {
	userIdStr := strconv.FormatInt(userId, 10)
	cache.Set("monitorauth:id:"+userIdStr, token, 60*60*24*365)
}

// 设置用户对应token
func SetUserToToken(token string, userId int64) {
	userIdStr := strconv.FormatInt(userId, 10)
	cache.Set("monitorauth:token:"+token, userIdStr, 60*60*24*365)
}

// 生成短位token用于面包多
func SetShorterToken(token string, userId int64) {
	userIdStr := strconv.FormatInt(userId, 10)
	cache.Set("monitororder:token:"+token, userIdStr, 60*60*24)
}
func GetUserShorterToken(token string) int64 {
	userIdStr := cache.Get("monitororder:token:" + token)
	var userId int64
	if userIdStr == "" {
		userId = 0
	} else {
		userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	}
	return userId
}

func DeleteUserShorterToken(token string) {
	cache.Del("monitororder:token:" + token)
}

func DelTokenToUser(userId string) {
	cache.Del("monitorauth:id:" + userId)
}
func DelUserToToken(token string) {
	cache.Del("monitorauth:token:" + token)
}

// 获取token
func GetTokenFromUser(userId int64) string {
	userIdStr := strconv.FormatInt(userId, 10)
	return cache.Get("monitorauth:id:" + userIdStr)
}

// 获取id
func GetUserFromToken(token string) int64 {
	userIdStr := cache.Get("monitorauth:token:" + token)
	var userId int64
	if userIdStr == "" {
		userId = 0
	} else {
		userId, _ = strconv.ParseInt(userIdStr, 10, 64)
	}
	return userId
}
