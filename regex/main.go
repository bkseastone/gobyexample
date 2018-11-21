package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	match, _ := regexp.MatchString("123p([a-z]+)ch", "peach")
	fmt.Println("查看字符串中是否匹配某个regex字符串", match)
	//编译一个正则对象
	r, _ := regexp.Compile("p([a-z]+)ch")
	//使用正则对象匹配
	fmt.Println("使用正则对象匹配 ", r.MatchString("peach"))

	fmt.Println("使用正则对象匹配一个字符串 ", r.FindString("peach punch"))

	fmt.Println("返回字符串的前后index[)", r.FindStringIndex("peach punch"))
	fmt.Println("返回字符串中匹配值和子组", r.FindStringSubmatch("peach punch"))

	fmt.Println("返回字符串中匹配值和子组的前后index[)", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("使用正则对象匹配-1(无限)个字符串",
		r.FindAllString("peach punch pinch", -1))
	fmt.Println("获取所有匹配的字符串和子组", r.FindAllStringSubmatch(
		"peach punch pinch", -1))
	fmt.Println("获取所有匹配的字符串和子组前后index", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println("使用正则对象匹配2个字符串",
		r.FindAllString("peach punch pinch", 2))
	fmt.Println("对字节数组的匹配 ", r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("MustCompile 就是一定要编译通过 不然panic", r)
	fmt.Println("将所有匹配到值替换为某个值",
		r.ReplaceAllString("a peach", "<fruit>"))
	fmt.Println("使用函数替换所有的匹配到的字符串",
		r.ReplaceAllStringFunc("a peach 234", strings.ToTitle))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("使用函数替换所有的匹配到的字节数组", string(out))
}
