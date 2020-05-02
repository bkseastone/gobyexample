package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

func DeepClone(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(dst)
}
func main() {
	user1 := &User{
		"buffge",
		26,
	}
	var user2 User
	var user3 = &User{}
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf) // 创建编码器
	err1 := encoder.Encode(user1)   // 编码
	if err1 != nil {
		log.Panic(err1)
	}
	fmt.Printf("序列化后：%x\n", buf.Bytes())
	// 反序列化
	byteEn := buf.Bytes()
	decoder := gob.NewDecoder(bytes.NewReader(byteEn)) // 创建解密器
	err2 := decoder.Decode(&user2)                     // 解密
	if err2 != nil {
		log.Panic(err2)
	}
	fmt.Println("反序列化后：", user2)
	// gob.Register(User{})
	if err := DeepClone(user1, user3); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("深克隆后：", user3)
}
