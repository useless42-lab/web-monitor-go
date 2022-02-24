package service

import (
	"WebMonitor/models"
)

func GetMonitorPolicyPaginationListService(teamId int64, page int, pageSize int) models.PaginationData {
	data := models.GetMonitorPolicyList(teamId, page, pageSize)
	total := models.GetMonitorPolicyCount(teamId)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func CheckCanMonitorPolicyDeleteService(policyId int64) bool {
	policyDeviceCount := models.GetDeviceNumberByPolicyId(policyId)
	if policyDeviceCount.Total > 0 {
		return false
	} else {
		return true
	}
}

func AddMonitorPolicyService(teamId int64, name string, frequency int, webMonitorType int, serverMonitorType int, apiMonitorType int, webHttpStatusCode string, webHttpRegexpText string, apiHttpStatusCode string, serverMemory float64, serverDisk float64, serverCpu float64, checkSSL int, checkSSLAdvance int, checkWhois int, checkWhoisAdvance int, failedWaitTimes int) string {
	plan := models.GetPlanBaseInfoByTeamId(teamId)
	monitorPolicyNumber := models.GetMonitorPolicyCount(teamId)
	if monitorPolicyNumber.Total < plan.PerTeamMonitorPolicyLimit {
		models.AddMonitorPolicy(teamId, name, frequency, webMonitorType, serverMonitorType, apiMonitorType, webHttpStatusCode, webHttpRegexpText, apiHttpStatusCode, serverMemory, serverDisk, serverCpu, checkSSL, checkSSLAdvance, checkWhois, checkWhoisAdvance, failedWaitTimes)
		return ""
	} else {
		return "监控策略数量超出上限"
	}
}
