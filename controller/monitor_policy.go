package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func GetMonitorPolicyList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	result := models.GetAllMonitorPolicyList(teamId)
	response.Success(c, 200, result)
}

func GetMonitorPolicyPaginationList(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetMonitorPolicyPaginationListService(teamId, page, pageSize)
	response.Success(c, 200, result)
}

func GetMonitorPolicy(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.Query("policy_id"), 10, 64)
	result := models.GetMonitorPolicy(teamId, policyId)
	response.Success(c, 200, result)
}

func DeleteMonitorPolicy(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.Query("policy_id"), 10, 64)
	canDelete := service.CheckCanMonitorPolicyDeleteService(policyId)
	result := models.GetTeamPolicyCount(teamId)
	if result.Total > 1 {
		if canDelete {
			models.DeleteMonitorPolicy(teamId, policyId)
			response.Success(c, 200, "")
		} else {
			response.Error(c, 9000, "该策略仍被分组或设备使用中，无法删除")
		}
	} else {
		response.Error(c, 9000, "无法删除团队下最后一个监控策略")
	}
}

func UpdateMonitorPolicy(c *gin.Context) {
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	name := c.PostForm("name")
	frequency, _ := strconv.Atoi(c.PostForm("frequency"))
	webMonitorType, _ := strconv.Atoi(c.PostForm("web_monitor_type"))
	serverMonitorType, _ := strconv.Atoi(c.PostForm("server_monitor_type"))
	apiMonitorType, _ := strconv.Atoi(c.PostForm("api_monitor_type"))
	webHttpStatusCode := c.PostForm("web_http_status_code")
	apiHttpStatusCode := c.PostForm("api_http_status_code")
	serverMemory, _ := strconv.ParseFloat(c.PostForm("server_memory"), 32)
	serverDisk, _ := strconv.ParseFloat(c.PostForm("server_disk"), 32)
	serverCpu, _ := strconv.ParseFloat(c.PostForm("server_cpu"), 32)
	checkSSL, _ := strconv.Atoi(c.PostForm("check_ssl"))
	checkSSLAdvance, _ := strconv.Atoi(c.PostForm("check_ssl_advance"))
	checkWhois, _ := strconv.Atoi(c.PostForm("check_whois"))
	checkWhoisAdvance, _ := strconv.Atoi(c.PostForm("check_whois_advance"))
	failedWaitTimes, _ := strconv.Atoi(c.PostForm("failed_wait_times"))
	monitorPolicyForm := MonitorPolicyForm{
		Name: name,
	}
	err1 := monitorPolicyForm.ValidateMonitorPolicyForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	models.UpdateMonitorPolicy(policyId, teamId, name, frequency, webMonitorType, serverMonitorType, apiMonitorType, webHttpStatusCode, apiHttpStatusCode, serverMemory, serverDisk, serverCpu, checkSSL, checkSSLAdvance, checkWhois, checkWhoisAdvance, failedWaitTimes)
	response.Success(c, 200, "")
}

type MonitorPolicyForm struct {
	Name string `json:"name" gorm:"column:name"`
}

func (form MonitorPolicyForm) ValidateMonitorPolicyForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("策略名称不能为空")),
	)
}

func AddMonitorPolicy(c *gin.Context) {
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	name := c.PostForm("name")
	frequency, _ := strconv.Atoi(c.PostForm("frequency"))
	webMonitorType, _ := strconv.Atoi(c.PostForm("web_monitor_type"))
	serverMonitorType, _ := strconv.Atoi(c.PostForm("server_monitor_type"))
	apiMonitorType, _ := strconv.Atoi(c.PostForm("api_monitor_type"))
	webHttpStatusCode := c.PostForm("web_http_status_code")
	webHttpRegexpText := c.PostForm("web_http_regexp_text")
	apiHttpStatusCode := c.PostForm("api_http_status_code")
	serverMemory, _ := strconv.ParseFloat(c.PostForm("server_memory"), 32)
	serverDisk, _ := strconv.ParseFloat(c.PostForm("server_disk"), 32)
	serverCpu, _ := strconv.ParseFloat(c.PostForm("server_cpu"), 32)
	checkSSL, _ := strconv.Atoi(c.PostForm("check_ssl"))
	checkSSLAdvance, _ := strconv.Atoi(c.PostForm("check_ssl_advance"))
	checkWhois, _ := strconv.Atoi(c.PostForm("check_whois"))
	checkWhoisAdvance, _ := strconv.Atoi(c.PostForm("check_whois_advance"))
	failedWaitTimes, _ := strconv.Atoi(c.PostForm("failed_wait_times"))

	monitorPolicyForm := MonitorPolicyForm{
		Name: name,
	}
	err1 := monitorPolicyForm.ValidateMonitorPolicyForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	err := service.AddMonitorPolicyService(teamId, name, frequency, webMonitorType, serverMonitorType, apiMonitorType, webHttpStatusCode, webHttpRegexpText, apiHttpStatusCode, serverMemory, serverDisk, serverCpu, checkSSL, checkSSLAdvance, checkWhois, checkWhoisAdvance, failedWaitTimes)
	if err != "" {
		response.Error(c, 3000, err)
	} else {
		response.Success(c, 200, "")
	}
}

func GetPlanFrequencyByUserId(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	result := models.GetPlanFrequencyByUserId(userIdInt)
	response.Success(c, 200, result)
}
