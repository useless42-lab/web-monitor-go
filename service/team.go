package service

import (
	"WebMonitor/cache"
	"WebMonitor/models"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

func GetTeamGroupMemberListService(teamId int64, page int, pageSize int) models.PaginationData {
	data := models.GetTeamGroupMemberList(teamId, page, pageSize)
	total := models.GetTeamGroupMemberNumber(teamId)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

type RTeamGroupInfo struct {
	Id               string           `json:"id" gorm:"column:id"`
	Name             string           `json:"name" gorm:"column:name"`
	UserId           string           `json:"user_id" gorm:"column:user_id"`
	CreatedAt        models.LocalTime `json:"created_at" gorm:"column:created_at"`
	MemberTotal      int              `json:"member_total"`
	DeviceGroupTotal int              `json:"device_group_total"`
	DeviceTotal      int              `json:"device_total"`
	Role             int              `json:"role" gorm:"column:role"`
}

func GetTeamGroupInfoService(teamId int64, userId int64) RTeamGroupInfo {
	userTeamInfo := models.GetUserTeamInfo(teamId, userId)
	info := models.GetTeamGroupInfo(teamId)
	memberTotal := models.GetTeamGroupMemberNumber(teamId)
	// deviceGroupTotal := models.GetTeamGroupDeviceGroupNumber(teamId)
	deviceGroupTotal := models.GetDeviceGroupCount(userId, teamId)
	deviceTotal := models.GetAllDeviceNumber(teamId)
	result := RTeamGroupInfo{
		Id:               info.Id,
		Name:             info.Name,
		UserId:           info.UserId,
		CreatedAt:        info.CreatedAt,
		MemberTotal:      memberTotal.Total,
		DeviceGroupTotal: deviceGroupTotal.Total,
		DeviceTotal:      deviceTotal.Total,
		Role:             userTeamInfo.Role,
	}
	return result
}

// 删除团队
func DeleteTeamGroupService(userId int64, teamId int64) string {
	// GetUserTeamGroupNumber
	count := models.GetOwnerTotalTeamNumber(userId)
	if count.Total < 2 {
		// 最后一个团队不能删除
		return "最后一个团队不能删除"
	} else {
		models.DeleteTeamGroup(userId, teamId)
		return ""
	}
}

// 退出团队
func ExitTeamService(userId int64, teamId int64) int {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 {
		// 创建者无法直接退出
		return 0
	} else {
		models.DeleteTeamMember(teamId, userId)
		return 1
	}
}

// 踢出成员
func KickOutTeamMemberService(userId int64, teamId int64, targetUserId int64) int {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 1 {
		// 普通成员没有权限
		return 0
	} else {
		models.DeleteTeamMember(teamId, targetUserId)
		return 1
	}
}

// 解除管理员
func RemoveAdminService(userId int64, teamId int64, targetUserId int64) {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 {
		// 创建人有权限
		models.UpdateUserTeamRole(teamId, targetUserId, 1)
	} else {
	}
}

// 任命管理员
func AddAdminService(userId int64, teamId int64, targetUserId int64) {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 {
		// 创建人有权限
		models.UpdateUserTeamRole(teamId, targetUserId, 3)
	} else {

	}
}

// 生成转让团队链接
func GenerateTransferTeamLinkService(teamId int64, userId int64) string {
	token := md5.New()
	token.Write([]byte(strconv.FormatInt(teamId, 10) + strconv.FormatInt(userId, 10) + time.Now().String()))
	finalToken := hex.EncodeToString(token.Sum(nil))
	// cache.Set("user:token:"+finalToken, strconv.Itoa(int(result.DefaultModel.ID)), 60*60*24*7)
	data := strconv.FormatInt(teamId, 10) + " " + strconv.FormatInt(userId, 10)
	cache.Set("team:group:"+finalToken, data, 30*60)
	return finalToken
}

// 生成邀请成员链接
func GenerateInviteTeamMemberLinkService(teamId int64, userId int64, role int) string {
	token := md5.New()
	token.Write([]byte(strconv.FormatInt(teamId, 10) + strconv.FormatInt(userId, 10) + strconv.Itoa(role) + time.Now().String()))
	finalToken := hex.EncodeToString(token.Sum(nil))
	// cache.Set("user:token:"+finalToken, strconv.Itoa(int(result.DefaultModel.ID)), 60*60*24*7)
	data := strconv.FormatInt(teamId, 10) + " " + strconv.FormatInt(userId, 10) + " " + strconv.Itoa(role)
	cache.Set("team:member:"+finalToken, data, 30*60)
	return finalToken
}

func GetTransferTeamInfoService(token string) models.RSimpleTeam {
	var result models.RSimpleTeam
	data := cache.Get("team:group:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		ownerId, _ := strconv.ParseInt(arr[1], 10, 64)
		result = models.GetSimpleTeamInfo(teamId, ownerId)
		return result
	} else {
		return result
	}
}

func TransferTeamGroupService(token string, userId int64) string {
	data := cache.Get("team:group:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		ownerId, _ := strconv.ParseInt(arr[1], 10, 64)

		planId := models.GetPlanIdByUserId(userId)
		plan := models.GetPlanBaseInfo(planId.Id)

		teamNumber := models.GetOwnerTotalTeamNumber(userId)
		if plan.TeamNumber <= teamNumber.Total {
			return "被转让人所拥有空间已达上限"
		}
		// 更改 team_group 创建人
		models.UpdateTeamOwner(teamId, ownerId, userId)
		// 检查接收者是否在该团队空间
		num := models.CheckIsUserInTeam(userId, teamId)
		if num.Total > 0 {
			// 已经存在该成员
			// 升级该成员权限至 创建者 2
			models.UpdateUserTeamRole(teamId, userId, 2)
		} else {
			models.AddUserTeam(userId, teamId, 2)
		}
		// 原创建者
		number := models.GetUserTeamGroupNumber(ownerId)
		models.DeleteTeamMember(teamId, ownerId)
		if number.Total > 0 {
		} else {
			// 转让后无团队空间 初始化
			InitNewTeamGroupService(ownerId)
		}
		cache.Del("team:group:" + token)
	} else {
		return "链接失效"
	}
	return ""
}

type RSimpleTeamWithRole struct {
	models.RSimpleTeam
	Role int `json:"role"`
}

func GetInviteTeamMemberInfoService(token string) RSimpleTeamWithRole {
	var result RSimpleTeamWithRole
	data := cache.Get("team:member:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		ownerId, _ := strconv.ParseInt(arr[1], 10, 64)
		role, _ := strconv.Atoi(arr[2])
		result.RSimpleTeam = models.GetSimpleTeamInfo(teamId, ownerId)
		result.Role = role
		return result
	} else {
		return result
	}
}

func CreateInviteTeamMemberService(token string, userId int64) int {
	data := cache.Get("team:member:" + token)
	if data != "" {
		arr := strings.Fields(data)
		teamId, _ := strconv.ParseInt(arr[0], 10, 64)
		role, _ := strconv.Atoi(arr[2])
		teamMemberTotal := models.GetTeamMemberCount(teamId)
		planBaseItem := models.GetPlanBaseInfoByTeamId(teamId)
		if teamMemberTotal.Total < planBaseItem.TeamMemberLimit {
			num := models.CheckIsUserInTeam(userId, teamId)
			if num.Total > 0 {
				// 已经存在该成员
				return 0
			} else {
				models.AddUserTeam(userId, teamId, role)
				return 1
			}
		}
	}
	return 0
}

func UpdateTeamGroupService(userId int64, teamId int64, name string) {
	role := models.CheckTeamMemberRole(userId, teamId)
	if role.Role == 2 || role.Role == 3 {
		// 创建人有权限
		models.UpdateTeamGroup(userId, teamId, name)
	} else {

	}
}
