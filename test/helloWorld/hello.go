package main

import "fmt"

func Hello(str string) string {
	return "Hello " + str
}

func main() {
	fmt.Println(Hello("World!"))
}
