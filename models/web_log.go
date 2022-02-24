package models

import (
	"time"
)

type WebLogItem struct {
	DefaultModel
	WebId        string `json:"web_id" gorm:"column:web_id"`
	Status       string `json:"status" gorm:"column:status"`
	StatusCode   int    `json:"status_code" gorm:"column:status_code"`
	Proto        string `json:"proto" gorm:"column:proto"`
	Elapsed      int64  `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int    `json:"check_success" gorm:"check_success"`
	Region       int    `json:"region" gorm:"region"`
}

type RWebLogItem struct {
	DefaultModel
	WebId        string    `json:"web_id" gorm:"column:web_id"`
	Status       string    `json:"status" gorm:"column:status"`
	StatusCode   int       `json:"status_code" gorm:"column:status_code"`
	Proto        string    `json:"proto" gorm:"column:proto"`
	Elapsed      int64     `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string    `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"check_success"`
	Region       int       `json:"region" gorm:"region"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
}

func AddWebLog(webId string, status string, statusCode int, proto string, elapsed int64, responseData string, checkSuccess int, region int) {
	webLog := WebLogItem{
		WebId:        webId,
		Status:       status,
		StatusCode:   statusCode,
		Proto:        proto,
		Elapsed:      elapsed,
		ResponseData: responseData,
		CheckSuccess: checkSuccess,
		Region:       region,
	}
	DB.Table("web_log").Create(&webLog)
}

func GetLatestWebLog(webId string) RWebLogItem {
	var result RWebLogItem
	sqlStr := `select * from web_log where web_id=@webId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"webId": webId,
	}).Scan(&result)
	return result
}

func GetWebLog(deviceId int64, day int, page int, pageSize int) []RWebLogItem {
	var result []RWebLogItem
	sqlStr := `select * FROM web_log where web_id=@deviceId and created_at between date_sub(now(),interval @day day) and now() order by id desc LIMIT @pageSize OFFSET @offset;`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
		"day":      day,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func GetWebLogTotal(deviceId int64, day int) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total FROM web_log where web_id=@deviceId and created_at between date_sub(now(),interval @day day) and now()`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
		"day":      day,
	}).Scan(&result)
	return result
}
