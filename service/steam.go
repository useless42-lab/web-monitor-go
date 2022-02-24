package service

import "WebMonitor/models"

func GetSteamServerListService(userId int64, groupId int64, page int, pageSize int) models.PaginationData {
	data := models.GetSteamServerList(userId, groupId, page, pageSize)
	total := models.GetDeviceCount(userId, groupId, 7)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func AddSteamServerService(userId int64, name string, path string, teamId int64, groupId int64, policyId int64, monitorRegion string) string {
	addStatus := CheckCanAddDevice(teamId)
	if addStatus {
		models.AddSteamServer(userId, name, path, groupId, policyId, monitorRegion)
		return ""
	} else {
		return "设备数量超出上限"
	}
}
