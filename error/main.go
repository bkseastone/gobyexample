package main

import (
	"errors"
	"fmt"
	"log"

	gherr "github.com/pkg/errors"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

/**
定义一个自定义错误类型
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("参数是%d - 报错信息:%s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		// 因为是一个指针实现了Error接口,所以这里要取引用
		// argError并么有实现error接口,但是&argError 实现了
		return -1, &argError{arg, "参数不能是42"}
	}
	return arg + 3, nil
}
func f3() error {
	return f4()
}
func f4() error {
	err := &argError{222, "test 222"}
	return gherr.Wrap(err, "buffge's warp info")
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("遇到了错误: ", e)
		} else {
			fmt.Println("参数没有问题,结果是: ", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	// 判断错误类型 可以当捕捉异常用
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	_, e2 := f2(42)
	// 这里的e2是一个实现了error接口的类型
	// 显然这个类型是一个指向argError指针
	// 并不是argError 就是一个error类型
	// argError 可以改为任何名字
	switch e2.(type) {
	case *argError:
		log.Println(e2)
		fmt.Println("错误类型是*argError")
		break
	default:
		fmt.Println("未知错误")
		break
	}
	err := f3()
	// 原始的error
	log.Printf("%T", gherr.Cause(err))
	log.Printf("%v", err)
	// 如果+v 会显示调用栈
	log.Printf("%+v", err)
}
