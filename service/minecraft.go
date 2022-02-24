package service

import "WebMonitor/models"

func GetMinecraftServerListService(userId int64, groupId int64, page int, pageSize int) models.PaginationData {
	data := models.GetMinecraftServerList(userId, groupId, page, pageSize)
	total := models.GetDeviceCount(userId, groupId, 8)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func AddMinecraftServerService(userId int64, name string, path string, teamId int64, groupId int64, policyId int64, platformVersion string, monitorRegion string) string {
	addStatus := CheckCanAddDevice(teamId)
	if addStatus {
		models.AddMinecraftServer(userId, name, path, groupId, policyId, platformVersion, monitorRegion)
		return ""
	} else {
		return "设备数量超出上限"
	}
}
