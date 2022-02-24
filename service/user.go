package service

import (
	"WebMonitor/models"
	"time"
)

func InitUserService(userId int64, username string, email string, expiredAt time.Time) {
	models.AddUser(userId, username, 2, expiredAt)
	models.AddNotification(userId, email, "", "", "", "", "")
	InitNewTeamGroupService(userId)
}
