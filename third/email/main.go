package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func main() {
	attachName := "main.go"
	smtpHost := "smtp.mxhichina.com"
	smtpPort := 465
	e := email.NewEmail()
	e.From = "buffge管理员 <noreply@buffge.com>"
	e.To = []string{"a@qq.com"}
	e.Bcc = []string{"b@qq.com"}
	e.Cc = []string{"c@qq.com"}
	e.Subject = "buffge测试标题"
	// e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1 style='color:red;'>buffge 测试内容</h1>")
	if _, err := e.AttachFile(attachName); err != nil {
		log.Fatalf("添加附件%s失败: %s", attachName, err)
	}
	auth := smtp.PlainAuth("", "noreply@buffge.com", "123123", smtpHost)
	begin := time.Now()
	if err := e.SendWithTLS(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, &tls.Config{InsecureSkipVerify: true,
		ServerName: smtpHost}); err != nil {
		log.Println("err: ", err)
	}
	log.Println(time.Since(begin))

}
