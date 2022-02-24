package controller

import (
	"WebMonitor/models"
	"WebMonitor/monitor"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ApiForm struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	TeamId   int64  `json:"team_id"`
	GroupId  int64  `json:"group_id"`
	PolicyId int64  `json:"policy_id"`
	Method   int64  `json:"method"`
	BodyType int64  `json:"body_type"`
	BodyJson string `json:"body_json"`
}

func (form ApiForm) ValidateApiForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Path, validation.Required.Error("网址不能为空")),
		validation.Field(&form.TeamId, validation.Required.Error("团队不能为空")),
		validation.Field(&form.GroupId, validation.Required.Error("分组不能为空")),
		validation.Field(&form.PolicyId, validation.Required.Error("策略不能为空")),
		validation.Field(&form.Method, validation.Required.Error("请求方法不能为空")),
		validation.Field(&form.BodyType, validation.Required.Error("不能为空")),
	)
}

type BodyJsonForm struct {
	BodyJson string `json:"body_json"`
}

func (form BodyJsonForm) ValidateBodyJsonForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.BodyJson, validation.Required.Error("不能为空"), is.JSON.Error("json格式有误")),
	)
}

func AddApi(c *gin.Context) {
	userId := c.GetInt64("userId")
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	method, _ := strconv.ParseInt(c.PostForm("method"), 10, 64)
	requestHeaders := c.PostForm("request_headers")
	bodyType, _ := strconv.ParseInt(c.PostForm("body_type"), 10, 64)
	bodyRaw := c.PostForm("body_raw")
	bodyJson := c.PostForm("body_json")
	bodyForm := c.PostForm("body_form")
	responseData := c.PostForm("response_data")
	basicUser := c.PostForm("basic_user")
	basicPassword := c.PostForm("basic_password")
	monitorRegion := c.PostForm("monitor_region")

	apiForm := ApiForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
		Method:   method,
		BodyType: bodyType,
	}
	bodyJsonForm := BodyJsonForm{
		BodyJson: bodyJson,
	}
	err := apiForm.ValidateApiForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	if bodyJson != "" {
		err = bodyJsonForm.ValidateBodyJsonForm()
		if err != nil {
			response.Error(c, 4000, response.ConvertValidationErrorToString(err))
			return
		}
	}

	result := service.AddApiService(userId, name, path, teamId, groupId, policyId, method, requestHeaders, bodyType, bodyRaw, bodyJson, bodyForm, responseData, basicUser, basicPassword, monitorRegion)

	if result.ID == 0 {
		response.Error(c, 3000, "设备数量超出上限")
	} else {
		response.Success(c, 200, result)
	}

}

func GetApiList(c *gin.Context) {
	userId := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetApiListService(userId, groupId, page, pageSize)
	response.Success(c, 200, result)
}

func UpdateApi(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	method, _ := strconv.ParseInt(c.PostForm("method"), 10, 64)
	requestHeaders := c.PostForm("request_headers")
	bodyType, _ := strconv.ParseInt(c.PostForm("body_type"), 10, 64)
	bodyRaw := c.PostForm("body_raw")
	bodyJson := c.PostForm("body_json")
	bodyForm := c.PostForm("body_form")
	responseData := c.PostForm("response_data")
	basicUser := c.PostForm("basic_user")
	basicPassword := c.PostForm("basic_password")
	monitorRegion := c.PostForm("monitor_region")

	apiForm := ApiForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
		Method:   method,
		BodyType: bodyType,
	}
	bodyJsonForm := BodyJsonForm{
		BodyJson: bodyJson,
	}
	err := apiForm.ValidateApiForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	if bodyJson != "" {
		err = bodyJsonForm.ValidateBodyJsonForm()
		if err != nil {
			response.Error(c, 4000, response.ConvertValidationErrorToString(err))
			return
		}
	}

	result := models.UpdateApi(id, name, path, groupId, policyId, method, requestHeaders, bodyType, bodyRaw, bodyJson, bodyForm, responseData, basicUser, basicPassword, monitorRegion)
	response.Success(c, 200, result)
}

func GetApi(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	result := models.GetApi(id)
	response.Success(c, 200, result)
}

func MockApi(c *gin.Context) {
	apiId, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	apiInfo := models.GetApi(apiId)
	var requestData string
	if apiInfo.BodyType == 1 {
		requestData = apiInfo.BodyRaw
	} else if apiInfo.BodyType == 2 {
		requestData = apiInfo.BodyJson
	} else if apiInfo.BodyType == 3 {
		requestData = apiInfo.BodyForm
	}
	_, _, body, _, _ := monitor.MonitorApi(apiInfo.Path, apiInfo.Method, apiInfo.RequestHeaders, apiInfo.BodyType, requestData, apiInfo.ResponseData, apiInfo.BasicUser, apiInfo.BasicPassword)
	response.Success(c, 200, body)
}

func GetFiledApiList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetFiledApiListService(userIdInt, page, pageSize)
	response.Success(c, 200, result)
}

func FileApi(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.FileDevice(id, 3)
	response.Success(c, 200, "")
}

func StartMonitorApi(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.StartMonitor(id, 3)
	response.Success(c, 200, "")
}

func PauseMonitorApi(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.PauseMonitor(id, 3)
	response.Success(c, 200, "")
}

func DeleteApi(c *gin.Context) {
	apiId, _ := strconv.ParseInt(c.Query("api_id"), 10, 64)
	models.DeleteDevice(apiId, 3)
	response.Success(c, 200, "")
}
