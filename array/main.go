package main

import "fmt"

/**
定义数组后数组的长度是不可以改变的
*/
func main() {
	//初始化数组
	arr0 := [5]int{1, 23}
	fmt.Println(arr0)
	//获取数组长度
	fmt.Printf("数组arr0d的长度是%v\n", len(arr0))
	//定义一个长度为fgo5的int数组
	var arr1 [5]int
	fmt.Println(arr1)
	//给数组赋值
	arr1 = [5]int{24, 34, 324, 234, 4}
	fmt.Println(arr1)
	//定义一个不定长度的int数组
	var arr2 []int
	fmt.Println("arr2: ", arr2)
	arr2 = append(arr2, 1)
	fmt.Println("arr2 after append: ", arr2)
	//初始化一个不定长度的切片,数组必须是定长
	arr3 := []int{1, 23, 3434, 34, 34, 35}
	fmt.Println("arr3: ", arr3)
	//二维数组
	var arr4 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr4[i][j] = i + j
		}
	}
	fmt.Println("二维数组: ", arr4)

}
