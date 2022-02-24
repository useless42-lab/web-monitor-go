package service

import "WebMonitor/models"

func GetMobileDeviceList(userId int64, deviceType int) []models.RMobileDeviceList {
	var result []models.RMobileDeviceList
	var deviceList []models.RSimpleDeviceItem
	if deviceType == 0 {
		deviceList = models.GetAllDeviceListByUserId(userId)
	} else {
		deviceList = models.GetAllTargetTypeDeviceListByUserId(userId, deviceType)
	}
	for _, item := range deviceList {
		var result1 models.RMobileDeviceList
		result1.Name = item.Name
		result1.Path = item.Path
		result1.DeviceType = item.DeviceType
		result1.Elapsed = models.GetDeviceElapsed(item.Id, item.DeviceType)
		result1.LatestLog = models.GetDeviceLatestLog(item.Id, item.DeviceType)
		result = append(result, result1)
	}
	return result
}
