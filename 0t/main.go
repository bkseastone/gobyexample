package main

import (
	"fmt"
	"time"

	. "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Print("\x1b[35mbuffge123\x1b[0m")
	fmt.Printf("PI is %+1.2e\n", Cyan(3.14))
	fmt.Print("buffge222")
	<-time.After(time.Second)
	fmt.Print("\x1b[2\x00\x00\x00A\x1b[J")
	<-time.After(time.Hour)

}
