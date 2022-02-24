package models

import (
	"WebMonitor/tools"
	"time"
)

type DeviceGroup struct {
	DefaultModel
	Name       string `json:"name" gorm:"column:name"`
	TeamId     int64  `json:"team_id" gorm:"column:team_id"`
	DeviceType int    `json:"device_type" gorm:"column:device_type"`
}

func AddDeviceGroup(teamId int64, name string, deviceType int) DeviceGroup {
	deviceGroup := DeviceGroup{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		TeamId:       teamId,
		Name:         name,
		DeviceType:   deviceType,
	}
	DB.Table("device_group").Create(&deviceGroup)
	return deviceGroup
}

type RDeviceGroup struct {
	Id         string    `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	DeviceType int       `json:"device_type" gorm:"device_type"`
	CreatedAt  LocalTime `json:"created_at" gorm:"colimn:created_at"`
}

func GetDeviceGroupDeviceNumber(groupId int64, deviceType int) RPTotal {
	var result RPTotal
	var device string
	if deviceType == 1 {
		device = "web_list"
	}
	if deviceType == 2 {
		device = "server_list"
	}
	if deviceType == 3 {
		device = "api_list"
	}
	if deviceType == 4 {
		device = "tcp_list"
	}
	if deviceType == 5 {
		device = "dns_list"
	}
	if deviceType == 6 {
		device = "heartbeat_list"
	}
	if deviceType == 7 {
		device = "steam_game_server_list"
	}
	if deviceType == 8 {
		device = "minecraft_server_list"
	}

	sqlStr := `select count(*) as total from ` + device + ` where deleted_at is null and group_id=@groupId and (status=1 or status=3)`
	DB.Raw(sqlStr, map[string]interface{}{"groupId": groupId}).Scan(&result)
	return result
}

func GetAllDeviceGroupList(userId int64, teamId int64) []RDeviceGroup {
	var result []RDeviceGroup
	sqlStr := `
	SELECT
        deviceg.id,
        deviceg.name,
				deviceg.device_type,
        deviceg.created_at
FROM
        device_group deviceg
LEFT JOIN user_team usert ON usert.team_id = deviceg.team_id
WHERE
usert.deleted_at is null and
	deviceg.deleted_at IS NULL AND
	usert.team_id = @teamId
AND usert.user_id = @userId
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "userId": userId}).Scan(&result)
	return result
}

func GetDeviceListByType(userId int64, teamId int64, deviceType int) []RDeviceGroup {
	var result []RDeviceGroup
	sqlStr := `
	SELECT
        deviceg.id,
        deviceg.name,
				deviceg.device_type,
        deviceg.created_at
FROM
        device_group deviceg
LEFT JOIN user_team usert ON usert.team_id = deviceg.team_id
WHERE
usert.deleted_at is null and
	deviceg.deleted_at IS NULL AND
	usert.team_id = @teamId
AND usert.user_id = @userId
AND deviceg.device_type=@deviceType
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "userId": userId, "deviceType": deviceType}).Scan(&result)
	return result
}

func GetDeviceGroupList(userId int64, teamId int64, page int, pageSize int) []RDeviceGroup {
	var result []RDeviceGroup
	sqlStr := `
	SELECT
        deviceg.id,
        deviceg.name,
				deviceg.device_type,
        deviceg.created_at
FROM
        device_group deviceg
LEFT JOIN user_team usert ON usert.team_id = deviceg.team_id
WHERE
usert.deleted_at is null and
	deviceg.deleted_at IS NULL AND
	usert.team_id = @teamId
AND usert.user_id = @userId LIMIT @pageSize OFFSET @offset
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "userId": userId, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}

func GetDeviceGroupCount(userId int64, teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `
	SELECT
        count(deviceg.id) as total
FROM
        device_group deviceg
LEFT JOIN user_team usert ON usert.team_id = deviceg.team_id
WHERE
	deviceg.deleted_at IS NULL AND
	usert.team_id = @teamId
AND usert.user_id = @userId
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "userId": userId}).Scan(&result)
	return result
}

func GetDeviceGroupCountByDeviceType(userId int64, teamId int64, deviceType int) RPTotal {
	var result RPTotal
	sqlStr := `
	SELECT
        count(deviceg.id) as total
FROM
        device_group deviceg
LEFT JOIN user_team usert ON usert.team_id = deviceg.team_id
WHERE
	deviceg.deleted_at IS NULL AND
	usert.team_id = @teamId
AND usert.user_id = @userId
AND deviceg.device_type=@deviceType
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "userId": userId, "deviceType": deviceType}).Scan(&result)
	return result
}

func GetDeviceGroupDetail(userId int64, teamId int64, groupId int64) DeviceGroup {
	var result DeviceGroup
	sqlStr := `
	SELECT
	deviceg.id,
	deviceg.name
FROM
	device_group deviceg
LEFT JOIN user_team usert ON usert.team_id = deviceg.team_id
WHERE
	usert.team_id = @teamId
AND usert.user_id = @userId
AND deviceg.id=@groupId
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "userId": userId, "groupId": groupId}).Scan(&result)
	return result
}

func UpdateDeviceGroup(userId int64, teamId int64, groupId int64, name string) {
	sqlStr := `
	UPDATE device_group AS ug
LEFT JOIN team_group AS tg ON tg.id = ug.team_id 
SET ug.name = @name,
team_id =@teamId
WHERE
	tg.user_id = @userId 
	AND ug.id =@groupId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"userId":  userId,
		"teamId":  teamId,
		"groupId": groupId,
		"name":    name,
	})
}

func DeleteDeviceGroup(userId int64, groupId int64) {
	sqlStr := `
	UPDATE device_group AS ug
LEFT JOIN team_group AS tg ON tg.id = ug.team_id 
SET ug.deleted_at = @deletedAt
WHERE
	tg.user_id = @userId 
	AND ug.id =@groupId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"deletedAt": time.Now(),
		"userId":    userId,
		"groupId":   groupId,
	})
}

func UpdateDeviceGroupName(groupId int64, name string) {
	sqlStr := `
	update device_group set name=@name where id=@groupId
	`
	DB.Exec(sqlStr, map[string]interface{}{"groupId": groupId, "name": name})
}
