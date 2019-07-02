package main

import (
	"fmt"
	n "github.com/jinzhu/now"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p("现在时间", now)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("初始化一个值", then)
	p("年: ", then.Year())
	p("月: ", int(then.Month()))
	p("日: ", then.Day())
	p("时: ", then.Hour())
	p("分: ", then.Minute())
	p("秒: ", then.Second())
	p("毫微秒(十亿分之一秒): ", then.Nanosecond())
	p("地区: ", then.Location())
	p("星期几: ", int(then.Weekday()))
	p("时间在什么之前: ", then.Before(now))
	p("时间在什么之后: ", then.After(now))
	p("时间相同: ", then.Equal(now))
	diff := now.Sub(then)
	p("时间时间的差距", diff)
	fmt.Printf("差距的时%.2f\n", diff.Hours())
	fmt.Printf("差距的分%.2f\n", diff.Minutes())
	fmt.Printf("差距的秒%.2f\n", diff.Seconds())
	p("差距的毫微秒", diff.Nanoseconds())
	p("时间加上差距", then.Add(diff))
	p("时间减去差距", then.Add(-diff))

	//now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println("现在的unix秒数", secs)
	fmt.Println("现在的unix毫秒数", millis)
	fmt.Println("现在的unix毫微秒数", nanos)

	fmt.Println("通过秒和毫秒的相加值创建一个time", time.Unix(secs, 1209))
	fmt.Println(time.Unix(0, nanos))
	p("格式化")
	//格式化
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	p("通过一定的格式格式化时间字符串", t.Format("3:04pm"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	p(t.Format("2006-01-02 15:04:05"))

	p("解析")
	//解析时间
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
	p("今天的起始时间", n.BeginningOfDay())
	today := n.New(time.Now())
	p("今天的起始时间", today.BeginningOfDay())
	tomorrow := today.AddDate(0, 0, 1)
	tomorrowBegin := n.New(tomorrow).BeginningOfDay()
	p("明天的起始时间", tomorrowBegin)
	p("还有%s今天结束", today.Sub(tomorrowBegin))

}
