package models

import "time"

type DnsLogItem struct {
	DefaultModel
	DnsId        string `json:"dns_id" gorm:"column:dns_id"`
	DnsType      int    `json:"dns_type" gorm:"column:dns_type"`
	Elapsed      int64  `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int    `json:"check_success" gorm:"column:check_success"`
}

type RDnsLogItem struct {
	DnsId        string    `json:"dns_id" gorm:"column:dns_id"`
	DnsType      int       `json:"dns_type" gorm:"column:dns_type"`
	Elapsed      int64     `json:"elapsed" gorm:"column:elapsed"`
	ResponseData string    `json:"response_data" gorm:"column:response_data"`
	CheckSuccess int       `json:"check_success" gorm:"column:check_success"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}

func AddDnsLog(dnsId string, dnsType int, elapsed int64, responseData string, checkSuccess int) {
	data := DnsLogItem{
		DnsId:        dnsId,
		DnsType:      dnsType,
		Elapsed:      elapsed,
		ResponseData: responseData,
		CheckSuccess: checkSuccess,
	}
	DB.Table("dns_log").Create(&data)
}

func GetLatestDnsLog(dnsId string) RDnsLogItem {
	var result RDnsLogItem
	sqlStr := `select * from dns_log where dns_id=@dnsId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"dnsId": dnsId,
	}).Scan(&result)
	return result
}
