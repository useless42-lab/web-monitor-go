package models

import "time"

type RApiLogItem struct {
	DefaultModel
	ApiId        string    `json:"api_id" gorm:"column:api_id"`
	Status       string    `json:"status" gorm:"column:status"`
	StatusCode   int       `json:"status_code" gorm:"column:status_code"`
	Proto        string    `json:"proto" gorm:"column:proto"`
	Elapsed      int64     `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string    `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"check_success"`
	Region       int       `json:"region" gorm:"region"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
}

func GetApiLog(deviceId int64, day int, page int, pageSize int) []RApiLogItem {
	var result []RApiLogItem
	sqlStr := `select * FROM api_log where api_id=@deviceId and created_at between date_sub(now(),interval @day day) and now() order by id desc LIMIT @pageSize OFFSET @offset;`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
		"day":      day,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func GetApiLogTotal(deviceId int64, day int) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total FROM api_log where api_id=@deviceId and created_at between date_sub(now(),interval @day day) and now()`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
		"day":      day,
	}).Scan(&result)
	return result
}
