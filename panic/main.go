package main

import "fmt"

/**
恐慌 有点类似于异常 抛出异常,然后用recover 捕获
*/
func main() {
	testDeferAndPanic()
	fmt.Println("下面会触发panic")
	panic("发生了一个恐慌")
	fmt.Println("这句话不会输出")
}

func testDeferAndPanic() {
	//就算发生了恐慌 defer 也会运行
	defer fmt.Println("this is defer")
	panic("panic happend")
}
