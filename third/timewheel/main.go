package main

import (
	"github.com/ouqiang/timewheel"
	"log"
	"time"
)

func main() {
	tw := timewheel.New(1*time.Second, 60, func(data interface{}) {
		log.Println(data)
	})
	tw.Start()
	tw.AddTimer(time.Second, 1, 1)
	tw.AddTimer(time.Second, 2, 2)
	select {}
}
