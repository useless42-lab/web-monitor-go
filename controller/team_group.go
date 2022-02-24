package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"WebMonitor/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	_ "github.com/joho/godotenv/autoload"
)

func GetTeamGroupList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	data := models.GetTeamGroupList(userIdInt)
	response.Success(c, 200, data)
}

type TeamGroupForm struct {
	Name string `json:"name" gorm:"column:name"`
}

func (form TeamGroupForm) ValidateTeamGroupForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("团队名称不能为空")),
	)
}

func AddTeamGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	name := c.PostForm("name")
	teamGroupForm := TeamGroupForm{
		Name: name,
	}
	err1 := teamGroupForm.ValidateTeamGroupForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	err := service.CreateTeamGroupService(userIdInt, name)
	if err != "" {
		response.Error(c, 3000, err)
	} else {
		response.Success(c, 200, "")
	}
}

func GetTeamGroupDetail(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	data := models.GetTeamGroupDetail(userIdInt, teamId)
	response.Success(c, 200, data)
}

func DeleteTeamGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	result := service.DeleteTeamGroupService(userIdInt, teamId)
	if result == "" {
		response.Success(c, 200, result)
	} else {
		response.Error(c, 6060, result)
	}
}

func GetTeamGroupInfo(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.Query("team_id"), 10, 64)
	num := models.CheckIsUserInTeam(userIdInt, teamId)
	if num.Total > 0 {
		result := service.GetTeamGroupInfoService(teamId, userIdInt)
		response.Success(c, 200, result)
	} else {
		response.Success(c, 9000, "无权限")
	}
}

func GenerateTransferTeamLink(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	result := service.GenerateTransferTeamLinkService(teamId, userIdInt)
	response.Success(c, 200, result)
}

func GenerateInviteTeamMemberLink(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	role, _ := strconv.Atoi(c.PostForm("role"))
	result := service.GenerateInviteTeamMemberLinkService(teamId, userIdInt, role)
	response.Success(c, 200, result)
}

func GetTransferTeamInfo(c *gin.Context) {
	var statusCode int
	token := c.Query("token")
	authorization := c.GetHeader("Authorization")
	if authorization != "" {
		data := utils.GetUserFromToken(authorization)
		if data != 0 {
			statusCode = 200
		} else {
			statusCode = 4001
		}
	} else {
		statusCode = 4001
	}

	result := service.GetTransferTeamInfoService(token)
	response.Success(c, statusCode, result)
}

func GetInviteTeamMemberInfo(c *gin.Context) {
	var statusCode int
	token := c.Query("token")
	authorization := c.GetHeader("Authorization")
	if authorization != "" {
		data := utils.GetUserFromToken(authorization)
		if data != 0 {
			statusCode = 200
		} else {
			statusCode = 4001
		}
	} else {
		statusCode = 4001
	}

	result := service.GetInviteTeamMemberInfoService(token)
	response.Success(c, statusCode, result)
}

func UpdateTeamGroup(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	name := c.PostForm("name")
	teamGroupForm := TeamGroupForm{
		Name: name,
	}
	err1 := teamGroupForm.ValidateTeamGroupForm()
	if err1 != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err1))
		return
	}
	service.UpdateTeamGroupService(userIdInt, teamId, name)
	response.Success(c, 200, "")
}
