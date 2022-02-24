package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type MBDResultStruct struct {
	OrderTime   int64   `json:"ordertime"`
	OrderAmount float64 `json:"orderamount"`
	Payway      string  `json:"payway"`
	OrderId     string  `json:"orderid"`
	CreatorId   string  `json:"creatorid"`
	State       string  `json:"state"`
	ExpireAt    int64   `json:"expired_at"`
	Rounds      int     `json:"rounds"`
	RealAmount  float64 `json:"real_amount"`
	VersionName string  `json:"version_name"`
	UrlKey      string  `json:"urlkey"`
}

type MBDResponse struct {
	Code      int `json:"code"`
	Result    MBDResultStruct
	ErrorInfo string `json:"error_info"`
}

func OnGetOrderDetailApi(orderId string) MBDResponse {
	path := "https://x.mianbaoduo.com/api/order-detail?order_id=" + orderId
	client := &http.Client{}
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Set("x-token", os.Getenv("MBD_TOKEN"))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var respJson MBDResponse
	json.Unmarshal([]byte(string(body)), &respJson)
	return respJson
}
