package models

import (
	"WebMonitor/tools"
	"time"
)

type StatusPageBaseItem struct {
	DefaultModel
	TeamId       int64  `json:"team_id" gorm:"column:team_id"`
	Name         string `json:"name" gorm:"column:name"`
	Logo         string `json:"logo" gorm:"column:logo"`
	Description  string `json:"description" gorm:"column:description"`
	HasPassword  int    `json:"has_password" gorm:"column:has_password"`
	Password     string `json:"password" gorm:"column:password"`
	HasPin       int    `json:"has_pin" gorm:"column:has_pin"`
	PinTitle     string `json:"pin_title" gorm:"column:pin_title"`
	PinMessage   string `json:"pin_message" gorm:"column:pin_message"`
	PinColor     string `json:"pin_color" gorm:"column:pin_color"`
	HasCopyright int    `json:"has_copyright" gorm:"has_copyright"`
	Copyright    string `json:"copyright" gorm:"copyright"`
	HasDomain    int    `json:"has_domain" gorm:"column:has_domain"`
	Domain       string `json:"domain" gorm:"column:domain"`
}

func AddStatusPageBase(teamId int64, name string, logo string, description string, hasPassword int, password string, hasPin int, pinTitle string, pinMessage string, pinColor string, hasCopyright int, copyright string, hasDomain int, domain string) StatusPageBaseItem {
	data := StatusPageBaseItem{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		TeamId:       teamId,
		Name:         name,
		Logo:         logo,
		Description:  description,
		HasPassword:  hasPassword,
		Password:     password,
		HasPin:       hasPin,
		PinTitle:     pinTitle,
		PinMessage:   pinMessage,
		PinColor:     pinColor,
		HasCopyright: hasCopyright,
		Copyright:    copyright,
		HasDomain:    hasDomain,
		Domain:       domain,
	}
	DB.Table("status_page").Create(&data)
	return data
}

