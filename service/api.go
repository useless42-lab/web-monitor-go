package service

import "WebMonitor/models"

func AddApiService(userId int64, name string, path string, teamId int64, groupId int64, policyId int64, method int64, requestHeaders string, bodyType int64, bodyRaw string, bodyJson string, bodyForm string, responseData string, basicUser string, basicPassword string, monitorRegion string) models.ApiItem {
	addStatus := CheckCanAddDevice(teamId)
	var result models.ApiItem
	if addStatus {
		result := models.AddApi(userId, name, path, groupId, policyId, method, requestHeaders, bodyType, bodyRaw, bodyJson, bodyForm, responseData, basicUser, basicPassword, monitorRegion)
		return result
	} else {
		return result
	}
}

func GetApiListService(userId int64, groupId int64, page int, pageSize int) models.PaginationData {
	data := models.GetApiList(userId, groupId, page, pageSize)
	total := models.GetDeviceCount(userId, groupId, 3)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}
