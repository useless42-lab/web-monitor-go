package models

type NotificationStruct struct {
	DefaultModel
	UserId     int64  `json:"user_id" gorm:"column:user_id"`
	Email      string `json:"email" gorm:"column:email"`
	Phone      string `json:"phone" gorm:"column:phone"`
	SMS        string `json:"sms" gorm:"column:sms"`
	Telegram   string `json:"telegram" gorm:"column:telegram"`
	Bark       string `json:"bark" gorm:"gorm:bark"`
	ServerChan string `json:"server_chan" gorm:"column:server_chan"`
}

type RNotificationStruct struct {
	UserId     string `json:"user_id" gorm:"column:user_id"`
	Email      string `json:"email" gorm:"column:email"`
	Phone      string `json:"phone" gorm:"column:phone"`
	SMS        string `json:"sms" gorm:"column:sms"`
	Telegram   string `json:"telegram" gorm:"column:telegram"`
	Bark       string `json:"bark" gorm:"gorm:bark"`
	ServerChan string `json:"server_chan" gorm:"column:server_chan"`
}

func AddNotification(userId int64, email string, phone string, sms string, telegram string, bark string, serverChan string) {
	notification := NotificationStruct{
		UserId:     userId,
		Email:      email,
		Phone:      phone,
		SMS:        sms,
		Telegram:   telegram,
		Bark:       bark,
		ServerChan: serverChan,
	}
	DB.Table("notification_base").Create(&notification)
}

func GetNotification(userId int64) RNotificationStruct {
	var result RNotificationStruct
	sqlStr := `select * from notification_base where user_id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

func UpdateNotification(userId int64, email string, phone string, sms string, telegram string, bark string, serverChan string) {
	sqlStr := `update notification_base set email=@email,phone=@phone,sms=@sms,telegram=@telegram,bark=@bark,server_chan=@serverChan where user_id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"email":      email,
		"phone":      phone,
		"sms":        sms,
		"telegram":   telegram,
		"bark":       bark,
		"serverChan": serverChan,
		"userId":     userId,
	})
}
