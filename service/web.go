package service

import (
	"WebMonitor/models"
	"WebMonitor/utils"
	"fmt"
	"strings"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func GetWebListService(userId int64, groupId int64, page int, pageSize int) models.PaginationData {
	data := models.GetWebList(userId, groupId, page, pageSize)
	total := models.GetDeviceCount(userId, groupId, 1)
	result := models.PaginationData{
		Data:      data,
		Total:     total.Total,
		PageIndex: page,
		PageSize:  pageSize,
	}
	return result
}

/*
检查是否能插入设备
*/
func CheckCanAddDevice(teamId int64) bool {
	planId := models.GetPlanIdByTeamId(teamId)
	plan := models.GetPlanBaseInfo(planId.Id)
	deviceNumber := models.GetAllDeviceNumber(teamId)
	if deviceNumber.Total < plan.PerTeamDeviceLimit {
		return true
	} else {
		return false
	}
}

func AddWebService(userId int64, name string, path string, teamId int64, groupId int64, policyId int64, basicUser string, basicPassword string, monitorRegion string) string {
	addStatus := CheckCanAddDevice(teamId)
	if addStatus {
		web := models.AddWeb(userId, name, path, groupId, policyId, basicUser, basicPassword, monitorRegion)
		go AddSSLConfig(web)
		go AddDomainWhois(web)
		return ""
	} else {
		return "设备数量超出上限"
	}
}

func AddDomainWhois(web models.WebItem) {
	webPathArr := strings.Split(web.Path, ".")
	if len(webPathArr) >= 2 {
		targetPath := webPathArr[len(webPathArr)-2] + "." + webPathArr[len(webPathArr)-1]
		whoisRaw, err := whois.Whois(targetPath)
		if err == nil {
			result, err := whoisparser.Parse(whoisRaw)
			if err == nil {
				tCreatedDate, _ := time.Parse(time.RFC3339, result.Domain.CreatedDate)
				tExpirationDate, _ := time.Parse(time.RFC3339, result.Domain.ExpirationDate)
				models.AddDomainWhois(web.DefaultModel.ID, tCreatedDate, tExpirationDate, result.Registrar.Name, result.Registrant.Name, result.Registrant.Email)
			}
		} else {
			fmt.Println(err)
		}
	}

}

func AddSSLConfig(web models.WebItem) {
	startTime, endTime, subject, issuer := utils.GetSSLInfo(web.Path)
	models.AddSSLConfig(web.DefaultModel.ID, startTime, endTime, subject, issuer)
}

func CheckSSL(web models.WebItem) {
	go OperateSSLConfig(web)
}

func OperateSSLConfig(web models.WebItem) {
	startTime, endTime, subject, issuer := utils.GetSSLInfo(web.Path)
	models.CheckSSLConfig(web.DefaultModel.ID, startTime, endTime, subject, issuer)
}

func CheckWhois(web models.WebItem) {
	go OperateWhoisConfig(web)
}

func OperateWhoisConfig(web models.WebItem) {
	webPathArr := strings.Split(web.Path, ".")
	if len(webPathArr) >= 2 {
		targetPath := webPathArr[len(webPathArr)-2] + "." + webPathArr[len(webPathArr)-1]
		whoisRaw, err := whois.Whois(targetPath)
		if err == nil {
			result, err := whoisparser.Parse(whoisRaw)
			if err == nil {
				tCreatedDate, _ := time.Parse(time.RFC3339, result.Domain.CreatedDate)
				tExpirationDate, _ := time.Parse(time.RFC3339, result.Domain.ExpirationDate)
				models.CheckWhois(web.DefaultModel.ID, tCreatedDate, tExpirationDate, result.Registrar.Name, result.Registrant.Name, result.Registrant.Email)
			}
		} else {
			fmt.Println(err)
		}
	}
}
