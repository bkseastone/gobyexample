package main

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	app := iris.New()
	app.Post("/", test)

	app.Listen(":8082")
}

type St struct {
	Name string `json:"name"`
	Num  int    `json:"num,omitempty"`
}

func test(ctx iris.Context) {
	st1 := &St{}
	if err := ctx.ReadBody(st1); err != nil {
		log.Println("bind error")
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}
	bts, _ := json.Marshal(st1)
	log.Println(st1)
	ctx.HTML(string(bts))
	ctx.StatusCode(iris.StatusCreated)

}
