package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("遍历数组求和:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("遍历数组取出值为3的那个index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("遍历映射")
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println("遍历映射的键")
	for k := range kvs {
		fmt.Println("key:", k)
	}
	fmt.Println("遍历字符串")
	for i, c := range "go" {
		fmt.Println(i, string(c))
	}
}
