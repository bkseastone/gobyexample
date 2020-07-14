package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var (
	pl = log.Println
)

func startClient(ip string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	if err != nil {
		pl("服务器端", conn.RemoteAddr(), "已关闭连接...")
		return
	}
	buf := make([]byte, 902400)
	pl("服务器端地址 ", conn.RemoteAddr())
	for {
		conn.Read(buf)
		reqStr := byte2Str(buf)
		pl("服务器端的返回是 ", reqStr)
		if reqStr == "ping" || reqStr == "/favicon.ico" {
			continue
		}
		resp, err := http.Get("http://127.0.0.1" + reqStr)
		if err != nil {
			_ = resp.Body.Close()
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			pl("请求本地失败", err.Error())
			return
		}
		_, _ = conn.Write(body)
		_ = resp.Body.Close()
	}
}
func byte2Str(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
func main() {
	startClient("47.96.13.188", 7788)
}

/*
	nat
	客户端1 client c1
	客户端2 client c2
	服务器 server s
c1 -> s
		s获取到c1的端口
3.打洞过程

双打洞客户端
（1）A请求Server。
（2）B请求Server。
（3）Server把A的IP和端口信息发给B。
（4）Server把B的IP和端口信息发给A。
（5）A利用信息给B发消息。（A信任B）
（6）B利用信息给A发消息。（B信任A）

单打洞客户端
    A 请求 S
	S 保存 A 的 端口 并且为A 分配一个域名
	S 定时 发送消息给 A
	A 定时 发送消息给 S
	B 用A分配到的域名访问 S http请求
	S 发送消息给A
	A 转发消息到本地的80端口
	A返回消息给S
	S 返回给B

*/
