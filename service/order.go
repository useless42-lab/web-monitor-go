package service

import "WebMonitor/models"

func GetOrderListService(userId int64, page int, pageSize int) models.PaginationData {
	data := models.GetOrderList(userId)
	total := models.GetOrderCount(userId)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}
