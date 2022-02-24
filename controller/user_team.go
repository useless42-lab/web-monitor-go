package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTeamGroupMemberList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	num := models.CheckIsUserInTeam(userIdInt, teamId)
	if num.Total > 0 {
		result := service.GetTeamGroupMemberListService(teamId, page, pageSize)
		response.Success(c, 200, result)
	} else {
		response.Success(c, 9000, "无权限")
	}
}

// 退出团队
func ExitTeam(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	canExit := service.ExitTeamService(userIdInt, teamId)
	if canExit == 1 {
		response.Success(c, 200, "退出成功")
	} else {
		response.Success(c, 200, "退出失败")
	}
}

// 踢出成员
func KickOutTeamMember(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	targetUserId, _ := strconv.ParseInt(c.PostForm("target_user_id"), 10, 64)

	result := service.KickOutTeamMemberService(userIdInt, teamId, targetUserId)
	if result == 1 {
		response.Success(c, 200, "踢出成员成功")
	} else {
		response.Success(c, 200, "踢出成员失败")
	}
}

func CreateInviteTeamMember(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	token := c.PostForm("token")
	num := service.CreateInviteTeamMemberService(token, userIdInt)
	if num != 1 {
		response.Success(c, 5000, "已经存在")
	} else {
		response.Success(c, 200, "")
	}
}

func TransferTeamGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	token := c.PostForm("token")
	err := service.TransferTeamGroupService(token, userIdInt)
	if err != "" {
		response.Error(c, 5000, err)
	} else {
		response.Success(c, 200, "")
	}
}

func RemoveAdmin(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	targetUserId, _ := strconv.ParseInt(c.PostForm("target_user_id"), 10, 64)
	service.RemoveAdminService(userIdInt, teamId, targetUserId)
	response.Success(c, 200, "")
}

func AddAdmin(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	targetUserId, _ := strconv.ParseInt(c.PostForm("target_user_id"), 10, 64)
	service.AddAdminService(userIdInt, teamId, targetUserId)
	response.Success(c, 200, "")
}
