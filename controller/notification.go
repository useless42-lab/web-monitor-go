package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"

	"github.com/gin-gonic/gin"
)

func UpdateNotification(c *gin.Context) {
	userId := c.GetInt64("userId")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	sms := c.PostForm("sms")
	telegram := c.PostForm("telegram")
	bark := c.PostForm("bark")
	serverChan := c.PostForm("server_chan")
	models.UpdateNotification(userId, email, phone, sms, telegram, bark, serverChan)
	response.Success(c, 200, "")
}

func GetNotification(c *gin.Context) {
	userId := c.GetInt64("userId")
	result := models.GetNotification(userId)
	response.Success(c, 200, result)
}
