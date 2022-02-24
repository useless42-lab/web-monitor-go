package controller

import (
	"WebMonitor/models"
	"WebMonitor/response"
	"WebMonitor/service"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrderList(c *gin.Context) {
	userId := c.GetInt64("userId")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	result := service.GetOrderListService(userId, page, pageSize)
	response.Success(c, 200, result)
}

type PlanJsonStruct struct {
	Data []PlanDataJsonstruct
}

type PlanDataJsonstruct struct {
	UrlKey string `json:"urlkey"`
	PlanId int    `json:"plan_id"`
	Time   int    `json:"time"`
}

func AddOrder(c *gin.Context) {
	userId := c.GetInt64("userId")
	orderId := c.PostForm("order_id")
	mbdDetail := OnGetOrderDetailApi(orderId)

	filePtr, err := os.Open("plan.json")
	if err != nil {
		return
	}
	defer filePtr.Close()

	var respJson PlanJsonStruct
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&respJson)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
	} else {
		var targetPlanId, targetPlanTime int
		for _, item := range respJson.Data {
			if item.UrlKey == mbdDetail.Result.UrlKey {
				targetPlanId = item.PlanId
				targetPlanTime = item.Time
			}
		}
		if mbdDetail.Code == 200 {
			if mbdDetail.Result.State == "success" {
				if targetPlanId == 0 {
					response.Error(c, 2011, "订单号有误")
				} else {
					isExist := models.IsExistOrder(userId, mbdDetail.Result.OrderId)
					if isExist {
						response.Error(c, 2001, "该订单已存在")
					} else {
						userPlan := models.GetUserPlan(userId)
						if userPlan.PlanId == 1 {
							models.AddOrder(userId, mbdDetail.Result.OrderTime, mbdDetail.Result.OrderAmount, mbdDetail.Result.State, mbdDetail.Result.Payway, mbdDetail.Result.OrderId, mbdDetail.Result.UrlKey, targetPlanId, targetPlanTime, 0)
							models.UpdateUserPlan(userId, targetPlanId, time.Now().AddDate(0, targetPlanTime, 0))
						} else {
							if time.Now().After(userPlan.ExpiredAt) {
								models.AddOrder(userId, mbdDetail.Result.OrderTime, mbdDetail.Result.OrderAmount, mbdDetail.Result.State, mbdDetail.Result.Payway, mbdDetail.Result.OrderId, mbdDetail.Result.UrlKey, targetPlanId, targetPlanTime, 0)
								models.UpdateUserPlan(userId, targetPlanId, time.Now().AddDate(0, targetPlanTime, 0))
							} else {
								models.AddOrder(userId, mbdDetail.Result.OrderTime, mbdDetail.Result.OrderAmount, mbdDetail.Result.State, mbdDetail.Result.Payway, mbdDetail.Result.OrderId, mbdDetail.Result.UrlKey, targetPlanId, targetPlanTime, 1)
							}
						}
						response.Success(c, 200, "")
					}
				}
			}
		} else {
			response.Error(c, 2010, "订单号有误")
		}
	}
}
