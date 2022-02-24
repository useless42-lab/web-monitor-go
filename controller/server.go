package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"WebMonitor/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ServerForm struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	TeamId   int64  `json:"team_id"`
	GroupId  int64  `json:"group_id"`
	PolicyId int64  `json:"policy_id"`
}

func (form ServerForm) ValidateServerForm() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Name, validation.Required.Error("名称不能为空")),
		validation.Field(&form.Path, validation.Required.Error("地址不能为空")),
		validation.Field(&form.TeamId, validation.Required.Error("团队不能为空")),
		validation.Field(&form.GroupId, validation.Required.Error("分组不能为空")),
		validation.Field(&form.PolicyId, validation.Required.Error("策略不能为空")),
	)
}

func getUniqueShorterStr() string {
	randomStr := utils.GetRandomStr(32)
	result := models.IsTokenExist(randomStr)
	if result {
		getUniqueShorterStr()
	} else {
		return randomStr
	}
	return randomStr
}

func AddServer(c *gin.Context) {
	userId := c.GetInt64("userId")
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	serverFrom := ServerForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}

	err := serverFrom.ValidateServerForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	result := service.AddServerService(userId, name, path, teamId, groupId, policyId, getUniqueShorterStr(), monitorRegion)
	if result != "" {
		response.Error(c, 3000, result)
	} else {
		response.Success(c, 200, "")
	}
}

func UpdateServer(c *gin.Context) {
	serverId, _ := strconv.ParseInt(c.PostForm("server_id"), 10, 64)
	name := c.PostForm("name")
	path := c.PostForm("path")
	teamId, _ := strconv.ParseInt(c.PostForm("team_id"), 10, 64)
	groupId, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	policyId, _ := strconv.ParseInt(c.PostForm("policy_id"), 10, 64)
	monitorRegion := c.PostForm("monitor_region")
	serverFrom := ServerForm{
		Name:     name,
		Path:     path,
		TeamId:   teamId,
		GroupId:  groupId,
		PolicyId: policyId,
	}

	err := serverFrom.ValidateServerForm()
	if err != nil {
		response.Error(c, 4000, response.ConvertValidationErrorToString(err))
		return
	}
	models.UpdateServer(serverId, name, path, policyId, monitorRegion)
	response.Success(c, 200, "")
}

func GetServerList(c *gin.Context) {
	userId := c.GetInt64("userId")
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetServerListService(userId, groupId, page, pageSize)
	response.Success(c, 200, result)
}

func RefreshServerToken(c *gin.Context) {
	userId := c.GetInt64("userId")
	serverId, _ := strconv.ParseInt(c.PostForm("server_id"), 10, 64)
	token := getUniqueShorterStr()
	models.RefreshServerToken(userId, serverId, token)
	response.Success(c, 200, "")
}

func AddServerLog(c *gin.Context) {
	token := c.Param("token")
	serverData := models.GetServerIdByToken(token)
	serverLogItem := models.GetLatestServerLog(serverData.Id)
	createdTime := serverLogItem.CreatedAt.Format("2006-01-02 15:04:05")
	createdTimeLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", createdTime, time.Local)
	createdTimeLocation = createdTimeLocation.Add(+time.Second * time.Duration(serverData.Frequency))
	if time.Now().After(createdTimeLocation) {
		var checkSuccessInt int = 0
		cpuPercent, _ := strconv.ParseFloat(c.PostForm("cpu_percent"), 64)
		memoryPercent, _ := strconv.ParseFloat(c.PostForm("memory_used_percent"), 64)
		diskPercent, _ := strconv.ParseFloat(c.PostForm("disk_used_percent"), 64)
		if serverData.ServerCpu > cpuPercent {
			checkSuccessInt = 1
		}
		if serverData.ServerMemory > memoryPercent {
			checkSuccessInt = 1
		}
		if serverData.ServerDisk > diskPercent {
			checkSuccessInt = 1
		}
		data := models.ServerLogItem{
			ServerId:          serverData.Id,
			CpuUser:           c.PostForm("cpu_user"),
			CpuSystem:         c.PostForm("cpu_system"),
			CpuIdle:           c.PostForm("cpu_idle"),
			CpuPercent:        c.PostForm("cpu_percent"),
			MemoryTotal:       c.PostForm("memory_total"),
			MemoryAvailable:   c.PostForm("memory_available"),
			MemoryUsed:        c.PostForm("memory_used"),
			MemoryUsedPercent: c.PostForm("memory_used_percent"),
			DiskTotal:         c.PostForm("disk_total"),
			DiskFree:          c.PostForm("disk_free"),
			DiskUsed:          c.PostForm("disk_used"),
			DiskUsedPercent:   c.PostForm("disk_used_percent"),
			NetSent:           c.PostForm("net_sent"),
			NetRecv:           c.PostForm("net_recv"),
			Elapsed:           0,
			CheckSuccess:      checkSuccessInt,
		}
		models.AddServerLog(data)
		response.Success(c, 200, "")
	} else {
		response.Error(c, 900, "未到服务器监控时间")
	}

}

func DeleteServer(c *gin.Context) {
	serverId, _ := strconv.ParseInt(c.Query("server_id"), 10, 64)
	models.DeleteDevice(serverId, 2)
	response.Success(c, 200, "")
}

func GetFiledServerList(c *gin.Context) {
	userIdInt := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetFiledServerListService(userIdInt, page, pageSize)
	response.Success(c, 200, result)
}

func FileServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.FileDevice(id, 2)
	response.Success(c, 200, "")
}

func StartMonitorServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.StartMonitor(id, 2)
	response.Success(c, 200, "")
}

func PauseMonitorServer(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	models.PauseMonitor(id, 2)
	response.Success(c, 200, "")
}
