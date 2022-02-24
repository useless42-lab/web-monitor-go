package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPlanBaseInfo(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	result := models.GetPlanBaseInfoByTeamId(teamId)
	response.Success(c, 200, result)
}

func GetUserPlanDetail(c *gin.Context) {
	userId := c.GetInt64("userId")
	data := models.GetUserPlanDetail(userId)
	response.Success(c, 200, data)
}

func GetPlanBaseList(c *gin.Context) {
	userId := c.GetInt64("userId")
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	token := utils.GetRandomStr(30) + timeNow
	planList := models.GetPlanBaseList()
	utils.SetShorterToken(token, userId)
	response.Success(c, 200, planList)
}
