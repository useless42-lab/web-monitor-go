package routes

import (
	"WebMonitor/controller"
	"WebMonitor/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func InitApiRoute() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	v1 := router.Group("/v1")
	{
		v1.GET("/monitor/heartbeat/:token", controller.AddHeartbeatLog)
		v1.GET("/team/transfer", controller.GetTransferTeamInfo)
		v1.GET("/team/invite", controller.GetInviteTeamMemberInfo)
		v1.POST("/monitor/server/:token", controller.AddServerLog)
		v1.GET("/view_status_page", controller.GetStatusPageDetail)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
			auth.POST("/register", controller.Register)
			auth.GET("/captcha", controller.GenerateCaptchaHandler)
			auth.POST("/reset/link", controller.GenerateResetPasswordLink)
			auth.POST("/reset", controller.ResetPassword)
		}
		user := v1.Group("/user")
		user.Use(middleware.CheckAuthMiddleware())
		{
			mobile := user.Group("/mobile")
			{
				mobile.GET("/device/list", controller.GetMobileDeviceList)
			}
			device := user.Group("/device")
			{
				device.GET("/simple/list", controller.GetAllDeviceList)
				device.POST("/group/transfer", controller.TransferDeviceGroup)
				device.GET("/detail", controller.GetDeviceDetail)
			}
			statusPage := user.Group("/status_page")
			{
				statusPage.POST("/add", controller.AddStatusPage)
				statusPage.GET("/list", controller.GetStatusPageList)
				statusPage.GET("/detail", controller.GetUserStatusPageDetail)
				statusPage.POST("/update", controller.UpdateStatusPage)
				statusPage.POST("/delete", controller.DeleteStatusPage)
			}
			web := user.Group("/web")
			{
				web.POST("/create", controller.AddWeb)
				web.GET("/list", controller.GetWebList)
				web.GET("/ssl", controller.GetWebSSLConfig)
				web.GET("/whois", controller.GetDomainWhois)
				web.GET("/overview", controller.GetWebOverView)
				web.DELETE("/delete", controller.DeleteWeb)
				web.GET("/filed/list", controller.GetFiledWebList)
				web.POST("/file", controller.FileWeb)
				web.POST("/monitor/start", controller.StartMonitorWeb)
				web.POST("/ssl/check", controller.CheckSSL)
				web.POST("/whois/check", controller.CheckWhois)
				web.POST("/pause", controller.PauseMonitorWeb)
				web.POST("/detail", controller.UpdateWeb)
			}
			server := user.Group("/server")
			{
				server.POST("/create", controller.AddServer)
				server.GET("/list", controller.GetServerList)
				server.POST("/token/refresh", controller.RefreshServerToken)
				server.GET("/log", controller.GetServerLog)
				server.DELETE("/delete", controller.DeleteServer)
				server.GET("/filed/list", controller.GetFiledServerList)
				server.POST("/file", controller.FileServer)
				server.POST("/monitor/start", controller.StartMonitorServer)
				server.POST("/pause", controller.PauseMonitorServer)
				server.POST("/detail", controller.UpdateServer)
			}
			api := user.Group("/api")
			{
				api.POST("/create", controller.AddApi)
				api.GET("/list", controller.GetApiList)
				api.POST("/update", controller.UpdateApi)
				api.GET("/detail", controller.GetApi)
				api.GET("/mock", controller.MockApi)
				api.GET("/filed/list", controller.GetFiledApiList)
				api.POST("/file", controller.FileApi)
				api.POST("/monitor/start", controller.StartMonitorApi)
				api.POST("/pause", controller.PauseMonitorApi)
				api.DELETE("/delete", controller.DeleteApi)
			}
			tcp := user.Group("tcp")
			{
				tcp.POST("/create", controller.AddTcp)
				tcp.GET("/list", controller.GetTcpList)
				tcp.GET("/filed/list", controller.GetFiledTcpList)
				tcp.POST("/file", controller.FileTcp)
				tcp.POST("/monitor/start", controller.StartMonitorTcp)
				tcp.POST("/pause", controller.PauseMonitorTcp)
				tcp.DELETE("/delete", controller.DeleteTcp)
				tcp.POST("/detail", controller.UpdateTcp)
			}
			dns := user.Group("dns")
			{
				dns.POST("/create", controller.AddDns)
				dns.GET("/list", controller.GetDnsList)
				dns.GET("/filed/list", controller.GetFiledDnsList)
				dns.POST("/file", controller.FileDns)
				dns.POST("/monitor/start", controller.StartMonitorDns)
				dns.POST("/pause", controller.PauseMonitorDns)
				dns.DELETE("/delete", controller.DeleteDns)
				dns.POST("/detail", controller.UpdateDns)
			}
			heartbeat := user.Group("heartbeat")
			{
				heartbeat.POST("/create", controller.AddHeartbeat)
				heartbeat.GET("/list", controller.GetHeartbeatList)
				heartbeat.GET("/filed/list", controller.GetFiledHeartbeatList)
				heartbeat.POST("/file", controller.FileHeartbeat)
				heartbeat.POST("/monitor/start", controller.StartMonitorHeartbeat)
				heartbeat.POST("/pause", controller.PauseMonitorHeartbeat)
				heartbeat.DELETE("/delete", controller.DeleteHeartbeat)
				heartbeat.POST("/detail", controller.UpdateHeartbeat)
			}
			steam := user.Group("steam")
			{
				steam.POST("/create", controller.AddSteamServer)
				steam.GET("/list", controller.GetSteamServerList)
				steam.GET("/filed/list", controller.GetFiledSteamServerList)
				steam.POST("/file", controller.FileSteamServer)
				steam.POST("/monitor/start", controller.StartMonitorSteamServer)
				steam.POST("/pause", controller.PauseMonitorSteamServer)
				steam.DELETE("/delete", controller.DeleteSteam)
				steam.POST("/detail", controller.UpdateSteamServer)
			}
			minecraft := user.Group("minecraft")
			{
				minecraft.POST("/create", controller.AddMinecraftServer)
				minecraft.GET("/list", controller.GetMinecraftServerList)
				minecraft.GET("/filed/list", controller.GetFiledMinecraftServerList)
				minecraft.POST("/file", controller.FileMinecraftServer)
				minecraft.POST("/monitor/start", controller.StartMonitorMinecraftServer)
				minecraft.POST("/pause", controller.PauseMonitorMinecraftServer)
				minecraft.DELETE("/delete", controller.DeleteMinecraft)
				minecraft.POST("/detail", controller.UpdateMinecraftServer)
			}
			team := user.Group("/team")
			{
				team.GET("/info", controller.GetTeamGroupInfo)
				team.GET("/list", controller.GetTeamGroupList)
				team.POST("/create", controller.AddTeamGroup)
				team.GET("/detail", controller.GetTeamGroupDetail)
				team.PATCH("/update", controller.UpdateTeamGroup)
				team.DELETE("/delete", controller.DeleteTeamGroup)
				member := team.Group("/member")
				{
					member.GET("/list", controller.GetTeamGroupMemberList)
					member.POST("/exit", controller.ExitTeam)
					member.POST("/kick", controller.KickOutTeamMember)

					// 生成转让团队链接
					member.POST("/generate/transfer", controller.GenerateTransferTeamLink)
					// 生成邀请成员链接
					member.POST("/generate/invite", controller.GenerateInviteTeamMemberLink)
					//
					member.POST("/admin/remove", controller.RemoveAdmin)
					member.POST("/admin/add", controller.AddAdmin)
				}
				team.POST("/invite", controller.CreateInviteTeamMember)
				team.POST("/transfer", controller.TransferTeamGroup)
			}
			group := user.Group("/group")
			{
				group.GET("/list", controller.GetDeviceGroup)
				group.GET("/list/type", controller.GetDeviceGroupType)
				group.GET("/list/pagination", controller.GetDeviceGroupPaginationList)
				group.GET("/detail", controller.GetDeviceGroupDetail)
				group.PATCH("/update", controller.UpdateDeviceGroup)
				group.PATCH("/update/name", controller.UpdateDeviceGroupName)
				group.DELETE("/delete", controller.DeleteGroup)
				group.POST("/create", controller.AddDeviceGroup)
				group.GET("/notification/list", controller.GetNotificationList)
				group.GET("/team/member/base/list", controller.GetTeamMemberListByGroupId)
				group.POST("/notification/list/add", controller.AddNotificationListItem)
				group.POST("/notification/list/delete", controller.DeleteNotificationListItem)
			}
			policy := user.Group("/policy")
			{
				policy.GET("/list", controller.GetMonitorPolicyList)
				policy.GET("/list/pagination", controller.GetMonitorPolicyPaginationList)
				policy.GET("/detail", controller.GetMonitorPolicy)
				policy.DELETE("/delete", controller.DeleteMonitorPolicy)
				policy.POST("/create", controller.AddMonitorPolicy)
				policy.PATCH("/update", controller.UpdateMonitorPolicy)
				policy.GET("/frequency", controller.GetPlanFrequencyByUserId)
			}
			userAuth := user.Group("/auth")
			{
				userAuth.POST("/password/change", controller.ChangePassword)
				userAuth.GET("/base", controller.GetBaseUserInfo)
				userAuth.GET("/third", controller.GetUserThirdPartyInfo)
				userAuth.POST("/third", controller.UpdateUserThirdPartyInfo)
			}
			plan := user.Group("/plan")
			{
				plan.GET("/config", controller.GetPlanBaseInfo)
				plan.GET("/list", controller.GetPlanBaseList)
				plan.GET("/user", controller.GetUserPlanDetail)
			}
			order := user.Group("/order")
			{
				order.POST("/add", controller.AddOrder)
				order.GET("/list", controller.GetOrderList)
			}
			notification := user.Group("notification")
			{
				notification.GET("/detail", controller.GetNotification)
				notification.POST("/update", controller.UpdateNotification)
			}
		}
	}
	router.Run(":" + os.Getenv("ROUTE_PORT"))
}
