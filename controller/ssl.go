package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckSSL(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	webItem := models.GetWebDetail(id)
	service.CheckSSL(webItem)
	response.Success(c, 200, "")
}

func CheckWhois(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	webItem := models.GetWebDetail(id)
	service.CheckWhois(webItem)
	response.Success(c, 200, "")
}
