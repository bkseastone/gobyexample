package main

import (
    "fmt"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func main() {

    f("同步程序")
    //立即开启一个协程,开好之后立即执行,不用等到后面同步程序执行完再执行
    go f("协程")
    go func(msg string) {
        fmt.Println(msg)
    }("再开一个协程")
    fmt.Println("这是第二个同步程序")
    for i := 0; i < 2000000000; i++ {
        if i == 2000000000-1 {
            //这句话会在协程之后输出
            fmt.Println("同步程序已循环20亿次")
        }
    }
    fmt.Scanln()
    fmt.Println("done")
}
