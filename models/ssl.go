package models

import (
	"WebMonitor/tools"
	"time"
)

type SSLItem struct {
	DefaultModel
	WebId     int64     `json:"web_id" gorm:"column:web_id"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time"`
	Subject   string    `json:"subject" gorm:"column:subject"`
	Issuer    string    `json:"issuer" gorm:"column:issuer"`
}

func AddSSLConfig(webId int64, startTime time.Time, endTime time.Time, subject string, issuer string) {
	sslConfig := SSLItem{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		WebId:        webId,
		StartTime:    startTime,
		EndTime:      endTime,
		Subject:      subject,
		Issuer:       issuer,
	}
	DB.Table("ssl_config").Create(&sslConfig)
}

func CheckSSLConfig(webId int64, startTime time.Time, endTime time.Time, subject string, issuer string) {
	sslConfig := SSLItem{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		WebId:        webId,
		StartTime:    startTime,
		EndTime:      endTime,
		Subject:      subject,
		Issuer:       issuer,
	}
	updateSSLConfig := SSLItem{
		WebId:     webId,
		StartTime: startTime,
		EndTime:   endTime,
		Subject:   subject,
		Issuer:    issuer,
	}
	err := DB.Table("ssl_config").Where("web_id = ?", &webId).First(&updateSSLConfig).Error
	if err != nil {
		DB.Table("ssl_config").Create(&sslConfig)
	} else {
		DB.Table("ssl_config").Where("web_id=?", webId).Updates(&updateSSLConfig)
	}
}

type RSSLItem struct {
	WebId     string    `json:"web_id" gorm:"column:web_id"`
	StartTime LocalTime `json:"start_time" gorm:"column:start_time"`
	EndTime   LocalTime `json:"end_time" gorm:"column:end_time"`
	Subject   string    `json:"subject" gorm:"column:subject"`
	Issuer    string    `json:"issuer" gorm:"column:issuer"`
	TEndTime  time.Time `json:"t_end_time" gorm:"column:end_time"`
}

func GetSslConfig(webId int64) RSSLItem {
	var result RSSLItem
	sqlStr := `select * from ssl_config where web_id=@webId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"webId": webId,
	}).Scan(&result)
	return result
}
