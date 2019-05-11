package main

import (
	"fmt"
	"os"
)

/**
 * 进程相关
 */
func main() {
	fmt.Println("当前pid = ", os.Getpid())
	fmt.Scanln()
}
