package service

import (
	"WebMonitor/models"
)

func InitNewTeamGroupService(userId int64) {
	team := models.AddTeamGroup(userId, "默认空间")
	models.AddMonitorPolicy(team.ID, "默认策略", 3600, 1, 1, 1, "200", "", "200", 100, 100, 100, 0, 2592000, 0, 2592000, 3)
	models.AddUserTeam(userId, int64(team.ID), 2)
	// 1网站 2服务器 3接口 4tcp 5dns 6心跳 7steam 8minecraft
	// 网站
	dWeb := models.AddDeviceGroup(int64(team.ID), "默认分组", 1)
	models.AddNotificationListItem(userId, dWeb.ID, 1, 1)
	// 服务器
	dServer := models.AddDeviceGroup(int64(team.ID), "默认分组", 2)
	models.AddNotificationListItem(userId, dServer.ID, 2, 1)
	// 接口
	dApi := models.AddDeviceGroup(int64(team.ID), "默认分组", 3)
	models.AddNotificationListItem(userId, dApi.ID, 3, 1)
	// tcp
	dTcp := models.AddDeviceGroup(int64(team.ID), "默认分组", 4)
	models.AddNotificationListItem(userId, dTcp.ID, 4, 1)
	// dns
	dDns := models.AddDeviceGroup(int64(team.ID), "默认分组", 5)
	models.AddNotificationListItem(userId, dDns.ID, 5, 1)
	// 心跳
	dHeartbeat := models.AddDeviceGroup(int64(team.ID), "默认分组", 6)
	models.AddNotificationListItem(userId, dHeartbeat.ID, 6, 1)
	// steam
	dSteam := models.AddDeviceGroup(int64(team.ID), "默认分组", 7)
	models.AddNotificationListItem(userId, dSteam.ID, 7, 1)
	// minecraft
	dMinecraft := models.AddDeviceGroup(int64(team.ID), "默认分组", 8)
	models.AddNotificationListItem(userId, dMinecraft.ID, 8, 1)
}

func CreateTeamGroupService(userId int64, teamName string) string {
	planId := models.GetPlanIdByUserId(userId)
	plan := models.GetPlanBaseInfo(planId.Id)
	teamNumber := models.GetUserTeamGroupNumber(userId)
	if teamNumber.Total < plan.TeamNumber {
		team := models.AddTeamGroup(userId, teamName)
		models.AddUserTeam(userId, int64(team.ID), 2)
		// 网站
		d1 := models.AddDeviceGroup(int64(team.ID), "默认分组", 1)
		// 服务器
		d2 := models.AddDeviceGroup(int64(team.ID), "默认分组", 2)
		// 接口
		d3 := models.AddDeviceGroup(int64(team.ID), "默认分组", 3)
		// tcp
		d4 := models.AddDeviceGroup(int64(team.ID), "默认分组", 4)
		// dns
		d5 := models.AddDeviceGroup(int64(team.ID), "默认分组", 5)
		// 心跳
		d6 := models.AddDeviceGroup(int64(team.ID), "默认分组", 6)
		// steam
		d7 := models.AddDeviceGroup(int64(team.ID), "默认分组", 7)
		// minecraft
		d8 := models.AddDeviceGroup(int64(team.ID), "默认分组", 8)

		// 1邮件
		models.AddNotificationListItem(userId, d1.ID, 1, 1)
		models.AddNotificationListItem(userId, d2.ID, 2, 1)
		models.AddNotificationListItem(userId, d3.ID, 3, 1)
		models.AddNotificationListItem(userId, d4.ID, 4, 1)
		models.AddNotificationListItem(userId, d5.ID, 5, 1)
		models.AddNotificationListItem(userId, d6.ID, 6, 1)
		models.AddNotificationListItem(userId, d7.ID, 7, 1)
		models.AddNotificationListItem(userId, d8.ID, 8, 1)
		return ""
	} else {
		return "团队空间数量超出上限"
	}
}

func GetDeviceGroupPaginationListService(userId int64, teamId int64, page int, pageSize int) models.PaginationData {
	data := models.GetDeviceGroupList(userId, teamId, page, pageSize)
	total := models.GetDeviceGroupCount(userId, teamId)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

func AddDeviceGroupService(userId int64, teamId int64, name string, deviceType int) string {
	planId := models.GetPlanIdByTeamId(teamId)
	plan := models.GetPlanBaseInfo(planId.Id)
	urlGroupNumber := models.GetDeviceGroupCount(userId, teamId)
	if urlGroupNumber.Total < plan.PerTeamGroupLimit {
		// 1网站 2服务器 3接口 4tcp 5dns 6心跳 7steam 8minecraft
		deviceGroup := models.AddDeviceGroup(teamId, name, deviceType)
		// 1 email
		models.AddNotificationListItem(userId, deviceGroup.ID, deviceType, 1)
		return ""
	} else {
		return "设备分组数量超出上限"
	}
}

func DeleteDeviceGroupService(userId int64, teamId int64, groupId int64, deviceType int) string {
	total := models.GetDeviceGroupDeviceNumber(groupId, deviceType)
	if total.Total > 0 {
		return "该分组下还有设备，无法删除"
	} else {
		total = models.GetDeviceGroupCountByDeviceType(userId, teamId, deviceType)
		if total.Total == 1 {
			return "无法删除团队设备分类下最后一个分组"
		} else {
			models.DeleteDeviceGroup(userId, groupId)
			return ""
		}
	}
}
