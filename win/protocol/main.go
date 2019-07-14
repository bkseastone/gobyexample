package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Printf("this is buffge protocol,curr pwd :%s", dir)
	// err := FilePutContents("buffge-protocol.logbuffge-protocol.log", "水电费很快就收到福建客户", 0666)
	// if err != nil {
	//     fmt.Println(err)
	// }
	fmt.Println(os.Args)
	fmt.Scanln()
}
