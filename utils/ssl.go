package utils

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func GetSSLInfo(seedUrl string) (startTime time.Time, endTime time.Time, subject string, issuer string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(seedUrl)

	if err != nil {
		fmt.Errorf(seedUrl, " 请求失败")
		// panic(err)
	}
	if resp != nil {
		if resp.TLS != nil {
			defer resp.Body.Close()
			certInfo := resp.TLS.PeerCertificates[0]
			// fmt.Println("开始时间:", certInfo.NotBefore)
			// fmt.Println("过期时间:", certInfo.NotAfter)

			// fmt.Println("ServerName:", resp.TLS.ServerName)
			// fmt.Println("组织信息:", certInfo.Subject.Names)

			// fmt.Println("颁发给:", certInfo.Subject.Names[len(certInfo.Subject.Names)-1].Value)
			// fmt.Println("颁发者:", certInfo.Issuer.Names[len(certInfo.Issuer.Names)-1].Value)
			return certInfo.NotBefore, certInfo.NotAfter, fmt.Sprintf("%v", certInfo.Subject.Names[len(certInfo.Subject.Names)-1].Value), fmt.Sprintf("%v", certInfo.Issuer.Names[len(certInfo.Issuer.Names)-1].Value)
		} else {
			return time.Now(), time.Now(), "0", "0"
		}
	} else {
		return time.Now(), time.Now(), "0", "0"
	}
}
