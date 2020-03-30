package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	log.Println(time.Now())
	// 分钟 小时 天 月(1-12) 星期(0-6)
	c.AddFunc("*/1 * * * *", func() { fmt.Println("每分钟执行一次A", time.Now()) })
	c.AddFunc("49-52 * * * *", func() { fmt.Println("每小时的49-52分都执行一次", time.Now()) })
	// c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	c.AddFunc("CRON_TZ=Asia/Shanghai 50 20 * * *", func() { fmt.Println("每天20:50 上海时间运行") })
	c.AddFunc("*/2 * * * *", func() { fmt.Println("每两分钟运行一次", time.Now()) })
	c.AddFunc("@every 3m", func() { fmt.Println("从现在开始每三分钟运行一次", time.Now()) })
	c.Start()
	defer c.Stop()
	time.Sleep(1 * time.Hour)
}
