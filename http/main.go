package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, " +
		"like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
)

var (
	pl = log.Println
	pf = log.Printf
)

/**
 *
 */
func main() {
	url := "http://:8080"
	tranSport := &http.Transport{}
	// 新建客户端
	client := &http.Client{
		Transport: tranSport,
	}
	req, err := http.NewRequest("GET",
		url, nil)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	randIp := "192.123.123.1"
	req.Header.Set("X-Real-Ip", randIp)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		pf("req fail: %s", err)
		return
	}
	defer resp.Body.Close()
	contentLen := resp.ContentLength
	if resp.ContentLength == -1 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			pf("获取http body 失败,错误信息: %s", err)
			return
		}
		contentLen = int64(len(body))
	}
	pl("content length: ", contentLen)
}
