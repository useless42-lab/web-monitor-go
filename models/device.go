package models

import (
	"strconv"
	"time"
)

func FilterDevice(deviceType int) string {
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
	return device
}

func FilterDeviceLog(deviceType int) string {
	var deviceLog string
	if deviceType == 1 {
		deviceLog = "web_log"
	}
	if deviceType == 2 {
		deviceLog = "server_log"
	}
	if deviceType == 3 {
		deviceLog = "api_log"
	}
	if deviceType == 4 {
		deviceLog = "tcp_log"
	}
	if deviceType == 5 {
		deviceLog = "dns_log"
	}
	if deviceType == 6 {
		deviceLog = "heartbeat_log"
	}
	if deviceType == 7 {
		deviceLog = "steam_game_server_log"
	}
	if deviceType == 8 {
		deviceLog = "minecraft_server_log"
	}
	return deviceLog
}

func FilterDeviceId(deviceType int) string {
	var deviceId string
	if deviceType == 1 {
		deviceId = "web_id"
	}
	if deviceType == 2 {
		deviceId = "server_id"
	}
	if deviceType == 3 {
		deviceId = "api_id"
	}
	if deviceType == 4 {
		deviceId = "tcp_id"
	}
	if deviceType == 5 {
		deviceId = "dns_id"
	}
	if deviceType == 6 {
		deviceId = "heartbeat_id"
	}
	if deviceType == 7 {
		deviceId = "steam_game_server_id"
	}
	if deviceType == 8 {
		deviceId = "minecraft_server_id"
	}
	return deviceId
}

type SimpleDeviceDetailItem struct {
	Id            string `json:"id" gorm:"column:id"`
	Name          string `json:"name" gorm:"column:name"`
	Path          string `json:"path" gorm:"column:path"`
	Token         string `json:"token" gorm:"column:token"`
	PolicyId      string `json:"policy_id" gorm:"column:policy_id"`
	BasicUser     string `json:"basic_user" gorm:"column:basic_user"`
	BasicPassword string `json:"basic_password" gorm:"column:basic_password"`
	DnsType       int    `json:"dns_type" gorm:"column:dns_type"`
	DnsServer     string `json:"dns_server" gorm:"column:dns_server"`
	MonitorRegion string `json:"monitor_region" gorm:"column:monitor_region"`
}

/*
获取设备详情
*/
func GetDeviceDetail(deviceId int64, deviceType int) SimpleDeviceDetailItem {
	var result SimpleDeviceDetailItem
	device := FilterDevice(deviceType)
	sqlStr := `select * from ` + device + ` where id=@deviceId and deleted_at is null and (status=1 or status=3)`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
	}).Scan(&result)
	return result
}

/*
获取某一设备的数量
包括可用和暂停的
*/
func GetDeviceCount(userId int64, groupId int64, deviceType int) RPTotal {
	var result RPTotal
	device := FilterDevice(deviceType)
	sqlStr := `select count(id) as total from ` + device + ` where user_id=@userId and group_id=@groupId and deleted_at is null and  (status=1 or status=3)`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":  userId,
		"groupId": groupId,
	}).Scan(&result)
	return result
}

/*
归档某个设备
*/
func FileDevice(id int64, deviceType int) {
	device := FilterDevice(deviceType)
	sqlStr := `update ` + device + ` set status=2 where id=@id`
	DB.Exec(sqlStr, map[string]interface{}{
		"id": id,
	})
}

/*
获取归档设备数量
*/
func GetFileCount(userId int64, deviceType int) RPTotal {
	device := FilterDevice(deviceType)
	var result RPTotal
	sqlStr := `SELECT
	count(al.id) as total
FROM
	(
		SELECT
			*
		FROM
			` + device + `
		WHERE
			group_id IN (
				SELECT
					id
				FROM
					device_group
				WHERE
					team_id IN (
						SELECT
							id
						FROM
							team_group
						WHERE
							user_id = @userId
					) and device_type=@deviceType
			) 
	) AS al
WHERE
	al.status = 2
AND al.deleted_at IS NULL`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId, "deviceType": deviceType}).Scan(&result)
	return result
}

