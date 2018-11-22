package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//创建一个cmd命令 不知道为什么dir命令不在path 中
	dirCmd := exec.Command("dir")

	dateOut, err := dirCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("dir: ")
	fmt.Println(string(dateOut))
}
