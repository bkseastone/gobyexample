package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var trash *os.File

func main() {
	var err error
	if trash, err = os.OpenFile(os.DevNull, os.O_WRONLY, 0200); err != nil {
		panic(err)
	}
	var (
		url, durationStr, tickerDelayStr string
		concurrent                       int
	)
	flag.StringVar(&url, "url", "", "要访问的url")
	flag.IntVar(&concurrent, "c", 100, "并发数")
	flag.StringVar(&durationStr, "time", "30s", "压测的时间 eg. 30s 1h")
	flag.StringVar(&tickerDelayStr, "ticker", "10s", "每隔多久显示当前压测情况 不填则不显示")
	flag.Parse()
	// url = "https://test.buffge.com/yarn.lock"
	// url = "https://nqhjdadn.zrbwlqa.com:10001/v1/users/register"
	// url = "http://chaojiappguanli.oss-accelerate.aliyuncs.com/apk/f7b632eea5b64b08b09aa6e44b1bfaa1/f7b632eea5b64b08b09aa6e44b1bfaa1.apk"
	// concurrent = 2000
	// durationStr = "60m"
	// tickerDelayStr = "10s"
	pressure(url, concurrent, durationStr, tickerDelayStr)
}
func pressure(url string, concurrency int, durationStr, tickerDelayStr string) {
	var successCount, failedCount, contentLen int
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}
	tickerDelay, err := time.ParseDuration(tickerDelayStr)
	if err != nil {
		panic(err)
	}
	quit := time.After(duration)
	begin := time.Now()
	log.Println("start pressure")
	contItemLen := make(chan int, concurrency)
	fail := make(chan byte, concurrency)
	for i := 0; i < concurrency; i++ {
		go spider(url, contItemLen, fail)
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
				log.Printf(" ticker 共用时%s\n 成功%d次\n 失败%d次\n 总字节数%.2fM\n 每秒字节数:%.1fM\n 每秒请求数: %."+
					"1f\n",
					duration, successCount, failedCount,
					float32(contentLen/1024/1024),
					float64(contentLen/1024/1024)/duration.Seconds(),
					float64(successCount)/duration.Seconds(),
				)
			}
		}()
	}
	signQuit := make(chan os.Signal, 1)
	signal.Notify(signQuit, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case t := <-contItemLen:
			successCount++
			contentLen += t
			go spider(url, contItemLen, fail)
		case <-fail:
			failedCount++
			go spider(url, contItemLen, fail)
		case <-quit:
			goto stop
		case <-signQuit:
			goto stop
		}
	}
stop:
	end := time.Now()
	duration = end.Sub(begin)
	log.Printf("访问%s\n 共用时%v\n 成功%d次\n 失败%d次\n 总字节数%.1fM\n 每秒字节数:%.1fM\n 每秒请求数: %.1f\n",
		url, duration, successCount, failedCount, float32(contentLen/1024/1024),
		float64(contentLen/1024/1024)/duration.Seconds(),
		float64(successCount)/duration.Seconds(),
	)
	log.Println("end pressure")
	if outputContentLenTicker != nil {
		outputContentLenTicker.Stop()
	}
}

var client = http.Client{
	Transport: &http.Transport{
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func spider(url string, contItemLen chan int, fail chan byte) {
	resp, err := client.Get(url)
	if err != nil {
		fail <- 0
		return
	}
	defer resp.Body.Close()
	length, _ := trash.ReadFrom(resp.Body)
	// log.Println(err)
	contItemLen <- int(length)
}
