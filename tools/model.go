package tools

type RNotificationItemStruct struct {
	NotificationType   int `json:"notification_type" gorm:"column:notification_type"`
	NotificationTarget string
	Email              string `json:"email" gorm:"column:email"`
	Phone              string `json:"phone" gorm:"column:phone"`
	SMS                string `json:"sms" gorm:"column:sms"`
	Telegram           string `json:"telegram" gorm:"column:telegram"`
	Bark               string `json:"bark" gorm:"gorm:bark"`
	ServerChan         string `json:"server_chan" gorm:"column:server_chan"`
}

type RDeviceItemJob struct {
	Name             string                  `json:"name" gorm:"name"`
	Path             string                  `json:"path" gorm:"path"`
	NotificationItem RNotificationItemStruct `json:"notification_item" gorm:"column:notification_item"`
}
