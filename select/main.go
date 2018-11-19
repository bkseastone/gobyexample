package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
    for i := 0; i < 2; i++ {
        //select 会堵塞 直到收到任意的一个通道
        //select 中只能有通道
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        default:
            println("default")
        }
    }
    /**
        非堵塞通道 就是加个default,不等待直接执行default
     */
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }
    /**
    这三行代码执行会报错
    我猜是因为第二行代码因为一直没人接收它,所以第三行代码不会执行,导致进程死锁
    msg := "hi"
    messages <- msg
    println(<-messages)
     */
    msg := "hi"
    select {
    case <-messages:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }

}
