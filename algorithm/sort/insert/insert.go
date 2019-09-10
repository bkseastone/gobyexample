package main

import (
	"fmt"
	"time"

	sortConf "github.com/buffge/gobyexample/algorithm/sort"
	"github.com/buffge/gobyexample/rand/utils"
)

func main() {
	dataCount := sortConf.DataCount
	// dataCount = 100000
	arr := utils.GenerateRandomIntData(0, 100, dataCount)
	// arr = []int{1, 2, 5, 3, 1}
	// fmt.Println(arr)
	now := time.Now()
	sort(arr)
	duration := time.Now().Sub(now)
	fmt.Println(arr[0:10])
	fmt.Printf("共用时 %s\n", duration)
}

// 插入排序
// 将数组分为 已排序为未排序2部分 初始时 左边1个值,右边 n-1个值
// 用右边的每个值依次与已排序的值比较,如果右边值更小则交换
// init ori: [1,2,5,3,1]	   -> [1] [2,5,3,1]
// 1. left:[1] right:[2,5,3,1] -> [1,2] [5,3,1]
// 2. left:[1,2] right:[5,3,1] -> [1,2,5] [3,1]
// 3. left:[1,2,5] right:[3,1] -> [1,2,3] [5,1]
// 4. left:[1,2,3] right:[5,1] -> [1,2,3,5] [1]
// 5. left:[1,2,3,5] right:[1] -> [1,2,3,1] [5]
// 6. left:[1,2,3,1] right:[5] -> [1,2,1,3] [5]
// 7. left:[1,2,1,3] right:[5] -> [1,1,2,3] [5]
func sort(arr []int) {
	len := len(arr)
	var tmp int
	for i := 1; i < len; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			// fmt.Printf("left: %d,right: %d\n", arr[0:i], arr[i:])
			tmp = arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = tmp
		}
		// fmt.Printf("left: %d,right: %d\n", arr[0:i], arr[i:])
	}
}
