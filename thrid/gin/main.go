package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vearne/golib/buffpool"
)

type SimplebodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w SimplebodyWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func Timeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// sync.Pool
		buffer := buffpool.GetBuff()

		blw := &SimplebodyWriter{body: buffer, ResponseWriter: c.Writer}
		c.Writer = blw

		finish := make(chan struct{})
		go func() { // 子协程只会将返回数据写入到内存buff中
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-time.After(t):
			c.Writer.WriteHeader(http.StatusGatewayTimeout)
			c.Abort()
			// 1. 主协程超时退出。此时，子协程可能仍在运行
			// 如果超时的话，buffer无法主动清除，只能等待GC回收
		case <-finish:
			// 2. 返回结果只会在主协程中被写入
			blw.ResponseWriter.Write(buffer.Bytes())
			buffpool.PutBuff(buffer)
		}
	}
}

func main() {
	// create new gin without any middleware
	engine := gin.New()
	// add timeout middleware with 2 second duration
	engine.Use(Timeout(time.Second * 2))
	engine.GET("/", func(context *gin.Context) {
		time.Sleep(time.Second * 5)
		context.String(300, "ssss")
	})

	// run the server
	log.Fatal(engine.Run(":8080"))
}
