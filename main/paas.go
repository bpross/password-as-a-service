package main

import (
	"fmt"
	pw "github.com/bpross/password-as-a-service/password"
)

func main() {
	password := "angryMonkey"
	passwordSha := pw.HashPassword512(password)
	fmt.Println(passwordSha)
}
