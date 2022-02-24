package controller

import (
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMobileDeviceList(c *gin.Context) {
	userId := c.GetInt64("userId")
	deviceType, _ := strconv.Atoi(c.Query("device_type"))
	result := service.GetMobileDeviceList(userId, deviceType)
	response.Success(c, 200, result)
}
