package math

import (
	"math/rand"
	"testing"
	"time"
)

func assertCorrectMessage(t *testing.T, got, want string) {
	// 声明此函数是一个辅助函数,
	// 最后如果测试未通过时会显示调用此辅助函数的行号
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func TestExtendedEuclid(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1; i++ {
		t.Run(t.Name(), func(t *testing.T) {
			a := rand.Intn(1e10)
			b := rand.Intn(1e10)
			a = 30
			b = 47
			_, x, y := ExtendedEuclid(a, b)
			if a*x+b*y == a*x+b*y {
				t.Errorf("ExtendedEuclid() %v*%v+%v*%v != %v", a, x, b, y, a*x+b*y)
			}
		})
	}
}