/*
暂停监控设备
*/
func PauseMonitor(id int64, deviceType int) {
	device := FilterDevice(deviceType)
	sqlStr := `update ` + device + ` set status=3 where id=@id`
	DB.Exec(sqlStr, map[string]interface{}{
		"id": id,
	})
}

/*
开始监控设备
*/
func StartMonitor(id int64, deviceType int) {
	device := FilterDevice(deviceType)
	sqlStr := `update ` + device + ` set status=1 where id=@id and status=3`
	DB.Exec(sqlStr, map[string]interface{}{
		"device": device,
		"id":     id,
	})
}

/*
获取单个团队下所有设备数量
*/
func GetAllDeviceNumber(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `SELECT
	count(id) AS total
FROM
	(
		SELECT
			id
		FROM
			web_list AS wl
		WHERE
			group_id IN (
				SELECT
					id
				FROM
					device_group
				WHERE
					team_id = @teamId
				AND deleted_at IS NULL
				AND (status=1 or status=3)
			)
		UNION ALL
			SELECT
				id
			FROM
				server_list AS sl
			WHERE
				group_id IN (
					SELECT
						id
					FROM
						device_group
					WHERE
						team_id = @teamId
					AND deleted_at IS NULL
					AND (status=1 or status=3)
				)
			UNION ALL
				SELECT
					id
				FROM
					api_list AS al
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id = @teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id
				FROM
					tcp_list AS tl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id = @teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id
				FROM
					dns_list AS dl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id = @teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
				UNION ALL
				SELECT
					id
				FROM
					heartbeat_list AS hbl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id = @teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id
				FROM
					steam_game_server_list AS sgsl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id = @teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id
				FROM
					minecraft_server_list AS msl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id = @teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
	) AS t`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId}).Scan(&result)
	return result
}

