package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"gin-gorm-oj/define"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	IsAdmin  int    `json:"is_admin"`
	jwt.RegisteredClaims
}

func GetMd5(s string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(s)))
}

var mySigningKey = []byte("gin-gorm-oj-key")

func GenerateToken(identity, name string, isAdmin int) (string, error) {
	userClaim := &UserClaims{
		Identity:         identity,
		Name:             name,
		IsAdmin:          isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{},
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
	e.From = "郭心月44 <lepengxi@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("<b>乔丹</b>！您的验证码是：<b>" + code + "</b>")
	// 返回 EOF 时，关闭SSL重试
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "lepengxi@163.com", "JRLFKOBTMYSJCKPO", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}

// 生成验证码
// GetRand
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// CodeSave
// 保存代码
func CodeSave(code []byte) (string, error) {
	dirName := "code/" + GetUUID()
	path := dirName + "/main.go"
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, _ = f.Write(code)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	return path, nil
}

// CheckGoCodeValid
// 检查golang代码的合法性
func CheckGoCodeValid(path string) (bool, error) {
	//ioutil.ReadFile => os.ReadFile
	//b, err := io.ioutil(path)
	b, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	code := string(b)
	for i := 0; i < len(code)-6; i++ {
		if code[i:i+6] == "import" {
			var flag byte
			for i = i + 7; i < len(code); i++ {
				if code[i] == ' ' {
					continue
				}
				flag = code[i]
				break
			}
			if flag == '(' {
				for i = i + 1; i < len(code); i++ {
					if code[i] == ')' {
						break
					}
					if code[i] == '"' {
						t := ""
						for i = i + 1; i < len(code); i++ {
							if code[i] == '"' {
								break
							}
							t += string(code[i])
						}
						if _, ok := define.ValidGolangPackageMap[t]; !ok {
							return false, nil
						}
					}
				}
			} else if flag == '"' {
				t := ""
				for i = i + 1; i < len(code); i++ {
					if code[i] == '"' {
						break
					}
					t += string(code[i])
				}
				if _, ok := define.ValidGolangPackageMap[t]; !ok {
					return false, nil
				}
			}
		}
	}
	return true, nil
}
