package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func GetDeviceGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	result := models.GetAllDeviceGroupList(userIdInt, teamId)
	// result := service.GetDeviceGroupPaginationListService(userIdInt, teamId, page, pageSize)
	response.Success(c, 200, result)
}

func GetDeviceGroupType(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	deviceType, _ := strconv.Atoi(c.Query("type"))
	result := models.GetDeviceListByType(userIdInt, teamId, deviceType)
	response.Success(c, 200, result)

}

func GetDeviceGroupPaginationList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	// result := models.GetDeviceGroupList(userIdInt, teamId)
	result := service.GetDeviceGroupPaginationListService(userIdInt, teamId, page, pageSize)
	response.Success(c, 200, result)
}

func GetDeviceGroupDetail(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	result := models.GetDeviceGroupDetail(userIdInt, teamId, groupId)
	response.Success(c, 200, result)
}

func UpdateDeviceGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	name := c.PostForm("name")
	urlGroupForm := DeviceGroupForm{
		Name:   name,
		TeamId: teamId,
	}
	err1 := urlGroupForm.ValidateDeviceGroupForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	models.UpdateDeviceGroup(userIdInt, teamId, groupId, name)
	response.Success(c, 200, "")
}

func DeleteGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	deviceType, _ := strconv.Atoi(c.Query("device_type"))
	result := service.DeleteDeviceGroupService(userIdInt, teamId, groupId, deviceType)
	if result != "" {
		response.Error(c, 4000, result)
	} else {
		response.Success(c, 200, "")
	}
}

type DeviceGroupForm struct {
	Name     string `json:"name" gorm:"column:name"`
	TeamId   int64  `json:"team_id" gorm:"column:team_id"`
	PolicyId int64  `json:"policy_id" gorm:"column:policy_id"`
}

func (form DeviceGroupForm) ValidateDeviceGroupForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("分组名称不能为空")),
		validation.Field(&form.TeamId, validation.NotNil.Error("团队空间不能为空")),
		validation.Field(&form.PolicyId, validation.NotNil.Error("监控策略不能为空")),
	)
}

func AddDeviceGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	name := c.PostForm("name")
	deviceType, _ := strconv.Atoi(c.PostForm("device_type"))
	urlGroupForm := DeviceGroupForm{
		Name:   name,
		TeamId: teamId,
	}
	err1 := urlGroupForm.ValidateDeviceGroupForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	err := service.AddDeviceGroupService(userIdInt, teamId, name, deviceType)
	if err != "" {
		response.Error(c, 3000, err)
	} else {
		response.Success(c, 200, "")
	}
}

func (form DeviceGroupForm) ValidateDeviceGroupFormName() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("分组名称不能为空")),
	)
}

func UpdateDeviceGroupName(c *gin.Context) {
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	name := c.PostForm("name")
	urlGroupForm := DeviceGroupForm{
		Name: name,
	}
	err := urlGroupForm.ValidateDeviceGroupFormName()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	models.UpdateDeviceGroupName(groupId, name)
	response.Success(c, 200, "")
}

func GetNotificationList(c *gin.Context) {
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	result := models.GetNotificationList(groupId)
	response.Success(c, 200, result)
}

func GetTeamMemberListByGroupId(c *gin.Context) {
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	result := models.GetTeamMemberListByGroupId(groupId)
	response.Success(c, 200, result)
}

func AddNotificationListItem(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	notificationType, _ := strconv.Atoi(c.PostForm("notification_type"))
	deviceType, _ := strconv.Atoi(c.PostForm("device_type"))
	models.AddNotificationListItem(userId, groupId, deviceType, notificationType)
	response.Success(c, 200, "")
}

func DeleteNotificationListItem(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.DeleteNotificationListItem(id)
	response.Success(c, 200, "")
}
