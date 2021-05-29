package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func a() {
	for {
		_ = 1
	}
}
func main() {
	go http.ListenAndServe(":6060", nil)
	for i := 0; i < 6; i++ {
		go a()
	}
	time.Sleep(time.Hour)
}
