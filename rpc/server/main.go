package main

import (
	"github.com/buffge/gobyexample/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// rpc.RegisterName("buffge", rpcdemo.CalcService{})
	err := rpc.Register(rpcdemo.CalcService{})
	if err != nil {
		panic(err)
	}
	// 可以注册多个服务
	err = rpc.Register(rpcdemo.Hello{})
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":7788")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		log.Println("接收到新请求")
		go jsonrpc.ServeConn(conn)
	}
}
