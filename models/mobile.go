package models

type ElapsedItem struct {
	Max string `json:"max_elapsed" gorm:"column:max_elapsed"`
	Min string `json:"min_elapsed" gorm:"column:min_elapsed"`
	Avg string `json:"avg_elapsed" gorm:"column:avg_elapsed"`
}

type LatestLogItem struct {
	CheckSuccess int       `json:"check_success" gorm:"column:check_success"`
	Elapsed      string    `json:"elapsed" gorm:"column:elapsed"`
	LogTime      LocalTime `json:"created_at" gorm:"column:created_at"`
}

type RMobileDeviceList struct {
	Name       string        `json:"name"`
	Path       string        `json:"path"`
	DeviceType int           `json:"device_type"`
	Elapsed    ElapsedItem   `json:"elapsed" grom:"-"`
	LatestLog  LatestLogItem `json:"latest_log" gorm:"-"`
}

func GetDeviceElapsed(deviceId string, deviceType int) ElapsedItem {
	var result ElapsedItem
	device := FilterDeviceLog(deviceType)
	deviceIdColumn := FilterDeviceId(deviceType)
	sqlStr := `
SELECT
	max(elapsed) AS max_elapsed,
	min(elapsed) AS min_elapsed,
	avg(elapsed) AS avg_elapsed
FROM
	` + device + `
WHERE
	` + deviceIdColumn + ` = @deviceId
`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
	}).Scan(&result)
	return result
}

func GetDeviceLatestLog(deviceId string, deviceType int) LatestLogItem {
	var result LatestLogItem
	deviceLog := FilterDeviceLog(deviceType)
	deviceIdColumn := FilterDeviceId(deviceType)
	sqlStr := `
	select created_at,check_success,elapsed from ` + deviceLog + ` where ` + deviceIdColumn + `=@deviceId order by id desc limit 1
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"deviceId": deviceId,
	}).Scan(&result)
	return result
}
