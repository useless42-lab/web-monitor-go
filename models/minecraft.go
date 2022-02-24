package models

import "WebMonitor/tools"

type MinecraftServerItem struct {
	DefaultModel
	UserId          int64  `json:"user_id" gorm:"column:user_id"`
	Name            string `json:"name" gorm:"column:name"`
	Path            string `json:"path" gorm:"column:path"`
	GroupId         int64  `json:"group_id" gorm:"column:group_id"`
	PolicyId        int64  `json:"policy_id" gorm:"column:policy_id"`
	MonitorRegion   string `json:"monitor_region" gorm:"column:monitor_region"`
	PlatformVersion string `json:"platform_version" gorm:"column:platform_version"`
}

type RMinecraftServerItem struct {
	Id              string    `json:"id" gorm:"column:id"`
	Name            string    `json:"name" gorm:"column:name"`
	Path            string    `json:"path" gorm:"column:path"`
	GroupId         int64     `json:"group_id" gorm:"column:group_id"`
	Frequency       int       `json:"frequency" gorm:"column:frequency"`
	FailedWaitTimes int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt       LocalTime `json:"created_at" gorm:"column:created_at"`
	Status          int       `json:"status" gorm:"column:status"`
	PlatformVersion string    `json:"platform_version" gorm:"column:platform_version"`
}

func AddMinecraftServer(userId int64, name string, path string, groupId int64, policyId int64, platformVersion string, monitorRegion string) MinecraftServerItem {
	minecraftServer := MinecraftServerItem{
		DefaultModel:    DefaultModel{ID: tools.GenerateSnowflakeId()},
		Name:            name,
		UserId:          userId,
		Path:            path,
		GroupId:         groupId,
		PolicyId:        policyId,
		PlatformVersion: platformVersion,
		MonitorRegion:   monitorRegion,
	}
	DB.Table("minecraft_server_list").Create(&minecraftServer)
	return minecraftServer
}

func UpdateMinecraftServer(minecraftId int64, name string, path string, policyId int64, platformVersion string, monitorRegion string) {
	sqlStr := `
	update minecraft_server_list set name=@name,path=@path,policy_id=@policyId,platform_version=@platformVersion,monitor_region=@monitorRegion where id=@minecraftId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"minecraftId":     minecraftId,
		"name":            name,
		"path":            path,
		"policyId":        policyId,
		"platformVersion": platformVersion,
		"monitorRegion":   monitorRegion,
	})
}

func GetMinecraftServerList(userId int64, groupId int64, page int, pageSize int) []RMinecraftServerItem {
	var result []RMinecraftServerItem
	sqlStr := `select * from minecraft_server_list where user_id=@userId and group_id=@groupId and deleted_at is null and  (status=1 or status=3) order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":   userId,
		"groupId":  groupId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func GetFiledMinecraftServerList(userID int64, page int, pageSize int) []RMinecraftServerItem {
	var result []RMinecraftServerItem
	sqlStr := `SELECT
	al.id,
	al.name,
	al.path,
	al.created_at,
	al.status
FROM
	(
		SELECT
			*
		FROM
			minecraft_server_list
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
					)
			)
	) AS al
WHERE
	al.status = 2
AND al.deleted_at IS NULL
LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userID, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}
