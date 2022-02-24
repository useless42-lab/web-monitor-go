package service

import "WebMonitor/models"

func GetDnsListService(userId int64, groupId int64, page int, pageSize int) models.PaginationData {
	data := models.GetDnsList(userId, groupId, page, pageSize)
	total := models.GetDeviceCount(userId, groupId, 5)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func AddDnsService(userId int64, name string, path string, dnsType int, dnsServer string, teamId int64, groupId int64, policyId int64, monitorRegion string) string {
	addStatus := CheckCanAddDevice(teamId)
	if addStatus {
		models.AddDns(userId, name, path, dnsType, dnsServer, groupId, policyId, monitorRegion)
		return ""
	} else {
		return "设备数量超出上限"
	}
}
