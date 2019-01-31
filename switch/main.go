package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	fmt.Print("i is ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 7:
		fmt.Println("four or five or seven")
	default:
		fmt.Println("un known")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("现在是周末")
	default:
		fmt.Println("现在不是周末")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("现在是12点以前")
	default:
		fmt.Println("现在是12点以后")
	}
	/**
	接收所有值
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("是一个布尔值")
		case int:
			fmt.Println("是一个整数值")
		default:
			fmt.Printf("未知类型 %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
