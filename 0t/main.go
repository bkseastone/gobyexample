package main

import "log"

func main() {
	var a uint8 = 0b10110011
	log.Printf("%b", ^a)
	log.Printf("%b", ^0b10110011)
	log.Printf("%b", ^1011)
}
