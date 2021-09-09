package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"strings"
)

var pubKey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp7r7hvHyMlJ4h1NuDX5l
MMJson2XuG9jQWVnh7fxyWAehVqvQ720n+vAfedVsvrDoNlSfJyvZTETejjGjCXM
c6wvbmSOmQ/96gUnDRMLM+Wr2VRKFPJGVCBhLuIjvQzQb3VSA4KbVWaPiVhvvMoW
aedyBLfjBhqGkOufVtiIGyp88Y4B6QaLqKLsibsMr084w4ZxWk9bmfjDKqlMAlOM
Q09POpGTLOEattNIljUuCCkKMkyngC1VyyUdxaNatrPoCo1cR8oP7Fpx0f9ElI8G
NmIfniOmQoyA9gn9vllITcFi0WiGlA7jjvuA28TXCyClIIrAYvH2ePKWHGMxaYnn
bQIDAQAB
-----END PUBLIC KEY-----
`

func main() {
	oriStr := "buffge"
	log.Println(encode(oriStr))
}

func encode(oriStr string) string {
	block, _ := pem.Decode([]byte(pubKey))
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	enBts, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(oriStr))
	if err != nil {
		panic(err)
	}
	return inyuEncode(base64.StdEncoding.EncodeToString(enBts))
}
func inyuEncode(str string) string {
	str = strings.ReplaceAll(str, "9", "!")
	str = strings.ReplaceAll(str, "b", "_")
	return strings.ReplaceAll(str, "2", "@")
}
