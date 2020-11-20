package main

import (
	"fmt"
	"log"
)

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

type Val struct {
	val int
}

func (v Val) set(newVal int) {
	v.val = newVal
}
func (v *Val) ptrSet(newVal int) {
	v.val = newVal
}
func main() {
	/**
	  这里右边是 *dog类型(指向dog类型的指针类型) 有speak方法
	  此时这个接口的实际类型 是 一个指针,这个指针有speak方法
	*/
	var animal0 animal = &dog{"小黑"}
	animal0.speak()
	v1 := Val{1}
	// 只要方法是值接收值 那么无法改变调用者本身
	v1.set(12)
	log.Println(v1.val) // 1
	v2 := &Val{1}
	v2.set(12)
	log.Println(v2.val) // 1
	// 如果方法是指针接收值 那么可以改变调用者本身
	v1.ptrSet(12)
	log.Println(v1.val) // 12
	v2.ptrSet(12)
	log.Println(v2.val) // 12
}
