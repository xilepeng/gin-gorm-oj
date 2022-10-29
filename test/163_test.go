package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "郭心月 <lepengxi@163.com>"
	e.To = []string{"18800138580@163.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<b>乔丹</b>！您的验证码是：<b>123456</b>")
	// 返回 EOF 时，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "lepengxi@163.com", "JRLFKOBTMYSJCKPO", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
