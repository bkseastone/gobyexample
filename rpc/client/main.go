package main

import (
	"github.com/buffge/gobyexample/rpc"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

// type Args struct {
//     A int `json:"A"`
//     B int `json:"B"`
// }

func main() {
	conn, err := net.Dial("tcp", "47.96.13.188:7788")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var sum int
	err = client.Call("CalcService.Add", rpcdemo.Args{1, 4}, &sum)
	log.Println(sum, err)
	err = client.Call("CalcService.Add", rpcdemo.Args{1324, 344}, &sum)
	log.Println(sum, err)
	// 报错 json: cannot unmarshal number into Go value of type rpcdemo.Args
	err = client.Call("CalcService.Add", 1, &sum)
	log.Println(sum, err)
	var res string
	// HELLO BUFFGE nil
	err = client.Call("Hello.Upper", "hello buffge", &res)
	log.Println(res, err)
}
