package models

type NotificationListItem struct {
	DefaultModel
	UserId           int64 `json:"user_id" gorm:"column:user_id"`
	GroupId          int64 `json:"group_id" gorm:"column:group_id"`
	DeviceType       int   `json:"device_type" gorm:"column:device_type"`
	NotificationType int   `json:"notification_type" gorm:"column:notification_type"`
}

type RNotificationListItem struct {
	Id               string    `json:"id" gorm:"column:id"`
	Username         string    `json:"username" gorm:"column:username"`
	NotificationType int       `json:"notification_type" gorm:"column:notification_type"`
	CreatedAt        LocalTime `json:"created_at"`
}

// 设备组 添加
func AddNotificationListItem(userId int64, groupId int64, deviceType int, notificationType int) {
	notification := NotificationListItem{
		UserId:           userId,
		GroupId:          groupId,
		DeviceType:       deviceType,
		NotificationType: notificationType,
	}
	DB.Table("notification_list").Create(&notification)
}

func DeleteNotificationListItem(id int64) {
	sqlStr := `delete from notification_list where id=@id`
	DB.Exec(sqlStr, map[string]interface{}{
		"id": id,
	})
}

func GetNotificationList(groupId int64) []RNotificationListItem {
	var result []RNotificationListItem
	sqlStr := `
	SELECT
	nl.id,
	u.username,
	nl.created_at,
	nl.notification_type 
FROM
	notification_list AS nl
	LEFT JOIN user as u ON u.id = nl.user_id 
WHERE
	nl.group_id = @groupId 
	AND nl.deleted_at IS NULL`
	DB.Raw(sqlStr, map[string]interface{}{
		"groupId": groupId,
	}).Scan(&result)
	return result
}

type RTeamMemberBase struct {
	UserId   string `json:"id" gorm:"column:id"`
	Username string `json:"username"`
}

func GetTeamMemberListByGroupId(groupId int64) []RTeamMemberBase {
	var result []RTeamMemberBase
	sqlStr := `
	SELECT
	u.id,
	u.username 
FROM
	(
	SELECT
		ut.user_id 
	FROM
		user_team AS ut 
	WHERE
		team_id =(
		SELECT
			dg.team_id 
		FROM
			device_group AS dg 
		WHERE
			id = @groupId 
			AND deleted_at IS NULL 
		) 
		AND deleted_at IS NULL 
	) AS a
	LEFT JOIN user AS u ON u.id = a.user_id
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"groupId": groupId,
	}).Scan(&result)
	return result
}
