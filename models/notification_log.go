package models

type NotificationLogItem struct {
	DefaultModel
	NotificationId int64 `json:"notification_id" gorm:"column:notification_id"`
	DeviceType     int   `json:"device_type" gorm:"column:device_type"`
	DeviceId       int64 `json:"device_id" gorm:"column:device_id"`
}

func AddNotificationLog(notificationId int64, deviceType int, deviceId int64) {
	notificationItem := NotificationLogItem{
		NotificationId: notificationId,
		DeviceType:     deviceType,
		DeviceId:       deviceId,
	}
	DB.Table("notification_log").Create(&notificationItem)
}

func GetLatestNotificationLog(notificationId int64, deviceType int, deviceId int64) NotificationLogItem {
	var result NotificationLogItem
	sqlStr := `select * from notification_log where notification_id=@notificationId and device_id=@deviceId and device_type=@deviceType order by id desc limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"notificationId": notificationId,
		"deviceType":     deviceType,
		"deviceId":       deviceId,
	}).Scan(&result)
	return result
}
