package models

import "WebMonitor/tools"

type HeartbeatItem struct {
	DefaultModel
	UserId        int64  `json:"user_id" gorm:"column:user_id"`
	Name          string `json:"name" gorm:"column:name"`
	Token         string `json:"token" gorm:"column:token"`
	GroupId       int64  `json:"group_id" gorm:"column:group_id"`
	PolicyId      int64  `json:"policy_id" gorm:"column:policy_id"`
	MonitorRegion string `json:"monitor_region" gorm:"column:monitor_region"`
}

type RHeartbeatItem struct {
	Id              string    `json:"id" gorm:"column:id"`
	Name            string    `json:"name" gorm:"column:name"`
	Token           string    `json:"token" gorm:"column:token"`
	GroupId         int64     `json:"group_id" gorm:"column:group_id"`
	Frequency       int       `json:"frequency" gorm:"column:frequency"`
	FailedWaitTimes int       `json:"failed_wait_times" gorm:"column:failed_wait_times"`
	CreatedAt       LocalTime `json:"created_at" gorm:"column:created_at"`
	Status          int       `json:"status" gorm:"column:status"`
}

func AddHeartbeat(userId int64, name string, token string, groupId int64, policyId int64, monitorRegion string) HeartbeatItem {
	heartbeat := HeartbeatItem{
		DefaultModel:  DefaultModel{ID: tools.GenerateSnowflakeId()},
		Name:          name,
		UserId:        userId,
		Token:         token,
		GroupId:       groupId,
		PolicyId:      policyId,
		MonitorRegion: monitorRegion,
	}
	DB.Table("heartbeat_list").Create(&heartbeat)
	return heartbeat
}

func UpdateHeartbeat(heartbeatId int64, name string, policyId int64, monitorRegion string) {
	sqlStr := `
	update heartbeat_list set name=@name,policy_id=@policyId,monitor_region=@monitorRegion where id=@heartbeatId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"heartbeatId":   heartbeatId,
		"name":          name,
		"policyId":      policyId,
		"monitorRegion": monitorRegion,
	})
}

func GetHearbeatList(userId int64, groupId int64, page int, pageSize int) []RHeartbeatItem {
	var result []RHeartbeatItem
	sqlStr := `select * from heartbeat_list where user_id=@userId and group_id=@groupId and deleted_at is null and  (status=1 or status=3) order by id desc LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":   userId,
		"groupId":  groupId,
		"pageSize": pageSize,
		"offset":   (page - 1) * pageSize,
	}).Scan(&result)
	return result
}

func GetFiledHeartbeatList(userID int64, page int, pageSize int) []RHeartbeatItem {
	var result []RHeartbeatItem
	sqlStr := `SELECT
	al.id,
	al.name,
	al.path,
	al.created_at,
	al.status
FROM
	(
		SELECT
			*
		FROM
			heartbeat_list
		WHERE
			group_id IN (
				SELECT
					id
				FROM
					device_group
				WHERE
					team_id IN (
						SELECT
							id
						FROM
							team_group
						WHERE
							user_id = @userId
					)
			)
	) AS al
WHERE
	al.status = 2
AND al.deleted_at IS NULL
LIMIT @pageSize OFFSET @offset`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userID, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}

func GetHeartbeatIdByToken(token string) RHeartbeatItem {
	var result RHeartbeatItem
	sqlStr := `
	SELECT
	*
FROM
	heartbeat_list
LEFT JOIN monitor_policy ON monitor_policy.id = heartbeat_list.policy_id
WHERE
	heartbeat_list.deleted_at IS NULL AND heartbeat_list.token=@token
AND heartbeat_list.status = 1
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"token": token,
	}).Scan(&result)
	return result
}
