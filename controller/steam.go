package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SteamServerForm struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	TeamId   int64  `json:"team_id"`
	GroupId  int64  `json:"group_id"`
	PolicyId int64  `json:"policy_id"`
}

func (form SteamServerForm) ValidateSteamServerForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Path, validation.Required.Error("网址不能为空")),
		validation.Field(&form.TeamId, validation.Required.Error("团队不能为空")),
		validation.Field(&form.GroupId, validation.Required.Error("分组不能为空")),
		validation.Field(&form.PolicyId, validation.Required.Error("策略不能为空")),
	)
}

func AddSteamServer(c *gin.Context) {
	userId := c.GetInt64("userId")
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	steamServerForm := SteamServerForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}
	err := steamServerForm.ValidateSteamServerForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	result := service.AddSteamServerService(userId, name, path, teamId, groupId, policyId, monitorRegion)
	if result != "" {
		response.Error(c, 3000, result)
	} else {
		response.Success(c, 200, "")
	}
}

func UpdateSteamServer(c *gin.Context) {
	steamServerId, _ := strconv.ParseInt(c.PostForm("steam_server_id"), 10, 64)
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	steamServerForm := SteamServerForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}
	err := steamServerForm.ValidateSteamServerForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	models.UpdateSteamServer(steamServerId, name, path, policyId, monitorRegion)
	response.Success(c, 200, "")
}

func GetSteamServerList(c *gin.Context) {
	userId := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetSteamServerListService(userId, groupId, page, pageSize)
	response.Success(c, 200, result)
}

func GetFiledSteamServerList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetFiledSteamServerListService(userIdInt, page, pageSize)
	response.Success(c, 200, result)
}

func FileSteamServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.FileDevice(id, 7)
	response.Success(c, 200, "")
}

func StartMonitorSteamServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.StartMonitor(id, 7)
	response.Success(c, 200, "")
}

func PauseMonitorSteamServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.PauseMonitor(id, 7)
	response.Success(c, 200, "")
}

func DeleteSteam(c *gin.Context) {
	apiId, _ := strconv.ParseInt(c.Query("steam_id"), 10, 64)
	models.DeleteDevice(apiId, 7)
	response.Success(c, 200, "")
}
