package service

import "WebMonitor/models"

func AddServerService(userId int64, name string, path string, teamId int64, groupId int64, policyId int64, token string, monitorRegion string) string {
	addStatus := CheckCanAddDevice(teamId)
	if addStatus {
		models.AddServer(userId, name, path, groupId, policyId, token, monitorRegion)
		return ""
	} else {
		return "设备数量超出上限"
	}
}

func GetServerListService(userId int64, groupId int64, page int, pageSize int) models.PaginationData {
	data := models.GetServerList(userId, groupId, page, pageSize)
	total := models.GetDeviceCount(userId, groupId, 2)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}
