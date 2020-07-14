package main

import (
	"log"
	"net/http"
)

func main() {
	responsePublicKeyURL, err := http.Get("https://gosspublic.alicdn.com/callback_pub_key_v1.pem")
	log.Println(responsePublicKeyURL, err)
}
