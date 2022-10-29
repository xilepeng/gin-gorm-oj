package test

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Identiey string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte("gin-gorm-oj-key")

// 生成 token
func TestGenerateToken(t *testing.T) {

	userClaim := &UserClaims{
		Identiey: "乔丹",
		Name:     "乔丹",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v\n %v\n", ss, err)

}

// 解析 token
func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IuS5lOS4uSIsIm5hbWUiOiLkuZTkuLkifQ.qV7OkycrE6yz-WN9kvPLQtfHVLJQYzxo69rvrGnJqnU"
	userClaim := new(UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil // 第一次错误写法：[]byte("mySigningKey"), nil   应为: []byte("gin-gorm-oj-key"),nil
	})

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println("err:", err)
	}
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if token.Valid {
	//	fmt.Println(userClaim)
	//}
}
