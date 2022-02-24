package models

import (
	"WebMonitor/tools"
)

type WebItem struct {
	DefaultModel
	UserId        int64  `json:"user_id" gorm:"column:user_id"`
	Name          string `json:"name" gorm:"column:name"`
	Path          string `json:"path" gorm:"column:path"`
	GroupId       int64  `json:"group_id" gorm:"column:group_id"`
	PolicyId      int64  `json:"policy_id" gorm:"column:policy_id"`
	BasicUser     string `json:"basic_user" gorm:"column:basic_user"`
	BasicPassword string `json:"basic_password" gorm:"column:basic_password"`
	MonitorRegion string `json:"monitor_region" gorm:"column:monitor_region"`
}

type RWebItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	Name              string    `json:"name" gorm:"column:name"`
	Path              string    `json:"path" gorm:"column:path"`
	GroupId           string    `json:"group_id" gorm:"column:group_id"`
	Frequency         int       `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int       `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int       `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int       `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode int       `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	WebHttpRegexpText string    `json:"web_http_regexp_text" gorm:"column:web_http_regexp_text"`
	ApiHttpStatusCode string    `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64   `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64   `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64   `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int       `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int       `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	FailedWaitTimes   int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt         LocalTime `json:"created_at" gorm:"column:created_at"`
	Status            int       `json:"status" gorm:"column:status"`
	BasicUser         string    `json:"basic_user" gorm:"column:basic_user"`
	BasicPassword     string    `json:"basic_password" gorm:"column:basic_password"`
}

func AddWeb(userId int64, name string, path string, groupId int64, policyId int64, basicUser string, basicPassword string, monitorRegion string) WebItem {
	web := WebItem{
		DefaultModel:  DefaultModel{ID: tools.GenerateSnowflakeId()},
		Name:          name,
		UserId:        userId,
		Path:          path,
		GroupId:       groupId,
		PolicyId:      policyId,
		BasicUser:     basicUser,
		BasicPassword: basicPassword,
		MonitorRegion: monitorRegion,
	}
	DB.Table("web_list").Create(&web)
	return web
}

/*
更新网站基础信息
*/
func UpdateWeb(webId int64, name string, path string, policyId int64, basicUser string, basicPassword string, monitorRegion string) {
	sqlStr := `
	update web_list set name=@name,path=@path,policy_id=@policyId,basic_user=@basicUser,basic_password=@basicPassword,monitor_region=@monitorRegion where id=@webId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"webId":         webId,
		"name":          name,
		"path":          path,
		"policyId":      policyId,
		"basicUser":     basicUser,
		"basicPassword": basicPassword,
		"monitorRegion": monitorRegion,
	})
}

/*
转移设备分组
*/
func TransferDeviceGroup(groupId int64, deviceId int64, deviceType int) {
	device := FilterDevice(deviceType)
	sqlStr := `update ` + device + ` set group_id=@groupId where id=@deviceId `
	DB.Exec(sqlStr, map[string]interface{}{
		"groupId":  groupId,
		"deviceId": deviceId,
	})
}

func GetWebList(userId int64, groupId int64, page int, pageSize int) []RWebItem {
	var result []RWebItem
	sqlStr := `select * from web_list where user_id=@userId and group_id=@groupId and deleted_at is null and  (status=1 or status=3) order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":   userId,
		"groupId":  groupId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

type RWebOverviewItem struct {
	Path           string    `json:"path" gorm:"column:path"`
	MaxElapsed     string    `json:"max_elapsed" gorm:"column:max_elapsed"`
	MinElapsed     string    `json:"min_elapsed" gorm:"column:min_elapsed"`
	AvgElapsed     string    `json:"avg_elapsed" gorm:"column:avg_elapsed"`
	SuccessPercent string    `json:"success_percent" gorm:"column:success_percent"`
	FailPercent    string    `json:"fail_percent" gorm:"column:fail_percent"`
	CreatedAt      LocalTime `json:"created_at" gorm:"column:created_at"`
}

func GetWebOverView(webId int64) RWebOverviewItem {
	var result RWebOverviewItem
	sqlStr := `
	SELECT
	path,
	max_elapsed,
	min_elapsed,
	avg_elapsed,
	round(countSuccess / total * 100, 2) AS success_percent,
	round(countFail / total * 100, 2) AS fail_percent,
	created_at
FROM
	(
		SELECT
			max(elapsed) AS max_elapsed,
			min(elapsed) AS min_elapsed,
			avg(elapsed) AS avg_elapsed,
			count(id) AS total
		FROM
			web_log
		WHERE
			web_id = @webId
	) AS a,
	(
		SELECT
			count(id) AS countFail
		FROM
			web_log
		WHERE
			status_code IS NULL
		AND web_id = @webId
	) AS b,
	(
		SELECT
			count(id) AS countSuccess
		FROM
			web_log
		WHERE
			status_code IS NOT NULL
		AND web_id = @webId
	) AS c,
	(
		SELECT
			created_at
		FROM
			web_log
		WHERE
			web_id = @webId
		ORDER BY
			id DESC
		LIMIT 1
	) AS d,
	(
		select path from web_list where id=@webId
	) as e
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"webId": webId,
	}).Scan(&result)
	return result
}

func GetFiledWebList(userID int64, page int, pageSize int) []RWebItem {
	var result []RWebItem
	sqlStr := `SELECT
	wl.id,
	wl.name,
	wl.path,
	wl.created_at,
	wl.status
FROM
	(
		SELECT
			*
		FROM
			web_list
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
	) AS wl
WHERE
	wl.status = 2
AND wl.deleted_at IS NULL
LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userID, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}

func GetWebDetail(id int64) WebItem {
	var result WebItem
	sqlStr := `select * from web_list where id=@id`
	DB.Raw(sqlStr, map[string]interface{}{
		"id": id,
	}).Scan(&result)
	return result
}
