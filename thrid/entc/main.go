package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/buffge/gobyexample/thrid/entc/ent"
	_ "github.com/go-sql-driver/mysql"
)

var uname = ""
var pwd = ""

func main() {

	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp("+
		"47.96.13.188:3306)/blog?parseTime=True&charset=utf8mb4&loc=Local", uname, pwd))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	art, err := client.Article.Get(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	str, _ := json.Marshal(art)
	fmt.Println(string(str))

}