/*
删除某个设备
*/
func DeleteDevice(id int64, deviceType int) {
	device := FilterDevice(deviceType)
	sqlStr := `
	update ` + device + ` set deleted_at=@deletedAt where id=@id
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"id":        id,
		"deletedAt": time.Now(),
	})
}

type RSimpleDeviceItem struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Path       string          `json:"path"`
	DeviceType int             `json:"device_type"`
	Log        []DeviceLogItem `json:"log" gorm:"-"`
}

/*
获取设备基础信息
*/
func GetSimpleDeviceItem(id string, deviceType int) RSimpleDeviceItem {
	device := FilterDevice(deviceType)
	var result RSimpleDeviceItem
	sqlStr := `select * from ` + device + ` where id=@id`
	DB.Raw(sqlStr, map[string]interface{}{
		"id": id,
	}).Scan(&result)
	result.DeviceType = deviceType

	log := GetStatusPageDeviceSimpleLog(id, deviceType)
	result.Log = log
	return result
}

/*
获取团队下所有设备列表
*/
func GetAllDeviceListByTeamId(teamId int64) []RSimpleDeviceItem {
	var result []RSimpleDeviceItem
	sqlStr := `
	SELECT
			id,name,1 as device_type
		FROM
			web_list AS wl
		WHERE
			group_id IN (
				SELECT
					id
				FROM
					device_group
				WHERE
					team_id =@teamId
				AND deleted_at IS NULL
				AND (status=1 or status=3)
			)
		UNION ALL
			SELECT
				id,name,2 as device_type
			FROM
				server_list AS sl
			WHERE
				group_id IN (
					SELECT
						id
					FROM
						device_group
					WHERE
						team_id =@teamId
					AND deleted_at IS NULL
					AND (status=1 or status=3)
				)
			UNION ALL
				SELECT
					id,name,3 as device_type
				FROM
					api_list AS al
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id =@teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,4 as device_type
				FROM
					tcp_list AS tl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
						team_id =@teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,5 as device_type
				FROM
					dns_list AS dl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
						team_id =@teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,6 as device_type
				FROM
					heartbeat_list AS hbl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
						team_id =@teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,7 as device_type
				FROM
					steam_game_server_list AS sgsl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
						team_id =@teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,8 as device_type
				FROM
					minecraft_server_list AS msl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
						team_id =@teamId
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

func GetAllTargetTypeDeviceListByUserId(userId int64, deviceType int) []RSimpleDeviceItem {
	var result []RSimpleDeviceItem
	device := FilterDevice(deviceType)
	sqlStr := `select id,name,path,` + strconv.Itoa(deviceType) + ` as device_type  from ` + device + ` WHERE
	group_id IN (
		SELECT
			id
		FROM
			device_group
		WHERE
			team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
		AND deleted_at IS NULL
		AND (status=1 or status=3)
	)`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

/*
获取用户所在团队所有设备列表
*/
func GetAllDeviceListByUserId(userId int64) []RSimpleDeviceItem {
	var result []RSimpleDeviceItem
	sqlStr := `
	SELECT
			id,name,1 as device_type,path
		FROM
			web_list AS wl
		WHERE
			group_id IN (
				SELECT
					id
				FROM
					device_group
				WHERE
					team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
				AND deleted_at IS NULL
				AND (status=1 or status=3)
			)
		UNION ALL
			SELECT
				id,name,2 as device_type,path
			FROM
				server_list AS sl
			WHERE
				group_id IN (
					SELECT
						id
					FROM
						device_group
					WHERE
						team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
					AND deleted_at IS NULL
					AND (status=1 or status=3)
				)
			UNION ALL
				SELECT
					id,name,3 as device_type,path
				FROM
					api_list AS al
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,4 as device_type,path
				FROM
					tcp_list AS tl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,5 as device_type,path
				FROM
					dns_list AS dl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,6 as device_type,token as path
				FROM
					heartbeat_list AS hbl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,7 as device_type,path
				FROM
					steam_game_server_list AS sgsl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
UNION ALL
				SELECT
					id,name,8 as device_type,path
				FROM
					minecraft_server_list AS msl
				WHERE
					group_id IN (
						SELECT
							id
						FROM
							device_group
						WHERE
							team_id in (select user_team.team_id from  user_team where user_id=@userId and deleted_at is null)
						AND deleted_at IS NULL
						AND (status=1 or status=3)
					)
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

/*
获取团队下使用该策略的设备数量
*/
func GetDeviceNumberByPolicyId(policyId int64) RPTotal {
	var result RPTotal
	sqlStr := `
	SELECT
	count(id) AS total
FROM
	(
		SELECT
			id
		FROM
			web_list AS wl
		WHERE
			policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
		UNION ALL
			SELECT
				id
			FROM
				server_list AS sl
			WHERE
				policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
			UNION ALL
				SELECT
					id
				FROM
					api_list AS al
				WHERE
					policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
UNION ALL
				SELECT
					id
				FROM
					tcp_list AS tl
				WHERE
					policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
UNION ALL
				SELECT
					id
				FROM
					dns_list AS dl
				WHERE
					policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
				UNION ALL
				SELECT
					id
				FROM
					heartbeat_list AS hbl
				WHERE
					policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
UNION ALL
				SELECT
					id
				FROM
					steam_game_server_list AS sgsl
				WHERE
					policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
UNION ALL
				SELECT
					id
				FROM
					minecraft_server_list AS msl
				WHERE
					policy_id =@policyId AND deleted_at IS NULL AND (status=1 or status=3)
	) AS t
	`
	DB.Raw(sqlStr, map[string]interface{}{"policyId": policyId}).Scan(&result)
	return result
}
