package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type UserClaims struct {
	Identiey string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(s)))
}

var mySigningKey = []byte("gin-gorm-oj-key")

func GenerateToken(identity, name string) (string, error) {
	userClaim := &UserClaims{
		Identiey: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v\n %v\n", ss, err)
	return ss, err
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	//tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IuS5lOS4uSIsIm5hbWUiOiLkuZTkuLkifQ.qV7OkycrE6yz-WN9kvPLQtfHVLJQYzxo69rvrGnJqnU"
	userClaim := new(UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil // 第一次错误写法：[]byte("mySigningKey"), nil   应为: []byte("gin-gorm-oj-key"),nil
	})

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims, claims.RegisteredClaims.Issuer)
		return userClaim, nil
	} else {
		fmt.Println("err:", err)
		return nil, err
	}
}

func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "郭心月 <lepengxi@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<b>乔丹</b>！您的验证码是：<b>" + code + "</b>")
	// 返回 EOF 时，关闭SSL重试
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "lepengxi@163.com", "JRLFKOBTMYSJCKPO", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
