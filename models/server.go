package models

import (
	"WebMonitor/tools"
)

type ServerItem struct {
	DefaultModel
	UserId        int64  `json:"user_id" gorm:"column:user_id"`
	Name          string `json:"name" gorm:"name"`
	Path          string `json:"path" gorm:"column:path"`
	GroupId       int64  `json:"group_id" gorm:"column:group_id"`
	PolicyId      int64  `json:"policy_id" gorm:"column:policy_id"`
	Token         string `json:"token" gorm:"column:token"`
	MonitorRegion string `json:"monitor_region" gorm:"column:monitor_region"`
}

type RServerItem struct {
	Id                string    `json:"id"gorm:"column:id"`
	Name              string    `json:"name" gorm:"name"`
	Path              string    `json:"path" gorm:"column:path"`
	GroupId           int64     `json:"group_id" gorm:"column:group_id"`
	Token             string    `json:"token" gorm:"column:token"`
	Frequency         int       `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int       `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int       `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int       `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode int       `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	ApiHttpStatusCode string    `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64   `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64   `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64   `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int       `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int       `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	FailedWaitTimes   int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt         LocalTime `json:"created_at" gorm:"column:created_at"`
	Status            int       `json:"status" gorm:"column:status"`
}

func AddServer(userId int64, name string, path string, groupId int64, policyId int64, token string, monitorRegion string) ServerItem {
	server := ServerItem{
		DefaultModel:  DefaultModel{ID: tools.GenerateSnowflakeId()},
		UserId:        userId,
		Name:          name,
		Path:          path,
		GroupId:       groupId,
		PolicyId:      policyId,
		Token:         token,
		MonitorRegion: monitorRegion,
	}
	DB.Table("server_list").Create(&server)
	return server
}

func UpdateServer(serverId int64, name string, path string, policyId int64, monitorRegion string) {
	sqlStr := `
	update server_list set name=@name,path=@path,policy_id=@policyId,monitor_region=@monitorRegion where id=@serverId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"name":          name,
		"path":          path,
		"serverId":      serverId,
		"policyId":      policyId,
		"monitorRegion": monitorRegion,
	})
}

func GetServerList(userId int64, groupId int64, page int, pageSize int) []RServerItem {
	var result []RServerItem
	sqlStr := `select * from server_list where user_id=@userId and group_id=@groupId and deleted_at is null and  (status=1 or status=3) order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":   userId,
		"groupId":  groupId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func IsTokenExist(token string) bool {
	var result RServerItem
	sqlStr := `select * from server_list where token=@token and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"token": token,
	}).Scan(&result)
	if result.Token != "" {
		return true
	} else {
		return false
	}
}

func RefreshServerToken(userId int64, serverId int64, token string) {
	sqlStr := `update server_list set token=@token where user_id=@userId and id=@serverId`
	DB.Exec(sqlStr, map[string]interface{}{
		"token":    token,
		"userId":   userId,
		"serverId": serverId,
	})
}

func GetServerIdByToken(token string) RServerItem {
	var result RServerItem
	sqlStr := `
	SELECT
	*
FROM
	server_list
LEFT JOIN monitor_policy ON monitor_policy.id = server_list.policy_id
WHERE
	server_list.deleted_at IS NULL AND server_list.token=@token
AND server_list.status = 1
`
	DB.Raw(sqlStr, map[string]interface{}{
		"token": token,
	}).Scan(&result)
	return result
}

func GetFiledServerList(userID int64, page int, pageSize int) []RServerItem {
	var result []RServerItem
	sqlStr := `SELECT
	sl.id,
	sl.name,
	sl.path,
	sl.created_at,
	sl.status
FROM
	(
		SELECT
			*
		FROM
			server_list
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
	) AS sl
WHERE
	sl.status = 2
AND sl.deleted_at IS NULL
LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userID, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}
