package models

import (
	"WebMonitor/tools"
	"time"
)

type TeamGroup struct {
	DefaultModel
	Name   string `json:"name" gorm:"column:name"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
}

// 创建团队空间
func AddTeamGroup(userId int64, name string) TeamGroup {
	data := TeamGroup{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		Name:         name,
		UserId:       userId,
	}
	DB.Table("team_group").Create(&data)
	return data
}

type UserTeamGroup struct {
	Id        string    `json:"id" grom:"column:id"`
	Name      string    `json:"name" gorm:"column:"name"`
	Username  string    `json:"username" gorm:"column:username"`
	Role      int       `json:"role" gorm:"column:role"`
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at" type:"datetime"`
}

// 获取用户的团队空间
func GetTeamGroupList(userId int64) []UserTeamGroup {
	var result []UserTeamGroup
	sqlStr := `
	SELECT tg.id,
        tg.name,
		tg.created_at,user.username,
		ut.role FROM user_team as ut
LEFT JOIN team_group as tg ON tg.id=ut.team_id
LEFT JOIN user ON user.id=tg.user_id
WHERE
	tg.deleted_at IS NULL
AND ut.deleted_at IS NULL
AND ut.user_id = @userId
	`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId}).Scan(&result)
	return result
}

// 获取团队空间详细信息
func GetTeamGroupDetail(userId int64, teamId int64) UserTeamGroup {
	var result UserTeamGroup
	sqlStr := `select * from team_group where user_id=@userId and id=@teamId`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId, "teamId": teamId}).Scan(&result)
	return result
}

// 更新团队空间
func UpdateTeamGroup(userId int64, teamId int64, name string) {
	sqlStr := `update team_group set name=@name where id=@teamId and user_id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{"userId": userId, "teamId": teamId, "name": name})
}

// 删除团队空间
func DeleteTeamGroup(userId int64, teamId int64) {
	sqlStr := `update team_group set deleted_at=@deletedAt where id=@teamId and user_id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{"userId": userId, "teamId": teamId, "deletedAt": time.Now()})
}

type RUserTeamRole struct {
	Id     string `json:"id" gorm:"column:id"`
	UserId string `json:"user_id" gorm:"column:user_id"`
	TeamId string `json:"team_id" gorm:"column:team_id"`
	Role   int    `json:"role" gorm:"role"`
}

// 检查用户权限
func CheckTeamMemberRole(userId int64, teamId int64) RUserTeamRole {
	var result RUserTeamRole
	sqlStr := `select * from user_team where deleted_at is null and user_id=@userId and team_id=@teamId`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId, "teamId": teamId}).Scan(&result)
	return result
}

// 更新团队创建者
func UpdateTeamOwner(teamId int64, oldOwnerId int64, newOwnerId int64) {
	sqlStr := `update team_group set user_id=@newOwnerId where id=@teamId and user_id=@oldOwnerId`
	DB.Exec(sqlStr, map[string]interface{}{
		"newOwnerId": newOwnerId,
		"oldOwnerId": oldOwnerId,
		"teamId":     teamId,
	})
}

// 团队删除成员
func DeleteTeamMember(teamId int64, userId int64) {
	sqlStr := `update user_team set deleted_at=@deletedAt where team_id=@teamId and user_id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"deletedAt": time.Now(),
		"teamId":    teamId,
		"userId":    userId,
	})
}

type ITeamMember struct {
	DefaultModel
	UserId int64 `json:"user_id" gorm:"column:user_id"`
	TeamId int64 `json:"team_id" gorm:"column:team_id"`
	Role   int   `json:"role" gorm:"column:role"`
}

// 团队新增成员
func AddTeamMember(teamId int64, userId int64, role int) {
	data := ITeamMember{
		DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
		UserId:       userId,
		TeamId:       teamId,
		Role:         role,
	}
	DB.Table("user_team").Create(&data)
}

type RSimpleTeam struct {
	Name     string `json:"name" gorm:"column:name"`
	Username string `json:"username" gorm:"column:username"`
}

func GetSimpleTeamInfo(teamId int64, userId int64) RSimpleTeam {
	var result RSimpleTeam
	sqlStr := `select tg.name,user.username from team_group as tg LEFT JOIN user on user.id=tg.user_id where tg.deleted_at is null and tg.user_id=@userId and tg.id=@teamId`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
		"userId": userId,
	}).Scan(&result)
	return result
}

func GetOwnerTotalTeamNumber(userId int64) RPTotal {
	var result RPTotal
	sqlStr := `
	select count(id) as total from team_group where user_id=@userId and deleted_at is null
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

/*
获取该团队下的成员人数
*/
func GetTeamMemberCount(teamId int64) RPTotal {
	var result RPTotal
	sqlStr := `
	select count(id) as total from user_team where team_id=@teamId and deleted_at is null
	`
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}

/*
根据用户编号获取用户的计划详情
*/
func GetPlanBaseInfoByUserId(userId int64) PlanBase {
	sqlStr := `
	select * from plan_base where id=(select plan_id from user where id=@userId)
	`
	var result PlanBase
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

/*
根据团队编号获取创建人的计划详情
*/
func GetPlanBaseInfoByTeamId(teamId int64) PlanBase {
	sqlStr := `
	select * from plan_base where id=(select plan_id from user where id=(select user_id from team_group where id=@teamId))
	`
	var result PlanBase
	DB.Raw(sqlStr, map[string]interface{}{
		"teamId": teamId,
	}).Scan(&result)
	return result
}
