package service

import "WebMonitor/models"

func GetWebLogService(deviceId int64, page int, pageSize int) models.PaginationData {
	devicePlan := models.GetPlanBaseInfoByDeviceId(deviceId, 1)
	data := models.GetWebLog(deviceId, devicePlan.ReportTimeLimit, page, pageSize)
	total := models.GetWebLogTotal(deviceId, devicePlan.ReportTimeLimit)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func GetApiLogService(deviceId int64, page int, pageSize int) models.PaginationData {
	devicePlan := models.GetPlanBaseInfoByDeviceId(deviceId, 3)
	data := models.GetApiLog(deviceId, devicePlan.ReportTimeLimit, page, pageSize)
	total := models.GetApiLogTotal(deviceId, devicePlan.ReportTimeLimit)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}
