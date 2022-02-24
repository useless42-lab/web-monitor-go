package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDeviceList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	result := models.GetAllDeviceListByTeamId(teamId)
	response.Success(c, 200, result)
}

func TransferDeviceGroup(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	deviceId, _ := strconv.ParseInt(c.PostForm("device_id"), 10, 64)
	deviceType, _ := strconv.Atoi(c.PostForm("device_type"))
	result := service.TransferDeviceGroupService(teamId, groupId, deviceId, deviceType)
	if result == "" {
		response.Success(c, 200, "")
	} else {
		response.Error(c, 6666, "")
	}
}

func GetDeviceDetail(c *gin.Context) {
	deviceId, _ := strconv.ParseInt(c.Query("device_id"), 10, 64)
	deviceType, _ := strconv.Atoi(c.Query("device_type"))
	result := models.GetDeviceDetail(deviceId, deviceType)
	response.Success(c, 200, result)
}
