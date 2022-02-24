package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type HeartbeatForm struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	TeamId   int64  `json:"team_id"`
	GroupId  int64  `json:"group_id"`
	PolicyId int64  `json:"policy_id"`
}

func (form HeartbeatForm) ValidateHeartbeatForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Token, validation.Required.Error("网址不能为空")),
		validation.Field(&form.TeamId, validation.Required.Error("团队不能为空")),
		validation.Field(&form.GroupId, validation.Required.Error("分组不能为空")),
		validation.Field(&form.PolicyId, validation.Required.Error("策略不能为空")),
	)
}

func AddHeartbeat(c *gin.Context) {
	userId := c.GetInt64("userId")
	name := c.PostForm("name")
	token := c.PostForm("token")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	heartbeatForm := HeartbeatForm{
		Name:     name,
		Token:    token,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}
	err := heartbeatForm.ValidateHeartbeatForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	result := service.AddHeartbeatService(userId, name, token, teamId, groupId, policyId, monitorRegion)
	if result != "" {
		response.Error(c, 3000, result)
	} else {
		response.Success(c, 200, "")
	}
}

func UpdateHeartbeat(c *gin.Context) {
	heartbeatId, _ := strconv.ParseInt(c.PostForm("heartbeat_id"), 10, 64)
	name := c.PostForm("name")
	token := c.PostForm("token")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	heartbeatForm := HeartbeatForm{
		Name:     name,
		Token:    token,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}
	err := heartbeatForm.ValidateHeartbeatForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	models.UpdateHeartbeat(heartbeatId, name, policyId, monitorRegion)
	response.Success(c, 200, "")
}

func GetHeartbeatList(c *gin.Context) {
	userId := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetHeartbeatListService(userId, groupId, page, pageSize)
	response.Success(c, 200, result)
}

func GetFiledHeartbeatList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetFiledHeartbeatListService(userIdInt, page, pageSize)
	response.Success(c, 200, result)
}

func FileHeartbeat(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.FileDevice(id, 6)
	response.Success(c, 200, "")
}

func StartMonitorHeartbeat(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.StartMonitor(id, 6)
	response.Success(c, 200, "")
}

func PauseMonitorHeartbeat(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.PauseMonitor(id, 6)
	response.Success(c, 200, "")
}

func AddHeartbeatLog(c *gin.Context) {
	token := c.Param("token")
	responseData := c.Query("msg")
	heartbeatData := models.GetHeartbeatIdByToken(token)
	heartbeatLogItem := models.GetLatestHeartbeatLog(heartbeatData.Id)
	createdTime := heartbeatLogItem.CreatedAt.Format("2006-01-02 15:04:05")
	createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
	createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(heartbeatData.Frequency))

	if time.Now().After(createdTimeLocation) {
		var checkSuccessInt int = 1
		models.AddHeartbeatLog(heartbeatData.Id, responseData, checkSuccessInt)
		response.Success(c, 200, "")
	} else {
		response.Error(c, 900, "未到监控时间")
	}
}

func DeleteHeartbeat(c *gin.Context) {
	apiId, _ := strconv.ParseInt(c.Query("heartbeat_id"), 10, 64)
	models.DeleteDevice(apiId, 6)
	response.Success(c, 200, "")
}
