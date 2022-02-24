package models

import "WebMonitor/tools"

type ApiItem struct {
	DefaultModel
	UserId         int64  `json:"user_id" gorm:"column:user_id"`
	Name           string `json:"name" gorm:"name"`
	Path           string `json:"path" gorm:"path"`
	GroupId        int64  `json:"group_id" gorm:"column:group_id"`
	PolicyId       int64  `json:"policy_id" gorm:"column:policy_id"`
	Method         int64  `json:"method" gorm:"method"`
	RequestHeaders string `json:"request_headers" gorm:"request_headers"`
	BodyType       int64  `json:"body_type" gorm:"body_type"`
	BodyRaw        string `json:"body_raw" gorm:"body_raw"`
	BodyJson       string `json:"body_json" gorm:"body_json"`
	BodyForm       string `json:"body_form" gorm:"body_form"`
	ResponseData   string `json:"response_data" gorm:"response_data"`
	BasicUser      string `json:"basic_user" gorm:"column:basic_user"`
	BasicPassword  string `json:"basic_password" gorm:"column:basic_password"`
	MonitorRegion  string `json:"monitor_region" gorm:"column:monitor_region"`
}

type RApiItem struct {
	Id                string    `json:"id" gorm:"column:id"`
	Name              string    `json:"name" gorm:"column:name"`
	Path              string    `json:"path" gorm:"column:path"`
	GroupId           int64     `json:"group_id" gorm:"column:group_id"`
	PolicyId          int64     `json:"policy_id" gorm:"column:policy_id"`
	Method            int       `json:"method" gorm:"column:method"`
	RequestHeaders    string    `json:"request_headers" gorm:"column:request_headers"`
	BodyType          int       `json:"body_type" gorm:"column:body_type"`
	BodyRaw           string    `json:"body_raw" gorm:"column:body_raw"`
	BodyJson          string    `json:"body_json" gorm:"column:body_json"`
	BodyForm          string    `json:"body_form" gorm:"column:body_form"`
	ResponseData      string    `json:"response_data" gorm:"column:response_data"`
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
	BasicUser         string    `json:"basic_user" gorm:"column:basic_user"`
	BasicPassword     string    `json:"basic_password" gorm:"column:basic_password"`
}

func AddApi(userId int64, name string, path string, groupId int64, policyId int64, method int64, requestHeaders string, bodyType int64, bodyRaw string, bodyJson string, bodyForm string, responseData string, basicUser string, basicPassword string, monitorRegion string) ApiItem {
	api := ApiItem{
		DefaultModel:   DefaultModel{ID: tools.GenerateSnowflakeId()},
		UserId:         userId,
		Name:           name,
		Path:           path,
		GroupId:        groupId,
		PolicyId:       policyId,
		Method:         method,
		RequestHeaders: requestHeaders,
		BodyType:       bodyType,
		BodyRaw:        bodyRaw,
		BodyJson:       bodyJson,
		BodyForm:       bodyForm,
		ResponseData:   responseData,
		BasicUser:      basicUser,
		BasicPassword:  basicPassword,
		MonitorRegion:  monitorRegion,
	}
	DB.Table("api_list").Create(&api)
	return api
}

func GetApiList(userId int64, groupId int64, page int, pageSize int) []RApiItem {
	var result []RApiItem
	sqlStr := `select * from api_list where user_id=@userId and group_id=@groupId and deleted_at is null and (status=1 or status=3) order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":   userId,
		"groupId":  groupId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func UpdateApi(id int64, name string, path string, groupId int64, policyId int64, method int64, requestHeaders string, bodyType int64, bodyRaw string, bodyJson string, bodyForm string, responseData string, basicUser string, basicPassword string, monitorRegion string) ApiItem {
	sqlStr := `update api_list set name=@name,path=@path,group_id=@groupId,policy_id=@policyId,method=@method,request_headers=@requestHeaders,body_type=@bodyType,body_raw=@bodyRaw,body_json=@bodyJson,body_form=@bodyForm,response_data=@responseData,basic_user=@basicUser,basic_password=@basicPassword,monitor_region=@monitorRegion where id=@id`
	DB.Exec(sqlStr, map[string]interface{}{
		"id":             id,
		"name":           name,
		"path":           path,
		"groupId":        groupId,
		"policyId":       policyId,
		"method":         method,
		"requestHeaders": requestHeaders,
		"bodyType":       bodyType,
		"bodyRaw":        bodyRaw,
		"bodyJson":       bodyJson,
		"bodyForm":       bodyForm,
		"responseData":   responseData,
		"basicUser":      basicUser,
		"basicPassword":  basicPassword,
		"monitorRegion":  monitorRegion,
	})
	api := ApiItem{
		Name:           name,
		Path:           path,
		GroupId:        groupId,
		PolicyId:       policyId,
		Method:         method,
		RequestHeaders: requestHeaders,
		BodyType:       bodyType,
		BodyRaw:        bodyRaw,
		BodyJson:       bodyJson,
		BodyForm:       bodyForm,
		ResponseData:   responseData,
		BasicUser:      basicUser,
		BasicPassword:  basicPassword,
		MonitorRegion:  monitorRegion,
	}
	return api
}

func GetApi(id int64) RApiItem {
	var result RApiItem
	sqlStr := `select * from api_list where id=@id`
	DB.Raw(sqlStr, map[string]interface{}{
		"id": id,
	}).Scan(&result)
	return result
}

func GetFiledApiList(userID int64, page int, pageSize int) []RApiItem {
	var result []RApiItem
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
			api_list
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
