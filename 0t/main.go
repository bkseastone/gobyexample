package main

import (
	"log"

	"github.com/alexedwards/argon2id"
)

func main() {
	log.Println(argon2id.CreateHash("qwer", argon2id.DefaultParams))
}
