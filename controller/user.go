package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"

	"github.com/gin-gonic/gin"
)

func GetBaseUserInfo(c *gin.Context) {
	userId := c.GetInt64("userId")
	result := models.GetBaseUserInfo(userId)
	response.Success(c, 200, result)
}

func GetUserThirdPartyInfo(c *gin.Context) {
	userId := c.GetInt64("userId")
	result := models.GetUserThirdPartyInfo(userId)
	response.Success(c, 200, result)
}

func UpdateUserThirdPartyInfo(c *gin.Context) {
	userId := c.GetInt64("userId")
	steamApiKey := c.PostForm("steam_api_key")
	models.UpdateUserThirdPartyInfo(userId, steamApiKey)
	response.Success(c, 200, "")
}
