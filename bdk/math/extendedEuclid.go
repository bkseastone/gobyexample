package math

import (
	"fmt"
)

// ExtendedEuclid 扩展欧几里得算法
// a * x + b * y = gcd(a,b)
// remainder : 余
func ExtendedEuclid(a, b int) (remainder, x, y int) {
	// b=0 时，gcd(a,b)=a，此时 x=1 , y=0
	// 当b != 0 时
	// 假设 有解, 因为有一个解就有无限个解 (x +ka ,y-ka)
	// 这里设2个解
	// ax1+by1=gcd(a,b) , = gcd(b,a%b)=ax2+by2
	// 即  ax1+by1 = bx2+(a%b)y2
	// 因为 a%b =  a-a/b*b
	// 所以 ax1+by1 = bx2+(a-a/b*b)y2
	// = bx2 + ay2 - a/b * b*y2
	// = ay2 + bx2 - b * a/b *y2
	// = ay2 + b(x2-a/b *y2)
	// 所以 x1 = y2, y1 =x2 -a/b *y2
	// 最终我们有一个解就能得出上一个解
	// 翻译上面代码
	x = 1
	y = 0
	remainder = a
	if b == 0 {
		return
	}
	fmt.Println(a, x, b, y)
	for b != 0 {
		lastX := x
		x = y
		y = lastX - a/b*y
		lastA := a
		a = b
		b = lastA % b
		remainder = a*x + b*y
		fmt.Println(a, x, b, y)
	}
	// x, y = y, x
	return
}
