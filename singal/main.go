package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	//新建一个提示,如果接受到系统的(退出)ctrl+c  或者 kill TERM(中断)
	//如果是kill 9 将会强制终止 程序无法捕捉信号, 还有SIGTSTP ctrl+z 也是无法捕捉到的
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("接受到了信号", sig)
		done <- true
	}()
	fmt.Println("持续等待信号")
	<-done
	fmt.Println("退出")
}
