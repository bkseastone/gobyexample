package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/alexedwards/argon2id"
	"log"
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
	fmt.Printf("md5值 %X\n", bs2)
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
	// 计算argon2id
	hash, err := argon2id.CreateHash("qwer", argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(hash)
	}
	// 验证argon2id
	match, err := argon2id.ComparePasswordAndHash("qwer", hash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Match: %v", match)
	pwdHash := "$argon2id$v=19$m=1024,t=2," +
		"p=2$OEVrSUtGSEFQNmRrengyNg$u0N+bFvuyrg5V0vgDxk2STa4Os8mnOzAm+Bi+tjvPa8"
	isRightPwd, err := argon2id.ComparePasswordAndHash("qwer", pwdHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("密码验证结果: %v", isRightPwd)
}
