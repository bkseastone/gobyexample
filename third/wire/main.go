//go:generate go run github.com/google/wire/cmd/wire

package main

import "log"

type A struct {
	b *B
}
type B struct {
	c *C
	d *D
}
type C struct {
}
type D struct {
}

func NewA(b *B) *A {
	return &A{b}
}
func NewB(c *C, d *D) *B {
	return &B{c, d}
}
func NewC() *C {
	return &C{}
}
func NewD() *D {
	return &D{}
}
func main() {
	a := InitA() // A 会自动注入
	a2 := &A{}   // 没有自动注入
	log.Println(a.b)
	log.Println(a2.b)
}
