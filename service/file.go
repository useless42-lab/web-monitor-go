package service

import "WebMonitor/models"

// 归档网站列表
func GetFiledWebListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledWebList(userId, page, pageSize)
	total := models.GetFileCount(userId, 1)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

// 归档服务器列表
func GetFiledServerListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledServerList(userId, page, pageSize)
	total := models.GetFileCount(userId, 2)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

// 归档接口列表
func GetFiledApiListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledApiList(userId, page, pageSize)
	total := models.GetFileCount(userId, 3)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

// 归档Tcp列表
func GetFiledTcpListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledTcpList(userId, page, pageSize)
	total := models.GetFileCount(userId, 4)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func GetFiledDnsListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledDnsList(userId, page, pageSize)
	total := models.GetFileCount(userId, 5)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

// heartbeat
func GetFiledHeartbeatListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledTcpList(userId, page, pageSize)
	total := models.GetFileCount(userId, 6)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func GetFiledSteamServerListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledTcpList(userId, page, pageSize)
	total := models.GetFileCount(userId, 7)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func GetFiledMinecraftServerListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetFiledTcpList(userId, page, pageSize)
	total := models.GetFileCount(userId, 8)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}
