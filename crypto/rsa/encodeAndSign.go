package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
)

// 加密和签名
func encodeAndSign() {
	oriData := "buffge"
	// 获取私钥
	pk, _ := loadPrivateKey(str2Bytes(privateKeyData))
	// 公钥加密
	pubEnStr := pubKeyEncode(oriData, &pk.PublicKey)
	// 私钥解密
	deData, _ := priKeyDecode(pubEnStr, pk)
	log.Println("公钥加密 == 私钥解密 ? = ", bytes2Str(deData) == oriData)
	signData, _ := pkSignWithSha256(str2Bytes(oriData), pk)
	err := pubVerySignWithSha256(str2Bytes(oriData), signData, &pk.PublicKey)
	log.Println("私钥签名 == 公钥验签 err = ", err)

}
func loadPrivateKey(bts []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bts)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
func getPubKeyBlock(privateKeyData string) (*pem.Block, error) {
	bts := str2Bytes(privateKeyData)
	pk, err := loadPrivateKey(bts)
	if err != nil {
		return nil, err
	}
	derPkix, err := x509.MarshalPKIXPublicKey(&pk.PublicKey)
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	// pubKey := pem.EncodeToMemory(block)
	// log.Println(bytes2Str(pubKey))
	return block, err
}
func pubKeyEncode(oriStr string, publicKey *rsa.PublicKey) string {
	// block, _ := pem.Decode([]byte(pubKey))
	// publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	// if err != nil {
	// 	panic(err)
	// }
	// publicKey := publicKeyInterface.(*rsa.PublicKey)
	enBts, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, str2Bytes(oriStr))
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(enBts)
}
func priKeyDecode(pubEnB64Data string, pk *rsa.PrivateKey) ([]byte, error) {
	cipherBts, err := base64.StdEncoding.DecodeString(pubEnB64Data)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, pk, cipherBts)
}
func pkSignWithSha256(data []byte, pk *rsa.PrivateKey) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA256, hashed)
}
func pubVerySignWithSha256(data, signData []byte, pubKey *rsa.PublicKey) error {
	hashed := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signData)
}
