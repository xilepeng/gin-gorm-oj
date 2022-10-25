package main

import "gin-gorm-oj/router"

func main() {
	r := router.Router()
	err := r.Run()
	if err != nil {
		return
	}
}
