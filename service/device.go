package service

import "WebMonitor/models"

func TransferDeviceGroupService(teamId int64, groupId int64, deviceId int64, deviceType int) string {
	addStatus := CheckCanAddDevice(teamId)
	if addStatus {
		models.TransferDeviceGroup(groupId, deviceId, deviceType)
		return ""
	} else {
		return "超出数量"
	}
}
