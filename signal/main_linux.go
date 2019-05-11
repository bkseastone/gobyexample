package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// 新建一个提示,如果接受到系统的(退出)ctrl+c  或者 kill TERM(中断)
	// 如果是kill 9 将会强制终止 程序无法捕捉信号, 还有SIGTSTP ctrl+z 也是无法捕捉到的
	// SIGINT : ctrl + c
	// SIGTERM : kill <pic>
	//  1:  "hangup",       挂起
	// 	2:  "interrupt",    ctrl + c
	// 	3:  "quit",         ctrl + \ 这个在win下不成功
	// 	4:  "illegal instruction", 非法指令  不懂
	// 	5:  "trace/breakpoint trap",
	// 	6:  "aborted",
	// 	7:  "bus error",
	// 	8:  "floating point exception",
	// 	9:  "killed",
	// 	10: "user defined signal 1", // linux 下才有 kill -
	// 	11: "segmentation fault",
	// 	12: "user defined signal 2",
	// 	13: "broken pipe",
	// 	14: "alarm clock",
	// 	15: "terminated",
	sysType := runtime.GOOS

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2,
	)
	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGINT:
			fmt.Println("接收到ctrl + c 信号")
		case syscall.SIGTERM:
			fmt.Println("接收到kill 信号")
		case syscall.SIGQUIT:
			fmt.Println("接收到ctrl + \\ 信号")
		case syscall.SIGUSR1:
			fmt.Println("接收到kill -SIGUSR1 <pid> 信号")
		case syscall.SIGUSR2:
			fmt.Println("接收到kill -SIGUSR2 <pid> 信号")
		}
		fmt.Println("接受到了信号", sig)
		done <- true
	}()
	fmt.Println("持续等待信号")
	<-done
	fmt.Println("退出")
}
