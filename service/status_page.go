package service

import (
	"WebMonitor/models"
	"encoding/json"
	"strconv"
)

type DeviceItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	DeviceType int    `json:"device_type"`
}

type DevicesItem struct {
	Device DeviceItem `json:"device"`
}

type DeviceGroup struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Devices []DevicesItem `json:"devices"`
}

func AddStatusPageService(teamId int64, name string, logo string, description string, hasPasswordInt int, password string, hasPinInt int, pinTitle string, pinMessage string, pinColor string, hasCopyright int, copyright string, hasDomain int, domain string, deviceGroup string) string {
	plan := models.GetPlanBaseInfoByTeamId(teamId)
	statusPageNumber := models.GetStatusPageCount(teamId)
	if statusPageNumber.Total < plan.StatusPageLimit {
		var deviceGroupItem []DeviceGroup
		json.Unmarshal([]byte(deviceGroup), &deviceGroupItem)
		statusPageResult := models.AddStatusPageBase(teamId, name, logo, description, hasPasswordInt, password, hasPinInt, pinTitle, pinMessage, pinColor, hasCopyright, copyright, hasDomain, domain)
		for _, item1 := range deviceGroupItem {
			statusPageGroupResult := models.AddStatusPageGroup(statusPageResult.ID, item1.Name)
			for _, item2 := range item1.Devices {
				models.AddStatusPageDevice(statusPageGroupResult.ID, item2.Device.Id, item2.Device.DeviceType)
			}
		}
		return ""
	} else {
		return "状态页数量超出上限"
	}
}

func GetStatusPageListService(teamId int64, page int, pageSize int) models.PaginationData {
	data := models.GetStatusPageList(teamId, page, pageSize)
	total := models.GetStatusPageCount(teamId)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func GetUserStatusPageDetail(teamId int64, statusPageId int64) models.RStatusPageItem {
	statusPageData := models.GetUserStatusPageItem(teamId, statusPageId)
	statusPageGroupData := models.GetStatusPageGroup(statusPageId)
	var statusPageDeviceData []models.RDeviceItem
	for i := 0; i < len(statusPageGroupData); i++ {
		statusPageGroupId, _ := strconv.ParseInt(statusPageGroupData[i].Id, 10, 64)
		statusPageDeviceData = models.GetStatusPageDevice(statusPageGroupId)
		statusPageGroupData[i].Devices = statusPageDeviceData
	}
	data := models.RStatusPageItem{
		Id:           statusPageData.Id,
		TeamId:       statusPageData.TeamId,
		Name:         statusPageData.Name,
		Logo:         statusPageData.Logo,
		Description:  statusPageData.Description,
		HasPassword:  statusPageData.HasPassword,
		Password:     statusPageData.Password,
		HasPin:       statusPageData.HasPin,
		PinTitle:     statusPageData.PinTitle,
		PinMessage:   statusPageData.PinMessage,
		PinColor:     statusPageData.PinColor,
		HasCopyright: statusPageData.HasCopyright,
		Copyright:    statusPageData.Copyright,
		HasDomain:    statusPageData.HasDomain,
		Domain:       statusPageData.Domain,
		DeviceGroup:  statusPageGroupData,
	}
	return data
}

func GetStatusPageDetail(statusPageId int64) models.RStatusPageItem {
	statusPageData := models.GetStatusPageItem(statusPageId)
	statusPageGroupData := models.GetStatusPageGroup(statusPageId)
	var statusPageDeviceData []models.RDeviceItem
	for i := 0; i < len(statusPageGroupData); i++ {
		statusPageGroupId, _ := strconv.ParseInt(statusPageGroupData[i].Id, 10, 64)
		statusPageDeviceData = models.GetStatusPageDevice(statusPageGroupId)
		statusPageGroupData[i].Devices = statusPageDeviceData
	}
	data := models.RStatusPageItem{
		Id:           statusPageData.Id,
		Name:         statusPageData.Name,
		Logo:         statusPageData.Logo,
		Description:  statusPageData.Description,
		HasPassword:  statusPageData.HasPassword,
		Password:     statusPageData.Password,
		HasPin:       statusPageData.HasPin,
		PinTitle:     statusPageData.PinTitle,
		PinMessage:   statusPageData.PinMessage,
		PinColor:     statusPageData.PinColor,
		HasCopyright: statusPageData.HasCopyright,
		Copyright:    statusPageData.Copyright,
		HasDomain:    statusPageData.HasDomain,
		Domain:       statusPageData.Domain,
		DeviceGroup:  statusPageGroupData,
	}
	return data
}

func UpdateStatusPageBaseService(statusPageId int64, teamId int64, name string, logo string, description string, hasPassword int, password string, hasPin int, pinTitle string, pinMessage string, pinColor string, hasCopyright int, copyright string, hasDomain int, domain string, deviceGroup string) string {
	models.UpdateStatusPageBase(statusPageId, teamId, name, logo, description, hasPassword, password, hasPin, pinTitle, pinMessage, pinColor, hasCopyright, copyright, hasDomain, domain)
	statusPageGroupData := models.GetStatusPageGroup(statusPageId)
	for _, item := range statusPageGroupData {
		// 根据 状态页编号 分组编号 硬删除
		statusPageGroupId, _ := strconv.ParseInt(item.Id, 10, 64)
		models.HardDeleteStatusPageDevice(statusPageGroupId)
		models.HardDeleteStatusPageGroup(statusPageGroupId)
	}
	var deviceGroupItem []DeviceGroup
	json.Unmarshal([]byte(deviceGroup), &deviceGroupItem)
	for _, item1 := range deviceGroupItem {
		statusPageGroupResult := models.AddStatusPageGroup(statusPageId, item1.Name)
		for _, item2 := range item1.Devices {
			models.AddStatusPageDevice(statusPageGroupResult.ID, item2.Device.Id, item2.Device.DeviceType)
		}
	}
	return ""
}
