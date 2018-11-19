package main

import "fmt"

/**
判断映射键是否存在用 v,ok = m[k] 根据ok来判断
不能用 m[k0] == m[k1] 来判断
*/
func main() {
	//创建一个string -> int 的映射
	m := make(map[string]string)

	m["u1"] = "buffge"
	m["u2"] = "zty"
	m["u3"] = ""

	fmt.Println("映射:", m)

	v1 := m["k1"]
	fmt.Println("映射中不存在的k1 : ", v1 == "")

	fmt.Println("映射的长度为:", len(m))
	//删除掉k2键
	delete(m, "k2")
	fmt.Println("映射:", m)

	_, ok := m["k2"]
	fmt.Println("映射中是否存在k2键:", ok)
	vu3, ok := m["u3"]
	fmt.Println("映射中是否存在u3键:", ok, vu3)
	fmt.Println("u3 == k2:", m["u3"] == m["k2"])
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("初始化的映射:", n)
}
