package models

import (
	"WebMonitor/tools"
)

type SteamServerItem struct {
	DefaultModel
	UserId        int64  `json:"user_id" gorm:"column:user_id"`
	Name          string `json:"name" gorm:"column:name"`
	Path          string `json:"path" gorm:"column:path"`
	GroupId       int64  `json:"group_id" gorm:"column:group_id"`
	PolicyId      int64  `json:"policy_id" gorm:"column:policy_id"`
	MonitorRegion string `json:"monitor_region" gorm:"column:monitor_region"`
}

type RSteamServerItem struct {
	Id              string    `json:"id" gorm:"column:id"`
	Name            string    `json:"name" gorm:"column:name"`
	Path            string    `json:"path" gorm:"column:path"`
	GroupId         int64     `json:"group_id" gorm:"column:group_id"`
	Frequency       int       `json:"frequency" gorm:"column:frequency"`
	FailedWaitTimes int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt       LocalTime `json:"created_at" gorm:"column:created_at"`
	Status          int       `json:"status" gorm:"column:status"`
}

func AddSteamServer(userId int64, name string, path string, groupId int64, policyId int64, monitorRegion string) SteamServerItem {
	steamServer := SteamServerItem{
		DefaultModel:  DefaultModel{ID: tools.GenerateSnowflakeId()},
		Name:          name,
		UserId:        userId,
		Path:          path,
		GroupId:       groupId,
		PolicyId:      policyId,
		MonitorRegion: monitorRegion,
	}
	DB.Table("steam_game_server_list").Create(&steamServer)
	return steamServer
}

func UpdateSteamServer(steamServerId int64, name, path string, policyId int64, monitorRegion string) {
	sqlStr := `
	update steam_game_server_list set name=@name,path=@path,policy_id=@policyId,monitor_region=@monitorRegion where id=@steamServerId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"steamServerId": steamServerId,
		"name":          name,
		"path":          path,
		"policyId":      policyId,
		"monitorRegion": monitorRegion,
	})
}

func GetSteamServerList(userId int64, groupId int64, page int, pageSize int) []RSteamServerItem {
	var result []RSteamServerItem
	sqlStr := `select * from steam_game_server_list where user_id=@userId and group_id=@groupId and deleted_at is null and  (status=1 or status=3) order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":   userId,
		"groupId":  groupId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func GetFiledSteamServerList(userID int64, page int, pageSize int) []RSteamServerItem {
	var result []RSteamServerItem
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
			steam_game_server_list
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
