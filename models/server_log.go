package models

import "time"

type ServerLogItem struct {
	DefaultModel
	ServerId          string `json:"server_id" gorm:"column:server_id"`
	CpuUser           string `json:"cpu_user" gorm:"column:cpu_user"`
	CpuSystem         string `json:"cpu_system" gorm:"column:cpu_system"`
	CpuIdle           string `json:"cpu_idle" gorm:"column:cpu_idle"`
	CpuPercent        string `json:"cpu_percent" gorm:"column:cpu_percent"`
	MemoryTotal       string `json:"memory_total" gorm:"column:memory_total"`
	MemoryAvailable   string `json:"memory_available" gorm:"memory_available"`
	MemoryUsed        string `json:"memory_used" gorm:"column:memory_used"`
	MemoryUsedPercent string `json:"memory_used_percent" gorm:"column:memory_used_percent"`
	DiskTotal         string `json:"disk_total" gorm:"column:disk_total"`
	DiskFree          string `json:"disk_free" gorm:"column:disk_free"`
	DiskUsed          string `json:"disk_used" gorm:"column:disk_used"`
	DiskUsedPercent   string `json:"disk_used_percent" gorm:"column:disk_used_percent"`
	NetSent           string `json:"net_sent" gorm:"column:net_sent"`
	NetRecv           string `json:"net_recv" gorm:"column:net_recv"`
	Elapsed           int    `json:"elapsed" gorm:"column:elapsed"`
	CheckSuccess      int    `json:"check_success" gorm:"check_success"`
}
type RServerLogItem struct {
	ServerId          string    `json:"server_id" gorm:"column:server_id"`
	CpuUser           string    `json:"cpu_user" gorm:"column:cpu_user"`
	CpuSystem         string    `json:"cpu_system" gorm:"column:cpu_system"`
	CpuIdle           string    `json:"cpu_idle" gorm:"column:cpu_idle"`
	CpuPercent        string    `json:"cpu_percent" gorm:"column:cpu_percent"`
	MemoryTotal       string    `json:"memory_total" gorm:"column:memory_total"`
	MemoryAvailable   string    `json:"memory_available" gorm:"memory_available"`
	MemoryUsed        string    `json:"memory_used" gorm:"column:memory_used"`
	MemoryUsedPercent string    `json:"memory_used_percent" gorm:"column:memory_used_percent"`
	DiskTotal         string    `json:"disk_total" gorm:"column:disk_total"`
	DiskFree          string    `json:"disk_free" gorm:"column:disk_free"`
	DiskUsed          string    `json:"disk_used" gorm:"column:disk_used"`
	DiskUsedPercent   string    `json:"disk_used_percent" gorm:"column:disk_used_percent"`
	NetSent           string    `json:"net_sent" gorm:"column:net_sent"`
	NetRecv           string    `json:"net_recv" gorm:"column:net_recv"`
	CheckSuccess      int       `json:"check_success" gorm:"check_success"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
}

func GetServerLog(serverId int64) []RServerLogItem {
	var result []RServerLogItem
	sqlStr := `
select * from server_log where server_id=@serverId and deleted_at is null
`
	DB.Raw(sqlStr, map[string]interface{}{
		"serverId": serverId,
	}).Scan(&result)
	return result
}

func AddServerLog(data ServerLogItem) {
	DB.Table("server_log").Create(&data)
}

func GetLatestServerLog(serverId string) RServerLogItem {
	var result RServerLogItem
	sqlStr := `select * from server_log where server_id=@serverId order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"serverId": serverId,
	}).Scan(&result)
	return result
}
