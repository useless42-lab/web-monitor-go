package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type MinecraftServerForm struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	TeamId   int64  `json:"team_id"`
	GroupId  int64  `json:"group_id"`
	PolicyId int64  `json:"policy_id"`
}

func (form MinecraftServerForm) ValidateMinecraftServerForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Path, validation.Required.Error("网址不能为空")),
		validation.Field(&form.TeamId, validation.Required.Error("团队不能为空")),
		validation.Field(&form.GroupId, validation.Required.Error("分组不能为空")),
		validation.Field(&form.PolicyId, validation.Required.Error("策略不能为空")),
	)
}

func AddMinecraftServer(c *gin.Context) {
	userId := c.GetInt64("userId")
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	platformVersion := c.PostForm("platform_version")
	minecraftServerForm := MinecraftServerForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}
	err := minecraftServerForm.ValidateMinecraftServerForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	result := service.AddMinecraftServerService(userId, name, path, teamId, groupId, policyId, platformVersion, monitorRegion)
	if result != "" {
		response.Error(c, 3000, result)
	} else {
		response.Success(c, 200, "")
	}
}

func UpdateMinecraftServer(c *gin.Context) {
	minecraftId, _ := strconv.ParseInt(c.PostForm("minecraft_id"), 10, 64)
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	platformVersion := c.PostForm("platform_version")
	minecraftServerForm := MinecraftServerForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}
	err := minecraftServerForm.ValidateMinecraftServerForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	models.UpdateMinecraftServer(minecraftId, name, path, policyId, platformVersion, monitorRegion)
	response.Success(c, 200, "")
}

func GetMinecraftServerList(c *gin.Context) {
	userId := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetMinecraftServerListService(userId, groupId, page, pageSize)
	response.Success(c, 200, result)
}

func GetFiledMinecraftServerList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetFiledMinecraftServerListService(userIdInt, page, pageSize)
	response.Success(c, 200, result)
}

func FileMinecraftServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.FileDevice(id, 8)
	response.Success(c, 200, "")
}

func StartMonitorMinecraftServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.StartMonitor(id, 8)
	response.Success(c, 200, "")
}

func PauseMonitorMinecraftServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.PauseMonitor(id, 8)
	response.Success(c, 200, "")
}

func DeleteMinecraft(c *gin.Context) {
	apiId, _ := strconv.ParseInt(c.Query("minecraft_id"), 10, 64)
	models.DeleteDevice(apiId, 8)
	response.Success(c, 200, "")
}
