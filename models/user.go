package models

import (
	"time"
)

type User struct {
	DefaultModel
	// Username  string    `json:"username" gorm:"column:username"`
	// Email     string    `json:"email" gorm:"column:email" gorm:"unique"`
	// Password  string    `json:"-" gorm:"column:password"`
	// Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Username  string    `json:"username" gorm:"column:username"`
	PlanId    int       `json:'plan_id gorm:"column:plan_id"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

// 添加用户
func AddUser(userId int64, username string, planId int, expiredAt time.Time) User {
	// encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }
	// user := User{
	// 	DefaultModel: DefaultModel{ID: tools.GenerateSnowflakeId()},
	// 	Username:     username,
	// 	Email:        email,
	// 	Password:     string(encryptPassword),
	// }
	user := User{
		DefaultModel: DefaultModel{ID: userId},
		Username:     username,
		PlanId:       planId,
		ExpiredAt:    expiredAt,
	}
	DB.Table("user").Create(&user)
	return user
}

// 检查该用户是否存在本系统
func CheckIsExistUserId(userId int64) bool {
	sqlStr := `select id from user where id=@userId`
	var result User
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	if result.DefaultModel.ID == 0 {
		return false
	} else {
		return true
	}
}

type BaseUserInfo struct {
	Id        string    `json:"id" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	PlanName  string    `json:"plan_name" gorm:"column:plan_name"`
	ExpiredAt LocalTime `json:"expired_at" gorm:"column:expired_at"`
}

func GetBaseUserInfo(userId int64) BaseUserInfo {
	var result BaseUserInfo
	sqlStr := `select user.id,username,pb.name as plan_name,user.expired_at from user left join plan_base as pb on pb.id=user.plan_id where user.id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{"userId": userId}).Scan(&result)
	return result
}

type UserPlanStruct struct {
	PlanId    int       `json:"plan_id" gorm:"column:plan_id"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

func GetUserPlan(userId int64) UserPlanStruct {
	var result UserPlanStruct
	sqlStr := `select plan_id,expired_at from user where id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

type ThirdPartyItem struct {
	SteamApiKey string `json:"steam_api_key" gorm:"column:steam_api_key"`
}

func GetUserThirdPartyInfo(userId int64) ThirdPartyItem {
	var result ThirdPartyItem
	sqlStr := `select steam_api_key from user where id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

func UpdateUserThirdPartyInfo(userId int64, steamApiKey string) {
	sqlStr := `update user set steam_api_key=@steamApiKey where id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"userId":      userId,
		"steamApiKey": steamApiKey,
	})
}

func UpdateUserPlan(userId int64, planId int, expiredAt time.Time) {
	sqlStr := `update user set plan_id=@planId , expired_at=@expiredAt where id=@userId`
	DB.Exec(sqlStr, map[string]interface{}{
		"userId":    userId,
		"expiredAt": expiredAt,
		"planId":    planId,
	})
}

type RUserPlan struct {
	Name      string    `json:"name" gorm:"column:name"`
	ExpiredAt LocalTime `json:"expired_at" gorm:"column:expired_at"`
}

func GetUserPlanDetail(userId int64) RUserPlan {
	var result RUserPlan
	sqlStr := `select user.expired_at,pb.name from user left join plan_base as pb on pb.id=user.plan_id where user.id=@userId`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}
