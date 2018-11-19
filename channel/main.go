package main

import (
    "fmt"
    "time"
)

func main() {

    c0 := make(chan int)
    c1 := make(chan int)
    timeout := time.After(time.Second * 3)
    go func() {
        time.AfterFunc(time.Second*1, func() {
            c0 <- 0
        })
        time.AfterFunc(time.Second*2, func() {
            c1 <- 0
        })
    }()
    go func() {
        select {
        case <-c0:
            println("c0先收到值")
        case <-c1:
            println("c1先收到值")
        case <-timeout:
            println("已超时")
        }
    }()

    //缓冲通道
    c2 := make(chan int, 32)
    go func() {
        for i := 0; i < 32; i++ {
            c2 <- i
            time.Sleep(time.Millisecond * 10)
        }
        close(c2)
    }()
    //在这里会一直卡住,直到通道被关闭
    for v := range c2 {
        fmt.Println("接收到chan值", v)
    }

    //通道通知
    workerComplete := make(chan byte)
    workerTimeout := time.After(time.Millisecond * 3000)
    go func() {
        fmt.Println("执行一些耗时的操作,比如查询数据库,请求网络api等等")
        time.Sleep(time.Millisecond * 2500)
        workerComplete <- 0
    }()
    select {
    case <-workerComplete:
        fmt.Println("耗时操作完成了")
    case <-workerTimeout:
        fmt.Println("两秒还没执行完,操作超时了")
    }

    //单向通道
    ping := make(chan byte)
    pong := make(chan byte)
    go func(ping chan<- byte) {
        time.Sleep(time.Millisecond * 1000)
        ping <- 22
    }(ping)
    go func(ping <-chan byte, pong chan<- byte) {
        v := <-ping
        pong <- v
    }(ping, pong)
    select {
    case <-pong:
        fmt.Println("已接收到ping信息,即将发出pong信息")
    }
}
