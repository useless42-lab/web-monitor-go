package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddStatusPage(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.PostForm("teamId"), 10, 64)
	planConfig := models.GetPlanBaseInfoByTeamId(teamId)
	name := c.PostForm("name")
	logo := c.PostForm("logo")
	description := c.PostForm("description")
	hasPassword, _ := strconv.Atoi(c.PostForm("hasPassword"))
	if planConfig.StatusPagePassword == 0 {
		hasPassword = 0
	}
	password := c.PostForm("password")
	hasPin := c.PostForm("hasPin")
	hasPinInt, _ := strconv.Atoi(hasPin)
	pinTitle := c.PostForm("pinTitle")
	pinMessage := c.PostForm("pinMessage")
	pinColor := c.PostForm("pinColor")
	if planConfig.StatusPageCustomStyle == 0 {
		pinColor = ""
	}
	hasCopyright, _ := strconv.Atoi(c.PostForm("hasCopyright"))
	if planConfig.StatusPageCopyright == 0 {
		hasCopyright = 0
	}
	hasDomain, _ := strconv.Atoi(c.PostForm("hasDomain"))
	if planConfig.StatusPageDomain == 0 {
		hasDomain = 0
	}
	domain := c.PostForm("domain")
	copyright := c.PostForm("copyright")
	deviceGroup := c.PostForm("deviceGroup")

	err := service.AddStatusPageService(teamId, name, logo, description, hasPassword, password, hasPinInt, pinTitle, pinMessage, pinColor, hasCopyright, copyright, hasDomain, domain, deviceGroup)
	if err != "" {
		response.Error(c, 3000, err)
	} else {
		response.Success(c, 200, "")
	}
}

func GetStatusPageList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetStatusPageListService(teamId, page, pageSize)
	response.Success(c, 200, result)
}

func GetUserStatusPageDetail(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	statusPageId, _ := strconv.ParseInt(c.Query("status_page_id"), 10, 64)
	data := service.GetUserStatusPageDetail(teamId, statusPageId)
	response.Success(c, 200, data)
}

func GetStatusPageDetail(c *gin.Context) {
	statusPageId, _ := strconv.ParseInt(c.Query("token"), 10, 64)
	password := c.Query("password")
	domain := c.Query("domain")
	data := service.GetStatusPageDetail(statusPageId)
	fmt.Println(domain)
	fmt.Println(data.Domain)
	if data.HasDomain == 1 {
		if domain == data.Domain {
			if data.HasPassword == 1 {
				if password == data.Password {
					response.Success(c, 200, data)
				} else {
					response.Success(c, 404, "")
				}
			} else {
				response.Success(c, 200, data)
			}
		} else {
			response.Success(c, 404, "")
		}
	} else {
		if data.HasPassword == 1 {
			if password == data.Password {
				response.Success(c, 200, data)
			} else {
				response.Success(c, 404, "")
			}
		} else {
			response.Success(c, 200, data)
		}
	}
}

func UpdateStatusPage(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.PostForm("teamId"), 10, 64)
	statusPageId, _ := strconv.ParseInt(c.PostForm("statusPageId"), 10, 64)
	planConfig := models.GetPlanBaseInfoByTeamId(teamId)
	name := c.PostForm("name")
	logo := c.PostForm("logo")
	description := c.PostForm("description")
	hasPassword, _ := strconv.Atoi(c.PostForm("hasPassword"))
	if planConfig.StatusPagePassword == 0 {
		hasPassword = 0
	}
	password := c.PostForm("password")
	hasPin := c.PostForm("hasPin")
	hasPinInt, _ := strconv.Atoi(hasPin)
	pinTitle := c.PostForm("pinTitle")
	pinMessage := c.PostForm("pinMessage")
	pinColor := c.PostForm("pinColor")
	if planConfig.StatusPageCustomStyle == 0 {
		pinColor = ""
	}
	hasCopyright, _ := strconv.Atoi(c.PostForm("hasCopyright"))
	if planConfig.StatusPageCopyright == 0 {
		hasCopyright = 0
	}
	copyright := c.PostForm("copyright")
	hasDomain, _ := strconv.Atoi(c.PostForm("hasDomain"))
	if planConfig.StatusPageDomain == 0 {
		hasDomain = 0
	}
	domain := c.PostForm("domain")
	deviceGroup := c.PostForm("deviceGroup")

	err := service.UpdateStatusPageBaseService(statusPageId, teamId, name, logo, description, hasPassword, password, hasPinInt, pinTitle, pinMessage, pinColor, hasCopyright, copyright, hasDomain, domain, deviceGroup)
	if err != "" {
		response.Error(c, 3000, err)
	} else {
		response.Success(c, 200, "")
	}
}

func DeleteStatusPage(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	statusPageId, _ := strconv.ParseInt(c.PostForm("statusPageId"), 10, 64)
	models.DeleteStatusPage(teamId, statusPageId)
	response.Success(c, 200, "")
}
