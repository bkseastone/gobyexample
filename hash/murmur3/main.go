package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

func main() {
	hash32 := murmur3.New32()
	hash64 := murmur3.New64()
	hash128 := murmur3.New128()
	_, _ = hash32.Write([]byte{'a'})
	_, _ = hash64.Write([]byte{'a'})
	_, _ = hash128.Write([]byte{'a'})
	fmt.Printf("%x\n", hash32.Sum32())
	fmt.Printf("%x\n", hash64.Sum64())
	v1, v2 := hash128.Sum128()
	fmt.Printf("%x%x\n", v1, v2)

}
