package models

import "WebMonitor/tools"

type UserTeam struct {
	DefaultModel
	UserId int64 `json:"user_id" grom:"column:user_id"`
	TeamId int64 `json:"team_id" gorm:"column:team_id"`
	Role   int   `json:"role" gorm:"column:role"`
}

func AddUserTeam(userId int64, teamId int64, role int) UserTeam {
	data := UserTeam{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		UserId:       userId,
		TeamId:       teamId,
		Role:         role,
	}
	DB.Table("user_team").Create(&data)
	return data
}

type RTeamMember struct {
	Id        string    `json:"id" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Role      int       `json:"role" gorm:"column:role"`
	TeamId    string    `json:"team_id" gorm:"column:team_id"`
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at"`
}

func GetTeamGroupMemberList(teamId int64, page int, pageSize int) []RTeamMember {
	var result []RTeamMember
	sqlStr := `
	select user.id,username,ut.role,team_id,ut.created_at from user_team as ut left join user on user.id=ut.user_id where team_id=@teamId and ut.deleted_at is null LIMIT @pageSize OFFSET @offset
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId, "pageSize": pageSize, "offset": (page - 1) * pageSize}).Scan(&result)
	return result
}

type RTeamGroupInfo struct {
	Id        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	UserId    string    `json:"user_id" gorm:"column:user_id"`
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at"`
}

type RUserTeamInfo struct {
	Role int `json:"role" gorm:"column:role"`
}

func GetUserTeamInfo(teamId int64, userId int64) RUserTeamInfo {
	var result RUserTeamInfo
	sqlStr := `select * from user_team where team_id=@teamId and user_id=@userId and deleted_at is null`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
		"userId": userId,
	}).Scan(&result)
	return result
}

func GetTeamGroupInfo(teamId int64) RTeamGroupInfo {
	var result RTeamGroupInfo
	sqlStr := `
	select tg.id ,tg.name,tg.user_id,tg.created_at from team_group as tg where deleted_at is null and id=@teamId
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId}).Scan(&result)
	return result
}

func GetTeamGroupMemberNumber(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `
	select count(*)as total from user_team where deleted_at is null and team_id=@teamId
	`
	DB.Raw(sqlStr, map[string]interface{}{"teamId": teamId}).Scan(&result)
	return result
}

func CheckIsUserInTeam(userId int64, teamId int64) RPTotal {
	sqlStr := `
	select count(*) as total from user_team where user_id=@userId and team_id=@teamId and deleted_at is null
	`
	var result RPTotal
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId, "teamId": teamId}).Scan(&result)
	return result
}

func UpdateUserTeamRole(teamId int64, userId int64, role int) {
	sqlStr := `update user_team set role=@role where team_id=@teamId and user_id=@userId and deleted_at is null`
	DB.Exec(sqlStr, map[string]interface{}{
		"role":   role,
		"teamId": teamId,
		"userId": userId,
	})
}

func GetUserTeamGroupNumber(userId int64) RPTotal {
	var result RPTotal
	sqlStr := `
	select count(id) as total from user_team where user_id=@userId and deleted_at is null
	`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId}).Scan(&result)
	return result
}
