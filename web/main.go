package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	pl = fmt.Println
)

func handHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
	header, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, string(header)) //这个写入到w的是输出到客户端的
}
func main() {
	port := 8893
	http.HandleFunc("/", handHttp)                          //设置访问的路由
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
