package main

import (
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var client = http.Client{
	Transport: &http.Transport{
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}
var req *http.Request

func main() {
	var (
		url, durationStr, tickerDelayStr, timeoutStr string
		concurrent                                   int
	)
	flag.StringVar(&url, "url", "", "要访问的url")
	flag.IntVar(&concurrent, "c", 100, "并发数")
	flag.StringVar(&durationStr, "time", "30s", "压测的时间 eg. 30s 1h")
	flag.StringVar(&timeoutStr, "timeout", "30s", "请求超时时间 eg. 30s 1m")
	flag.StringVar(&tickerDelayStr, "ticker", "10s", "每隔多久显示当前压测情况 不填则不显示")
	flag.Parse()
	url = "http://localnor.com:8898/prom"
	concurrent = 10000
	durationStr = "60m"
	tickerDelayStr = "10s"
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}
	tickerDelay, err := time.ParseDuration(tickerDelayStr)
	if err != nil {
		panic(err)
	}
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		panic(err)
	}
	client.Timeout = timeout
	req, _ = http.NewRequest(http.MethodGet, url, nil)
	pressure(url, concurrent, duration, tickerDelay)
}
func pressure(url string, concurrency int, duration, tickerDelay time.Duration) {
	var successCount, failedCount, contentLen int
	quit := time.After(duration)
	begin := time.Now()
	log.Println("start pressure")
	contItemLen := make(chan int, concurrency)
	fail := make(chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		go spider(contItemLen, fail)
	}
	var outputContentLenTicker *time.Ticker
	if tickerDelay != 0 {
		outputContentLenTicker = time.NewTicker(tickerDelay)
		go func() {
			for t := range outputContentLenTicker.C {
				if contentLen == 0 {
					continue
				}
				duration := t.Sub(begin)
				lenMb := float64(contentLen / 1024 / 1024)
				durationSec := duration.Seconds()
				log.Printf(" ticker 共用时%s\n 成功%d次\n 失败%d次\n 总字节数%.2fM\n"+
					" 每秒字节数:%.1fM\n 每秒请求数: %.1f\n",
					duration, successCount, failedCount,
					lenMb,
					lenMb/durationSec,
					float64(successCount)/durationSec,
				)
			}
		}()
	}
	signQuit := make(chan os.Signal)
	signal.Notify(signQuit, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case t := <-contItemLen:
			successCount++
			contentLen += t
			go spider(contItemLen, fail)
		case <-fail:
			failedCount++
			go spider(contItemLen, fail)
		case <-quit:
			goto stop
		case <-signQuit:
			goto stop
		}
	}
stop:
	duration = time.Since(begin)
	lenMb := float64(contentLen / 1024 / 1024)
	durationSec := duration.Seconds()
	log.Printf("访问%s\n 共用时%v\n 成功%d次\n 失败%d次\n 总字节数%.1fM\n"+
		" 每秒字节数:%.1fM\n 每秒请求数: %.1f\n",
		url, duration, successCount, failedCount, lenMb,
		lenMb/durationSec,
		float64(successCount)/durationSec,
	)
	log.Println("end pressure")
	if outputContentLenTicker != nil {
		outputContentLenTicker.Stop()
	}
}

func spider(contItemLen chan int, fail chan struct{}) {
	resp, err := client.Do(req)
	if err != nil {
		fail <- struct{}{}
		return
	}
	defer resp.Body.Close()
	length, _ := io.Copy(ioutil.Discard, resp.Body)
	contItemLen <- int(length)
}
