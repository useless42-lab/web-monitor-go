package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func MonitorApi(path string, method int, headers string, bodyType int, requestData string, responseData string, basicUser string, basicPassword string) (status string, statusCode int, body string, proto string, elapsed int64) {
	t := time.Now()
	type JsonItem struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	// // get delete
	// payload := nil
	// // post
	// payload := strings.NewReader("a=111")
	// // post form
	// payload1 := url.Values{"key": {"value"}, "id": {"123"}}
	// // post json/put/patch
	// payload := strings.NewReader("{\"name\":\"123\"}")

	var methodStr string = "GET"
	payload := strings.NewReader("")
	urlValues := url.Values{}

	switch method {
	case 1:
		methodStr = "GET"
		requestData = ""
	case 2:
		methodStr = "POST"
		if bodyType == 3 {
			var requestDataBlob = []byte(requestData)
			var requestDataContent []JsonItem
			err := json.Unmarshal(requestDataBlob, &requestDataContent)
			if err != nil {
				fmt.Println("error:", err)
			}
			for _, v := range requestDataContent {
				urlValues.Set(v.Name, v.Value)
			}
			requestData = urlValues.Encode()
		}
	case 3:
		methodStr = "PUT"
	case 4:
		methodStr = "PATCH"
	case 5:
		methodStr = "DELETE"
	case 6:
		methodStr = "HEAD"
	case 7:
		methodStr = "OPTIONS"
	default:
		methodStr = "GET"
		requestData = ""
	}
	payload = strings.NewReader(requestData)
	req, _ := http.NewRequest(methodStr, path, payload)
	if method != 1 {
		if bodyType == 3 {
			// post form
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		} else {
			// post json/put/patch
			req.Header.Add("Content-Type", "application/json")
		}
	}

	if basicUser == "" && basicPassword == "" {
		req.SetBasicAuth(basicUser, basicPassword)
	}

	if method != 1 {
		var headersBlob = []byte(headers)
		var headersContent []JsonItem
		err := json.Unmarshal(headersBlob, &headersContent)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			for _, v := range headersContent {
				req.Header.Add(v.Name, v.Value)
			}
		}
	}

	response, err1 := http.DefaultClient.Do(req)
	fmt.Println("path", path)
	fmt.Println("err1", err1)
	if err1 == nil {
		status = response.Status
		statusCode = response.StatusCode
		proto = response.Proto
		defer response.Body.Close()
	} else {
		status = ""
		statusCode = -1
		proto = ""
	}
	var bodyStr string
	if statusCode != -1 {
		body, _ := ioutil.ReadAll(response.Body)
		bodyStr = string(body)
	} else {
		bodyStr = ""
	}

	elapsedTime := time.Since(t)

	return status, statusCode, bodyStr, proto, int64(elapsedTime)
}