func UpdateStatusPageBase(statusPageId int64, teamId int64, name string, logo string, description string, hasPassword int, password string, hasPin int, pinTitle string, pinMessage string, pinColor string, hasCopyright int, copyright string, hasDomain int, domain string) {
	sqlStr := `
	update status_page set name=@name,logo=@logo,description=@description,has_password=@hasPassword,password=@password,has_pin=@hasPin,pin_title=@pinTitle,pin_message=@pinMessage,pin_color=@pinColor,has_copyright=@hasCopyright,copyright=@copyright,has_domain=@hasDomain,domain=@domain where id=@statusPageId and team_id=@teamId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"statusPageId": statusPageId,
		"teamId":       teamId,
		"name":         name,
		"logo":         logo,
		"description":  description,
		"hasPassword":  hasPassword,
		"password":     password,
		"hasPin":       hasPin,
		"pinTitle":     pinTitle,
		"pinMessage":   pinMessage,
		"pinColor":     pinColor,
		"hasCopyright": hasCopyright,
		"copyright":    copyright,
		"hasDomain":    hasDomain,
		"domain":       domain,
	})
}

type StatusPageGroupItem struct {
	DefaultModel
	StatusPageId int64  `json:"status_page_id" gorm:"column:status_page_id"`
	Name         string `json:"name" gorm:"name"`
}

func AddStatusPageGroup(statusPageId int64, name string) StatusPageGroupItem {
	data := StatusPageGroupItem{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		StatusPageId: statusPageId,
		Name:         name,
	}
	DB.Table("status_page_group").Create(&data)
	return data
}

type StatusPageDeviceItem struct {
	DefaultModel
	StatusPageGroupId int64  `json:"status_page_group_id" gorm:"status_page_group_id"`
	DeviceId          string `json:"device_id" gorm:"column:device_id"`
	DeviceType        int    `json:"device_type" gorm:"column:device_type"`
}

func AddStatusPageDevice(statusPageGroupId int64, deviceId string, deviceType int) StatusPageDeviceItem {
	data := StatusPageDeviceItem{
		DefaultModel:      DefaultModel{ID: tools.GenerateSnowflakeId()},
		StatusPageGroupId: statusPageGroupId,
		DeviceId:          deviceId,
		DeviceType:        deviceType,
	}
	DB.Table("status_page_device").Create(&data)
	return data
}

func GetStatusPageCount(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from status_page where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

type RStatusPageBaseItem struct {
	Id          string    `json:"id" gorm:"column:id"`
	TeamId      string    `json:"team_id" gorm:"column:team_id"`
	Name        string    `json:"name" gorm:"column:name"`
	Logo        string    `json:"logo" gorm:"column:logo"`
	Description string    `json:"description" gorm:"column:description"`
	HasPassword int       `json:"has_password" gorm:"column:has_password"`
	Password    string    `json:"password" gorm:"column:password"`
	HasPin      int       `json:"has_pin" gorm:"column:has_pin"`
	PinTitle    string    `json:"pin_title" gorm:"column:pin_title"`
	PinMessage  string    `json:"pin_message" gorm:"column:pin_message"`
	HasDomain   int       `json:"has_domain" gorm:"column:has_domain"`
	Domain      string    `json:"domain" gorm:"column:domain"`
	CreatedAt   LocalTime `json:"created_at" gorm:"column:created_at"`
}

func GetStatusPageList(teamId int64, page int, pageSize int) []RStatusPageBaseItem {
	var result []RStatusPageBaseItem
	sqlStr := `select * from status_page where team_id =@teamId  and deleted_at is null order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId":   teamId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

type RStatusPageItem struct {
	Id           string                 `json:"id" gorm:"column:id"`
	TeamId       string                 `json:"teamId" gorm:"column:team_id"`
	Name         string                 `json:"name" gorm:"column:name"`
	Logo         string                 `json:"logo" gorm:"column:logo"`
	Description  string                 `json:"description" gorm:"column:description"`
	HasPassword  int                    `json:"hasPassword" gorm:"column:has_password"`
	Password     string                 `json:"password" gorm:"column:"password"`
	HasPin       int                    `json:"hasPin" gorm:"column:has_pin"`
	PinTitle     string                 `json:"pinTitle" gorm:"column:pin_title"`
	PinMessage   string                 `json:"pinMessage" gorm:"column:pin_message"`
	PinColor     string                 `json:"pinColor" gorm:"column:pin_color"`
	HasCopyright int                    `json:"hasCopyright" gorm:"has_copyright"`
	Copyright    string                 `json:"copyright" gorm:"copyright"`
	HasDomain    int                    `json:"hasDomain" gorm:"column:has_domain"`
	Domain       string                 `json:"domain" gorm:"column:domain"`
	DeviceGroup  []RStatusPageGroupItem `json:"deviceGroup" gorm:"-"`
}

func GetUserStatusPageItem(teamId int64, statusPageId int64) RStatusPageItem {
	var result RStatusPageItem
	sqlStr := `select * from  status_page where id=@statusPageId and team_id=@teamId`
	DB.Raw(sqlStr, map[string]interface{}{
		"statusPageId": statusPageId,
		"teamId":       teamId,
	}).Scan(&result)
	return result
}

func GetStatusPageItem(statusPageId int64) RStatusPageItem {
	var result RStatusPageItem
	sqlStr := `select * from  status_page where id=@statusPageId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"statusPageId": statusPageId,
	}).Scan(&result)
	return result
}

type RStatusPageGroupItem struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Devices []RDeviceItem `json:"devices" gorm:"-"`
}

type RDeviceItem struct {
	Device RSimpleDeviceItem `json:"device" gorm:"-"`
}

func GetStatusPageGroup(statusPageId int64) []RStatusPageGroupItem {
	var result []RStatusPageGroupItem
	sqlStr := `select name,id from status_page_group where status_page_id=@statusPageId`
	DB.Raw(sqlStr, map[string]interface{}{
		"statusPageId": statusPageId,
	}).Scan(&result)
	return result
}

func GetStatusPageDevice(statusPageGroupId int64) []RDeviceItem {
	var data []RSimpleDeviceItem
	var result []RDeviceItem
	sqlStr := `select device_id as id,device_type from status_page_device where status_page_group_id=@statusPageGroupId`
	DB.Raw(sqlStr, map[string]interface{}{
		"statusPageGroupId": statusPageGroupId,
	}).Scan(&data)
	for i := 0; i < len(data); i++ {
		if data[i].DeviceType == 0 {
			data[i].DeviceType = 1
		}
		result = append(result, RDeviceItem{
			Device: GetSimpleDeviceItem(data[i].Id, data[i].DeviceType),
		})
		// result = append(result, GetSimpleDeviceItem(data[i].Id, data[i].DeviceType))
	}
	return result
}

type DeviceLogItem struct {
	CheckSuccess int       `json:"check_success" gorm:"column:check_success"`
	CreatedAt    LocalTime `json:"created_at" gorm:"column:created_at"`
}

func GetStatusPageDeviceSimpleLog(deviceId string, deviceType int) []DeviceLogItem {
	var result []DeviceLogItem
	deviceLog := FilterDeviceLog(deviceType)
	deviceIdColumn := FilterDeviceId(deviceType)
	sqlStr := `select check_success,created_at from ` + deviceLog + ` where ` + deviceIdColumn + `=@deviceId order by id desc limit 30`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
	}).Scan(&result)
	var targetLength int = 30
	var lengthResult int
	lengthResult = targetLength - len(result)
	if len(result) < targetLength {
		tmpData := DeviceLogItem{
			CheckSuccess: 2,
		}
		for i := 0; i < lengthResult; i++ {
			result = append(result, tmpData)
		}
	}
	return result
}

/*
危险操作
硬删除状态页分组
*/
func HardDeleteStatusPageGroup(statusPageGroupId int64) {
	sqlStr := `delete from status_page_group where id=@statusPageGroupId`
	DB.Exec(sqlStr, map[string]interface{}{
		"statusPageGroupId": statusPageGroupId,
	})
}

/*
危险操作
硬删除状态页分组设备
*/
func HardDeleteStatusPageDevice(statusPageGroupId int64) {
	sqlStr := `delete from status_page_device where status_page_group_id=@statusPageGroupId`
	DB.Exec(sqlStr, map[string]interface{}{
		"statusPageGroupId": statusPageGroupId,
	})
}

func DeleteStatusPage(teamId int64, statusPageId int64) {
	sqlStr := `update status_page set deleted_at=@deletedAt where id=@statusPageId and team_id=@teamId`
	DB.Exec(sqlStr, map[string]interface{}{
		"teamId":       teamId,
		"deletedAt":    time.Now(),
		"statusPageId": statusPageId,
	})
}
