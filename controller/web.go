package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type WebForm struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	TeamId   int64  `json:"team_id"`
	GroupId  int64  `json:"group_id"`
	PolicyId int64  `json:"policy_id"`
}

func (form WebForm) ValidateWebForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Path, validation.Required.Error("网址不能为空"), is.URL.Error("必须是网址")),
		validation.Field(&form.TeamId, validation.Required.Error("团队不能为空")),
		validation.Field(&form.GroupId, validation.Required.Error("分组不能为空")),
		validation.Field(&form.PolicyId, validation.Required.Error("策略不能为空")),
	)
}

func AddWeb(c *gin.Context) {
	userId := c.GetInt64("userId")
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	basicUser := c.PostForm("basic_user")
	basicPassword := c.PostForm("basic_password")
	monitorRegion := c.PostForm("monitor_region")
	webForm := WebForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}

	err := webForm.ValidateWebForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	result := service.AddWebService(userId, name, path, teamId, groupId, policyId, basicUser, basicPassword, monitorRegion)
	if result != "" {
		response.Error(c, 3000, result)
	} else {
		response.Success(c, 200, "")
	}
}

func UpdateWeb(c *gin.Context) {
	name := c.PostForm("name")
	path := c.PostForm("path")
	webId, _ := strconv.ParseInt(c.PostForm("web_id"), 10, 64)
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	basicUser := c.PostForm("basic_user")
	basicPassword := c.PostForm("basic_password")
	monitorRegion := c.PostForm("monitor_region")
	webForm := WebForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}

	err := webForm.ValidateWebForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	models.UpdateWeb(webId, name, path, policyId, basicUser, basicPassword, monitorRegion)
	response.Success(c, 200, "")
}

func GetWebList(c *gin.Context) {
	userId := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetWebListService(userId, groupId, page, pageSize)
	response.Success(c, 200, result)
}

func GetWebSSLConfig(c *gin.Context) {
	webId, _ := strconv.ParseInt(c.Query("web_id"), 10, 64)
	result := models.GetSslConfig(webId)
	response.Success(c, 200, result)
}

func GetDomainWhois(c *gin.Context) {
	webId, _ := strconv.ParseInt(c.Query("web_id"), 10, 64)
	result := models.GetDomainWhois(webId)
	response.Success(c, 200, result)
}

func GetWebOverView(c *gin.Context) {
	webId, _ := strconv.ParseInt(c.Query("web_id"), 10, 64)
	result := models.GetWebOverView(webId)
	response.Success(c, 200, result)
}

func DeleteWeb(c *gin.Context) {
	webId, _ := strconv.ParseInt(c.Query("web_id"), 10, 64)
	models.DeleteDevice(webId, 1)
	response.Success(c, 200, "")
}

func GetFiledWebList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetFiledWebListService(userIdInt, page, pageSize)
	response.Success(c, 200, result)
}

func FileWeb(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.FileDevice(id, 1)
	response.Success(c, 200, "")
}

func StartMonitorWeb(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.StartMonitor(id, 1)
	response.Success(c, 200, "")
}

func PauseMonitorWeb(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.PauseMonitor(id, 1)
	response.Success(c, 200, "")
}
