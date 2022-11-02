package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	// go run code-user/main.go
	cmd := exec.Command("go", "run", "code/code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	_, _ = io.WriteString(stdinPipe, "1 1\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	// 根据测试的输入案例进行运行，拿到输出结果和标准输出结果进行比对

	println("Err:", string(stderr.Bytes()))
	fmt.Println(out.String())

	println(out.String() == "2\n")

}
