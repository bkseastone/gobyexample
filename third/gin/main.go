package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

var arr []*user

type (
	user struct {
		name string
	}
)

func main() {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		a := &user{
			name: "buff",
		}
		arr = append(arr, a)
		c.String(200, "ssss")
	})
	go http.ListenAndServe(":6601", nil)
	log.Fatal(engine.Run(":8080"))
}
