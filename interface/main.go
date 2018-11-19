package main

import "fmt"

/**
  定义一个接口 animal 此类型 有speak方法
*/
type animal interface {
	speak()
}

// 定义一个对象狗
type dog struct {
	Name string
}

//
func (t *dog) speak() {
	fmt.Println("汪汪汪")
}

type cat struct {
	Name string
}

func (t *cat) speak() {
	fmt.Println("喵喵喵")
}

func main() {
	/**
	  这里右边是 *dog类型(指向dog类型的指针类型) 有speak方法
	  此时这个接口的实际类型 是 一个指针,这个指针有speak方法
	*/
	var animal0 animal = &dog{"小黑"}
	animal0.speak()
}
