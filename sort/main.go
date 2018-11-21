package main

import (
	"fmt"
	"sort"
)

type User struct {
	age int
}
type Users []User

/**
实现sort接口必须要有这3个方法
*/
func (s *Users) Len() int {
	return len(*s)
}
func (s *Users) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
func (s *Users) Less(i, j int) bool {
	return (*s)[i].age < (*s)[j].age
}
func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("字符串排序后:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("切片排序后:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("切片已排序 : ", s)

	times := []string{"2006-01-02 15:04:05", "2018-11-11 11:11:11"}
	sort.Strings(times)
	fmt.Println("日期排序后:", times)
	users := []User{{12}, {324}, {34}}
	sort.Sort((*Users)(&users))
	fmt.Println("实现sort接口", users)
}
