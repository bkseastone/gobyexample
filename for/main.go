package main

import "fmt"

/**
for 循环
*/

func main() {
	i := 1
	//相当于while
	for i < 5 {
		fmt.Println(i)
		i++
	}
	for i = 7; i < 9; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("这是无限循环")
		break
	}
	//可以continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
