package models

import (
	"WebMonitor/tools"
	"time"
)

type MonitorPolicy struct {
	DefaultModel
	PolicyId          int64   `json:"policy_id" gorm:"column:id"`
	TeamId            int64   `json:"team_id" gorm:"column:team_id"`
	Name              string  `json:"name" gorm:"column:name"`
	Frequency         int     `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int     `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int     `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int     `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode string  `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	WebHttpRegexpText string  `json:"web_http_regexp_text" gorm:"column:web_http_regexp_text"`
	ApiHttpStatusCode string  `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64 `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64 `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64 `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int     `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int     `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	CheckWhois        int     `json:"check_whois" gorm:"column:check_whois"`
	CheckWhoisAdvance int     `json:"check_whois_advance" gorm:"column:check_whois_advance"`
	FailedWaitTimes   int     `json:"failed_wait_times" gorm:"column:failed_wait_times"`
}

type RMonitorPolicyDetail struct {
	Id                string  `json:"id" gorm:"column:id"`
	TeamId            string  `json:"team_id" gorm:"column:team_id"`
	Name              string  `json:"name" gorm:"column:name"`
	Frequency         int     `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int     `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int     `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int     `json:"api_monitor_type" gorm:"column:api_monitor_type"`
	WebHttpStatusCode string  `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	WebHttpRegexpText string  `json:"web_http_regexp_text" gorm:"column:web_http_regexp_text"`
	ApiHttpStatusCode string  `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64 `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64 `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64 `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int     `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int     `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	CheckWhois        int     `json:"check_whois" gorm:"column:check_whois"`
	CheckWhoisAdvance int     `json:"check_whois_advance" gorm:"column:check_whois_advance"`
	FailedWaitTimes   int     `json:"failed_wait_times" gorm:"column:failed_wait_times"`
}

type IMonitorPolicy struct {
	DefaultModel
	TeamId            int64   `json:"team_id" gorm:"column:team_id"`
	Name              string  `json:"name" gorm:"column:name"`
	Frequency         int     `json:"frequency" gorm:"column:frequency"`
	WebMonitorType    int     `json:"web_monitor_type" gorm:"column:web_monitor_type"`
	ServerMonitorType int     `json:"server_monitor_type" gorm:"column:server_monitor_type"`
	ApiMonitorType    int     `json:"api_monitor_type" gorm:"api_monitor_type"`
	WebHttpStatusCode string  `json:"web_http_status_code" gorm:"column:web_http_status_code"`
	WebHttpRegexpText string  `json:"web_http_regexp_text" gorm:"column:web_http_regexp_text"`
	ApiHttpStatusCode string  `json:"api_http_status_code" gorm:"column:api_http_status_code"`
	ServerMemory      float64 `json:"server_memory" gorm:"column:server_memory"`
	ServerDisk        float64 `json:"server_disk" gorm:"column:server_disk"`
	ServerCpu         float64 `json:"server_cpu" gorm:"column:server_cpu"`
	CheckSSL          int     `json:"check_ssl" gorm:"column:check_ssl"`
	CheckSSLAdvance   int     `json:"check_ssl_advance" gorm:"column:check_ssl_advance"`
	CheckWhois        int     `json:"check_whois" gorm:"column:check_whois"`
	CheckWhoisAdvance int     `json:"check_whois_advance" gorm:"column:check_whois_advance"`
	FailedWaitTimes   int     `json:"failed_wait_times" gorm:"column:failed_wait_times"`
}

// 添加策略
func AddMonitorPolicy(teamId int64, name string, frequency int, webMonitorType int, serverMonitorType int, apiMonitorType int, webHttpStatusCode string, webHttpRegexpText string, apiHttpStatusCode string, serverMemory float64, serverDisk float64, serverCpu float64, checkSSL int, checkSSLAdvance int, checkWhois int, checkWhoisAdvance int, failedWaitTimes int) IMonitorPolicy {
	data := IMonitorPolicy{
		DefaultModel:      DefaultModel{ID: tools.GenerateSnowflakeId()},
		TeamId:            teamId,
		Name:              name,
		Frequency:         frequency,
		WebMonitorType:    webMonitorType,
		ServerMonitorType: serverMonitorType,
		ApiMonitorType:    apiMonitorType,
		WebHttpStatusCode: webHttpStatusCode,
		WebHttpRegexpText: webHttpRegexpText,
		ApiHttpStatusCode: apiHttpStatusCode,
		ServerMemory:      serverMemory,
		ServerDisk:        serverDisk,
		ServerCpu:         serverCpu,
		CheckSSL:          checkSSL,
		CheckSSLAdvance:   checkSSLAdvance,
		CheckWhois:        checkWhois,
		CheckWhoisAdvance: checkWhoisAdvance,
		FailedWaitTimes:   failedWaitTimes,
	}
	DB.Table("monitor_policy").Create(&data)
	return data
}

type RMonitorPolicy struct {
	Id        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Frequency string    `json:"frequency" gorm:"frequency"`
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at"`
}

