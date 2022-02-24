package models

type OrderItemStruct struct {
	DefaultModel
	UserId         int64   `json:"user_id" gorm:"column:user_id"`
	OrderTime      int64   `json:"order_time" gorm:"column:order_time"`
	OrderAmount    float64 `json:"order_amount" gorm:"column:order_amount"`
	State          string  `json:"state" gorm:"column:state"`
	Payway         string  `json:"payway" gorm:"column:payway"`
	OrderId        string  `json:"order_id" gorm:"column:order_id"`
	UrlKey         string  `json:"url_key" gorm:"column:url_key"`
	TargetPlanId   int     `json:"target_plan_id" gorm:"column:target_plan_id"`
	TargetPlanTime int     `json:"target_plan_time" gorm:"column:target_plan_time"`
	Status         int     `json:"status" gorm:"column:status"`
}

type ROrderItemStruct struct {
	UserId         int64     `json:"user_id" gorm:"column:user_id"`
	OrderTime      int64     `json:"order_time" gorm:"column:order_time"`
	OrderAmount    float64   `json:"order_amount" gorm:"column:order_amount"`
	State          string    `json:"state" gorm:"column:state"`
	Payway         string    `json:"payway" gorm:"column:payway"`
	OrderId        string    `json:"order_id" gorm:"column:order_id"`
	UrlKey         string    `json:"url_key" gorm:"column:url_key"`
	TargetPlanId   int       `json:"target_plan_id" gorm:"column:target_plan_id"`
	TargetPlanTime int       `json:"target_plan_time" gorm:"column:target_plan_time"`
	Status         int       `json:"status" gorm:"column:status"`
	CreatedAt      LocalTime `json:"created_at" gorm:"column:created_at"`
}

type RSimpleOrderItem struct {
	OrderId   string    `json:"order_id" gorm:"column:order_id"`
	Status    int       `json:"status" gorm:"column:status"`
	CreatedAt LocalTime `json:"created_at" gorm:"column:created_at"`
}

func AddOrder(userId int64, orderTime int64, orderAmount float64, state string, payway string, orderId string, urlKey string, targetPlanId int, targetPlanTime int, status int) OrderItemStruct {
	order := OrderItemStruct{
		UserId:         userId,
		OrderTime:      orderTime,
		OrderAmount:    orderAmount,
		State:          state,
		Payway:         payway,
		OrderId:        orderId,
		UrlKey:         urlKey,
		TargetPlanId:   targetPlanId,
		TargetPlanTime: targetPlanTime,
		Status:         status,
	}
	DB.Table("order_list").Create(&order)
	return order
}

func GetOrderList(userId int64) []RSimpleOrderItem {
	var result []RSimpleOrderItem
	sqlStr := `select order_id,status,created_at from order_list where user_id=@userId and status=1`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

func GetOrderCount(userId int64) RPTotal {
	var result RPTotal
	sqlStr := `select count(id) as total from order_list where user_id=@userId and status=1`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}

func IsExistOrder(userId int64, mbdOrderId string) bool {
	sqlStr := `select order_id from order_list where user_id=@userId and order_id=@mbdOrderId`
	var result RSimpleOrderItem
	DB.Raw(sqlStr, map[string]interface{}{
		"userId":     userId,
		"mbdOrderId": mbdOrderId,
	}).Scan(&result)
	if result.OrderId == "" {
		return false
	} else {
		return true
	}
}

func UseOrder(orderListId int64, userId int64) {
	sqlStr := `
	update order_list set status=0 where id=@orderListId and user_id=@userId
	`
	DB.Exec(sqlStr, map[string]interface{}{
		"orderListId": orderListId,
		"userId":      userId,
	})
}

func GetUserOrder(userId int64) OrderItemStruct {
	var result OrderItemStruct
	sqlStr := `select * from order_list where user_id=@userId and status=1 limit 1`
	DB.Raw(sqlStr, map[string]interface{}{
		"userId": userId,
	}).Scan(&result)
	return result
}
