package main

import (
	"fmt"
	"runtime"
)

var (
	pl = fmt.Println
)

//
func main() {
	pl("cpu 逻辑核心数: ", runtime.NumCPU())
	pl("系统: ", runtime.GOOS)
}
