package main

import "fmt"

type User struct {
	Id   uint
	Name string
	Age  int
	//这是私有变量 不可以读取,这里不知道为什么还能读取
	wifeName string
}

//方法
func (t *User) login() {
	fmt.Printf("普通用户%s你好你已登录\n", t.Name)
}

func main() {
	u0 := User{1, "buffge", 24, "zty"}
	fmt.Println("现在u0的id是 ", u0.Id)
	u0.Id = 2
	fmt.Println("设置u0的id为2之后 id = ", u0.Id)
	u1 := &u0
	u1.Id = 3
	fmt.Println("u1为指向u0的指针,改变u1的id后 u0 id = ", u0.Id)
	fmt.Println(u0)
	u0.login()
}