func GetAllMonitorPolicyList(teamId int64) []RMonitorPolicy {
	var result []RMonitorPolicy
	sqlStr := "select * from monitor_policy where team_id=@teamId and deleted_at is null"
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId}).Scan(&result)
	return result
}

func GetMonitorPolicyList(teamId int64, page int, pageSize int) []RMonitorPolicy {
	var result []RMonitorPolicy
	sqlStr := "select * from monitor_policy where team_id=@teamId and deleted_at is null LIMIT @pageSize OFFSET @offset"
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}

func GetMonitorPolicyCount(teamId int64) RPTotal {
	var total RPTotal
	sqlStr := `select count(*) as total from monitor_policy where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId}).Scan(&total)
	return total
}

func GetMonitorPolicy(teamId int64, policyId int64) RMonitorPolicyDetail {
	var result RMonitorPolicyDetail
	sqlStr := "select * from monitor_policy where team_id=@teamId and id=@policyId"
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "policyId": policyId}).Scan(&result)
	return result
}

func DeleteMonitorPolicy(teamId int64, policyId int64) {
	sqlStr := `
	update monitor_policy set deleted_at=@deletedAt where team_id=@teamId and id=@policyId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"deletedAt": time.Now(),
		"teamId":    teamId,
		"policyId":  policyId,
	})
}

func UpdateMonitorPolicy(policyId int64, teamId int64, name string, frequency int, webMonitorType int, serverMonitorType int, apiMonitorType int, webHttpStatusCode string, apiHttpStatusCode string, serverMemory float64, serverDisk float64, serverCpu float64, checkSSL int, checkSSLAdvance int, checkWhois int, checkWhoisAdvance int, failedWaitTimes int) {
	sqlStr := `
	update monitor_policy set name=@name,frequency=@frequency,web_monitor_type=@webMonitorType,server_monitor_type=@serverMonitorType,api_monitor_type=@apiMonitorType,web_http_status_code=@webHttpStatusCode,api_http_status_code=@apiHttpStatusCode,server_memory=@serverMemory,server_disk=@serverDisk,server_cpu=@serverCpu,check_ssl=@checkSSL,check_ssl_advance=@checkSSLAdvance,check_whois=@checkWhois,check_whois_advance=@checkWhoisAdvance,failed_wait_times=@failedWaitTimes where id=@policyId and team_id=@teamId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"policyId":          policyId,
		"teamId":            teamId,
		"name":              name,
		"frequency":         frequency,
		"webMonitorType":    webMonitorType,
		"serverMonitorType": serverMonitorType,
		"apiMonitorType":    apiMonitorType,
		"webHttpStatusCode": webHttpStatusCode,
		"apiHttpStatusCode": apiHttpStatusCode,
		"serverMemory":      serverMemory,
		"serverDisk":        serverDisk,
		"serverCpu":         serverCpu,
		"checkSSL":          checkSSL,
		"checkSSLAdvance":   checkSSLAdvance,
		"checkWhois":        checkWhois,
		"checkWhoisAdvance": checkWhoisAdvance,
		"failedWaitTimes":   failedWaitTimes,
	})
}

func GetDefaultMonitorPolicy(teamId int64) RMonitorPolicy {
	var result RMonitorPolicy
	sqlStr := `
	SELECT
	* 
FROM
	monitor_policy 
WHERE
	team_id = @teamId 
	LIMIT 1
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId}).Scan(&result)
	return result
}

/*
获取团队下策略数量
*/
func GetTeamPolicyCount(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from monitor_policy where team_id=@teamId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}
