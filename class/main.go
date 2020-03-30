package main

import "fmt"

type A struct {
	AName string
}
type B struct {
	*A
	BName string
}

func (t *A) Call() {
	fmt.Printf("%v\n", t)
}
func main() {
	b := &B{
		A: &A{
			AName: "buffge A",
		},
		BName: "buffge B",
	}
	b.Call()
}
