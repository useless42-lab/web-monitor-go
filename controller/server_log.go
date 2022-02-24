package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetServerLog(c *gin.Context) {
	serverId, _ := strconv.ParseInt(c.Query("server_id"), 10, 64)
	result := models.GetServerLog(serverId)
	response.Success(c, 200, result)
}
