package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	var seedUrl string
	flag.StringVar(&seedUrl, "url", "https://www.baidu.com", "默认Baidu首页")
	flag.Parse()
	resp, err := client.Get(seedUrl)
	defer resp.Body.Close()

	if err != nil {
		fmt.Errorf(seedUrl, "请求失败，请检查域名是否正确且存在")
		panic(err)
	}

	certInfo := resp.TLS.PeerCertificates[0]
	fmt.Println("域名证书过期时间：", certInfo.NotAfter)
	fmt.Println("域名证书组织信息：", certInfo.Subject)
	fmt.Println("证书颁发机构信息", certInfo.Issuer)
}
