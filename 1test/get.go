package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get() {
	//跳过证书验证
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	//生成要访问的url
	url := "https://test.buffge.com/t.php"
	//提交请求
	req, err := http.NewRequest("GET", url, nil)
	//增加header选项
	req.Header.Add("Cookie", "xxxxxx")
	req.Header.Add("User-Agent", "xxx")
	req.Header.Add("X-Requested-With", "xxxx")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	resp, _ := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("从body中读取数据失败")
		}
		for k, v := range resp.Header {
			fmt.Println(k, " : ", v)
		}
		fmt.Println(resp.ContentLength)
		fmt.Println(resp.Proto)
		fmt.Println(resp.ProtoMajor)
		fmt.Println(string(body))
	}
}
