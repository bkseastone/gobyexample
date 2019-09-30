package main

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want string) {
	// 声明此函数是一个辅助函数,
	// 最后如果测试未通过时会显示调用此辅助函数的行号
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func TestHello(t *testing.T) {
	// 子任务
	t.Run("Say Hello World!", func(t *testing.T) {
		got := Hello("World!")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello Buffge!", func(t *testing.T) {
		got := Hello("Buffge!")
		want := "Hello Buffge!"
		assertCorrectMessage(t, got, want)
	})
}
func ExampleHello() {
	res := Hello("Buffge!")
	fmt.Println(res)
	// Output: Hello Buffge!
}
