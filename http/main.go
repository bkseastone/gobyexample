package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	var (
		url string
	)
	flag.StringVar(&url, "url", "http://baidu.com", "要访问的url")
	// 新建客户端
	client := &http.Client{}
	req, err := http.NewRequest("GET",
		url, nil)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	randIp := strconv.Itoa(rand.Intn(221-58)+58) + "." +
		strconv.Itoa(rand.Intn(0xff)) + "." +
		strconv.Itoa(rand.Intn(0xff)) + "." +
		strconv.Itoa(rand.Intn(0xff))
	req.Header.Set("X-FORWARDED-FOR", randIp)
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
