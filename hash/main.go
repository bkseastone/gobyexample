package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)
import "fmt"

func main() {
	s := "buffge"
	fmt.Println("原始字符串", s)
	//md5
	m5 := md5.New()
	m5.Write([]byte(s))
	//计算hash值
	bs2 := m5.Sum(nil)
	fmt.Printf("md5值 %x\n", bs2)
	//sha1
	//新建sha1对象
	h := sha1.New()
	//将字节数组写入到sha1对象中
	h.Write([]byte(s))
	//计算hash值
	bs := h.Sum(nil)
	fmt.Printf("sha1值 %x\n", bs)
	//sha256
	s256 := sha256.New()
	s256.Write([]byte(s))
	//计算hash值
	bs256 := s256.Sum(nil)
	fmt.Printf("sha256值 %x\n", bs256)
	//sha256
	s512 := sha512.New()
	s512.Write([]byte(s))
	//计算hash值
	bs512 := s512.Sum(nil)
	fmt.Printf("sha512值 %x\n", bs512)
}
