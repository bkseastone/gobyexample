package main

import (
	"fmt"
	"strings"
)
import "net/url"

func main() {
	s := "postgres://host.com/path?k=123v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("协议: ", u.Scheme)
	fmt.Println("用户信息: ", u.User)
	fmt.Println("用户名: ", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("用户密码: ", p)

	fmt.Println("域名包含端口: ", u.Host)
	//host, port, _ := net.SplitHostPort(u.Host)
	hostInfo := strings.Split(u.Host, ":")
	fmt.Println("域名: ", hostInfo[0])
	var port string
	if len(hostInfo) == 1 {
		port = "80"
	} else {
		port = hostInfo[1]
	}
	fmt.Println("端口: ", port)
	fmt.Println("路径: ", u.Path)
	fmt.Println("hash值: ", u.Fragment)
	fmt.Println("原生查询语句: ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("查询语句", m)
	fmt.Println("k的值", m["k"][0])
}
