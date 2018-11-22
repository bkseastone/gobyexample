package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//设置环境变量
	os.Setenv("FOO", "1")
	//读取环境变量
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	//读取系统的所有环境变量
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
}
